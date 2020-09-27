package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Property struct {
	ID          string
	Area        int64
	Owner       sdk.AccAddress
	Location    string
	Type        string
	RefProperty string
	PerSqCost   sdk.Coin
}

func NewProperty(id string, area int64, owner sdk.AccAddress, location string,
	_type string, refPropID string, cost sdk.Coin) Property {
	return Property{
		ID:          id,
		Area:        area,
		Owner:       owner,
		Location:    location,
		Type:        _type,
		RefProperty: refPropID,
		PerSqCost:   cost,
	}
}
