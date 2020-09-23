package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/dharani/types"
	"github.com/dharani/x/dharani/types"
)

func (k Keeper) SetPropertyCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryLengthPrefixed(count)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.PropertyCountKey, value)
}

func (k Keeper) GetPropertyCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.PropertyCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &count)
	return count
}

func (k Keeper) SetProperty(ctx sdk.Context, id types2.PropertyID, property types.Property) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetPropertyKey(id)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(property)

	store.Set(key, value)
}

func (k Keeper) GetProperty(ctx sdk.Context, id types2.PropertyID) (property *types.Property) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetPropertyKey(id)

	value := store.Get(key)
	if value == nil {
		return nil
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &property)
	
	return property
}
