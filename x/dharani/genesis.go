package dharani

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dharani/x/dharani/types"

	// abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, property := range data.Properties {
		key := types.GetPropertyKey(property.Owner, []byte(property.ID))
		k.SetProperty(ctx, key, property)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	return ExportGenesisState(ctx, k)
}

func ExportGenesisState(ctx sdk.Context, k Keeper) GenesisState {
	properties := k.GetAllProperties(ctx)
	return GenesisState{
		Properties: properties,
	}
}
