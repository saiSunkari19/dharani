package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Property struct {
	ID          string         `json:"id"`
	Area        uint64         `json:"area"`
	Owner       sdk.AccAddress `json:"owner"`
	Location    string         `json:"location"`
	Type        string         `json:"type"`
	RefProperty string         `json:"ref_property"`
	PerSqCost   sdk.Coin       `json:"per_sq_cost"`
}

func NewProperty(id string, area uint64, owner sdk.AccAddress, location string,
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
