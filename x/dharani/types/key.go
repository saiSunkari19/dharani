package types

import (
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
	PropertyCountKey = []byte{0x00}
	PropertyKey      = []byte{0x01}
	
	TypeOwn  = "OWN"
	TypeSell = "SELL"
)

func GetPropertyKey(id []byte) []byte {
	return append(PropertyKey, id...)
}

func GetPropertyID(count uint64) string {
	return types.PropertyIDPrefix + strconv.Itoa(int(count))
}
