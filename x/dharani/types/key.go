package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dharani/types"
	"strconv"
)

const (
	// ModuleName is the name of the module
	ModuleName = "dharani"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)

var (
	PropertyCountKey = []byte{0x00}
	PropertyKey      = []byte{0x01}

	TypeOwn  = "OWN"
	TypeSell = "SELL"
)

func GetPropertyCountKey(c uint64) []byte {
	return append(PropertyCountKey, sdk.Uint64ToBigEndian(c)...)
}

func GetPropertyKey(id []byte) []byte {
	return append(PropertyKey, id...)
}

func GetPropertyID(count uint64) string {
	return types.PropertyIDPrefix + strconv.Itoa(int(count))
}