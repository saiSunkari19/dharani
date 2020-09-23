package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dharani/types"
)

type Property struct {
	ID          types.PropertyID
	Area        int64
	Owner       sdk.AccAddress
	Location    string
	Type        string
	RefProperty types.PropertyID
	Cost        sdk.Coin
}

func NewProperty(id types.PropertyID, area int64, owner sdk.AccAddress, location string,
	_type string, refPropID types.PropertyID, cost sdk.Coin) Property {
	return Property{
		ID:          id,
		Area:        area,
		Owner:       owner,
		Location:    location,
		Type:        _type,
		RefProperty: refPropID,
		Cost:        cost,
	}
}
