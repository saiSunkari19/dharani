package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Property struct {
	ID              string `json:"id"`
	UniqueID        string `json:"unique_id"`
	TotalShares     uint64 `json:"total_shares"`
	RemainingShares uint64 `json:"remaining_s_hares"`
	SoldShares      uint64 `json:"sold_shares"`
	AvailableShares uint64 `json:"available_shares"`

	Owner    sdk.AccAddress `json:"owner"`
	Location string         `json:"location"`
	Type     string         `json:"type"`

	CostAtBought sdk.Coin `json:"cost_at_bought"`
	PerSqFtCost  sdk.Coin `json:"per_sq_cost"`
}

func NewProperty(id string, shares uint64, owner sdk.AccAddress, location string,
	_type string, cost sdk.Coin) Property {
	return Property{
		ID:              id,
		TotalShares:     shares,
		RemainingShares: shares,
		SoldShares:      0,
		AvailableShares: shares,

		Owner:    owner,
		Location: location,
		Type:     _type,

		CostAtBought: sdk.Coin{},
		PerSqFtCost:  cost,
	}
}

func (p Property) NewMarketPlaceProperty(uniqueID string, shares uint64, _type string, perSqCost sdk.Coin) Property {
	return Property{
		ID:              p.ID,
		UniqueID:        uniqueID,
		TotalShares:     shares,
		RemainingShares: shares,

		Owner:    p.Owner,
		Location: p.Location,
		Type:     _type,

		PerSqFtCost: perSqCost,
	}
}
