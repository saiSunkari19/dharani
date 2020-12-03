package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dharani/x/dharani/types"
)

func (k Keeper) SetPropertyCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryBare(count)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.PropertyCountKey, value)
}

func (k Keeper) GetPropertyCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.PropertyCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryBare(value, &count)
	return count
}

func (k Keeper) SetMarketPalcePropertyCount(ctx sdk.Context, count uint64) {
	value := k.cdc.MustMarshalBinaryBare(count)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.MarketPlacePropertyCountKey, value)
}

func (k Keeper) GetMarketPlacePropertyCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.MarketPlacePropertyCountKey)
	if value == nil {
		return 0
	}

	k.cdc.MustUnmarshalBinaryBare(value, &count)
	return count
}

func (k Keeper) SetProperty(ctx sdk.Context, key []byte, property types.Property) {
	store := ctx.KVStore(k.storeKey)

	value := k.cdc.MustMarshalBinaryBare(property)

	store.Set(key, value)
}

func (k Keeper) GetProperty(ctx sdk.Context, key []byte) (property *types.Property) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(key)
	if value == nil {
		return nil
	}

	k.cdc.MustUnmarshalBinaryBare(value, &property)

	return property
}

func (k Keeper) GetAllProperties(ctx sdk.Context) []types.Property {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PropertyKey)
	defer iterator.Close()

	var properties []types.Property
	for ; iterator.Valid(); iterator.Next() {
		var property types.Property
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &property)
		properties = append(properties, property)
	}

	return properties
}

func (k Keeper) GetPropertyByAddress(ctx sdk.Context, addr sdk.AccAddress) []types.Property {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PropertyKey)
	defer iterator.Close()

	var properties []types.Property
	for ; iterator.Valid(); iterator.Next() {
		var property types.Property
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &property)
		if bytes.Compare(property.Owner.Bytes(), addr.Bytes()) == 0 {
			properties = append(properties, property)
		}
	}

	return properties
}
