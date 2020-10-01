package app

import (
	"encoding/json"
	"io"
	"os"
	
	"github.com/dharani/x/dharani"
	
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"
	
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	codecstd "github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	port "github.com/cosmos/cosmos-sdk/x/ibc/05-port"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	// this line is used by starport scaffolding
)

const appName = "dharani"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.dharanicli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.dharanid")
	
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		
		ibc.AppModuleBasic{},
		transfer.AppModuleBasic{},
		
		dharani.AppModuleBasic{},
	
	)
	
	maccPerms = map[string][]string{
		auth.FeeCollectorName:           nil,
		distr.ModuleName:                nil,
		staking.BondedPoolName:          {auth.Burner, auth.Staking},
		staking.NotBondedPoolName:       {auth.Burner, auth.Staking},
		transfer.GetModuleAccountName(): {auth.Minter, auth.Burner},
	}
)

func MakeCodecs() (*std.Codec, *codec.Codec) {
	cdc := std.MakeCodec(ModuleBasics)
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	
	sdk.RegisterInterfaces(interfaceRegistry)
	ModuleBasics.RegisterInterfaceModules(interfaceRegistry)
	appCodec := std.NewAppCodec(cdc, interfaceRegistry)
	
	return appCodec, cdc
}

type NewApp struct {
	*bam.BaseApp
	cdc *codec.Codec
	
	invCheckPeriod uint
	
	keys    map[string]*sdk.KVStoreKey
	tKeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey
	
	subspaces map[string]params.Subspace
	
	accountKeeper    auth.AccountKeeper
	bankKeeper       bank.Keeper
	stakingKeeper    staking.Keeper
	capabilityKeeper *capability.Keeper
	slashingKeeper   slashing.Keeper
	distrKeeper      distr.Keeper
	paramsKeeper     params.Keeper
	ibcKeeper        *ibc.Keeper
	transferKeeper   transfer.Keeper
	dharaniKeeper    dharani.Keeper
	
	scopedIBCKeeper      capability.ScopedKeeper
	scopedTransferKeeper capability.ScopedKeeper
	
	mm *module.Manager
	
	sm *module.SimulationManager
}

var _ simapp.App = (*NewApp)(nil)

func NewInitApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string,
	baseAppOptions ...func(*bam.BaseApp),
) *NewApp {
	appCodec, cdc := MakeCodecs()
	
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	
	keys := sdk.NewKVStoreKeys(
		auth.StoreKey,
		bank.StoreKey,
		staking.StoreKey,
		distr.StoreKey,
		slashing.StoreKey,
		params.StoreKey,
		ibc.StoreKey,
		transfer.StoreKey,
		capability.StoreKey,
		
		dharani.StoreKey,
	
	
	)
	
	tKeys := sdk.NewTransientStoreKeys(params.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capability.MemStoreKey)
	
	var app = &NewApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
		subspaces:      make(map[string]params.Subspace),
	}
	
	app.paramsKeeper = params.NewKeeper(appCodec, keys[params.StoreKey], tKeys[params.TStoreKey])
	app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.paramsKeeper.Subspace(staking.DefaultParamspace)
	app.subspaces[distr.ModuleName] = app.paramsKeeper.Subspace(distr.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	
	app.subspaces[dharani.ModuleName] = app.paramsKeeper.Subspace(dharani.DefaultParamspace)
	
	bApp.SetParamStore(app.paramsKeeper.Subspace(bam.Paramspace).WithKeyTable(std.ConsensusParamsKeyTable()))
	
	app.capabilityKeeper = capability.NewKeeper(appCodec, keys[capability.StoreKey], memKeys[capability.MemStoreKey])
	scopedIBCKeeper := app.capabilityKeeper.ScopeToModule(ibc.ModuleName)
	scopedTransferKeeper := app.capabilityKeeper.ScopeToModule(transfer.ModuleName)
	
	app.accountKeeper = auth.NewAccountKeeper(
		appCodec,
		keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount,
		maccPerms,
	)
	
	app.bankKeeper = bank.NewBaseKeeper(
		appCodec,
		keys[bank.StoreKey],
		app.accountKeeper,
		app.subspaces[bank.ModuleName],
		app.ModuleAccountAddrs(),
	)
	
	stakingKeeper := staking.NewKeeper(
		appCodec,
		keys[staking.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		app.subspaces[staking.ModuleName],
	)
	
	app.distrKeeper = distr.NewKeeper(
		appCodec,
		keys[distr.StoreKey],
		app.subspaces[distr.ModuleName],
		app.accountKeeper,
		app.bankKeeper,
		&stakingKeeper,
		auth.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	
	app.slashingKeeper = slashing.NewKeeper(
		appCodec,
		keys[slashing.StoreKey],
		&stakingKeeper,
		app.subspaces[slashing.ModuleName],
	)
	
	app.ibcKeeper = ibc.NewKeeper(
		app.cdc,
		appCodec,
		keys[ibc.StoreKey],
		stakingKeeper,
		scopedIBCKeeper,
	)
	
	app.transferKeeper = transfer.NewKeeper(
		appCodec,
		keys[transfer.StoreKey],
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.bankKeeper,
		scopedTransferKeeper,
	)
	
	app.dharaniKeeper = dharani.NewKeeper(app.bankKeeper, app.cdc, keys[dharani.StoreKey])
	
	transferModule := transfer.NewAppModule(app.transferKeeper)
	
	ibcRouter := port.NewRouter()
	ibcRouter.AddRoute(transfer.ModuleName, transferModule)
	
	app.ibcKeeper.SetRouter(ibcRouter)
	
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distrKeeper.Hooks(),
			app.slashingKeeper.Hooks(),
		),
	)
	
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(appCodec, app.accountKeeper),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(appCodec, *app.capabilityKeeper),
		slashing.NewAppModule(appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		distr.NewAppModule(appCodec, app.distrKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		params.NewAppModule(app.paramsKeeper),
		
		dharani.NewAppModule(app.dharaniKeeper, app.bankKeeper),
		
		transferModule,
	
	)
	
	app.mm.SetOrderBeginBlockers(
		distr.ModuleName,
		slashing.ModuleName,
		staking.ModuleName,
		ibc.ModuleName,
	)
	
	app.mm.SetOrderEndBlockers(staking.ModuleName)
	
	app.mm.SetOrderInitGenesis(
		capability.ModuleName,
		auth.ModuleName,
		distr.ModuleName,
		staking.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		ibc.ModuleName,
		genutil.ModuleName,
		transfer.ModuleName,
		
		dharani.ModuleName,
	
	
	)
	
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
	
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	
	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)
	app.MountMemoryStores(memKeys)
	
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.accountKeeper,
			app.bankKeeper,
			*app.ibcKeeper,
			ante.DefaultSigVerificationGasConsumer,
		),
	)
	app.SetEndBlocker(app.EndBlocker)
	
	if loadLatest {
		err := app.LoadLatestVersion()
		if err != nil {
			tmos.Exit(err.Error())
		}
	}
	
	ctx := app.BaseApp.NewUncachedContext(true, abci.Header{})
	app.capabilityKeeper.InitializeAndSeal(ctx)
	
	app.scopedIBCKeeper = scopedIBCKeeper
	app.scopedTransferKeeper = scopedTransferKeeper
	
	return app
}

type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState() GenesisState {
	cdc := codecstd.MakeCodec(ModuleBasics)
	return ModuleBasics.DefaultGenesis(cdc)
	
}

func (app *NewApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	
	return app.mm.InitGenesis(ctx, app.cdc, genesisState)
}

func (app *NewApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *NewApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *NewApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

func (app *NewApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[auth.NewModuleAddress(acc).String()] = true
	}
	
	return modAccAddrs
}

func (app *NewApp) Codec() *codec.Codec {
	return app.cdc
}

func (app *NewApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

func GetMaccPerms() map[string][]string {
	modAccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		modAccPerms[k] = v
	}
	return modAccPerms
}
