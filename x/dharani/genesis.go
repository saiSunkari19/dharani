package dharani

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, property := range data.Properties {
		k.SetProperty(ctx, property.ID, property)
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