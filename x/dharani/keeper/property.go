package keeper

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) SetProperty(ctx sdk.Context, id string, property types.Property) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetPropertyKey([]byte(id))
	value := k.cdc.MustMarshalBinaryLengthPrefixed(property)

	store.Set(key, value)
}

func (k Keeper) GetProperty(ctx sdk.Context, id string) (property *types.Property) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetPropertyKey([]byte(id))
	value := store.Get(key)
	if value == nil {
		return nil
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &property)

	return property
}

func (k Keeper) GetAllProperties(ctx sdk.Context) []types.Property {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PropertyKey)
	defer iterator.Close()

	var properties []types.Property
	for ; iterator.Valid(); iterator.Next() {
		var property types.Property
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &property)
		properties = append(properties, property)
	}

	return properties
}

func (k Keeper) GetPropertyByAddress(ctx sdk.Context, addr sdk.AccAddress) []types.Property {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.PropertyKey)
	defer iterator.Close()

	var properties []types.Property
	for ; iterator.Valid(); iterator.Next() {
		var property types.Property
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), property)
		if bytes.Compare(property.Owner.Bytes(), addr.Bytes()) == 0 {
			properties = append(properties, property)
		}
	}

	return properties
}
