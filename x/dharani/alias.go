package dharani

import (
	"github.com/dharani/x/dharani/keeper"
	"github.com/dharani/x/dharani/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
	EventTypeMsgAddProperty  = types.EventTypeMsgAddProperty
	EventTypeMsgSellProperty = types.EventTypeMsgSellProperty
	EventTypeMsgBuyProperty  = types.EventTypeMsgBuyProperty

	AttributeKeyFromAddress = types.AttributeKeyFromAddress
	AttributeKeyPropertyID  = types.AttributeKeyPropertyID
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	// TODO: Fill out module types
)
