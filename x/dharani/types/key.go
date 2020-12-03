package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/dharani/types"
)

const (
	ModuleName = "dharani"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	PropertyCountKey            = []byte{0x00}
	MarketPlacePropertyCountKey = []byte{0x01}
	PropertyKey                 = []byte{0x02}
	PropertyMarketPlaceKey      = []byte{0x03}
	PropertyBoughtKey           = []byte{0x04}

	TypeOwn    = "OWN"
	TypeBought = "BOUGHT"
)

func GetPropertyKey(owner sdk.AccAddress, id []byte) []byte {
	return append(PropertyKey, append(owner, id...)...)
}

func GetMarketPlacePropertyKey(id []byte) []byte {
	return append(PropertyMarketPlaceKey, id...)
}

func GetPropertyBoughtKey(id []byte) []byte {
	return append(PropertyBoughtKey, id...)
}

func GetPropertyID(count uint64) string {
	return types.PropertyIDPrefix + strconv.Itoa(int(count))
}
