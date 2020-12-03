package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dharani/types"
)

var _ sdk.Msg = (*MsgAddProperty)(nil)

type MsgAddProperty struct {
	From     sdk.AccAddress `json:"from"`
	Area     uint64         `json:"area"`
	Location string         `json:"location"`
}

func (msg MsgAddProperty) Type() string {
	return "add_property"
}

func (msg MsgAddProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}
	if msg.Area <= 0 {
		return ErrInvalidShares
	}
	if msg.Location == "" {
		return ErrInvalidLocation
	}

	return nil
}

func (msg MsgAddProperty) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAddProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgAddProperty) Route() string {
	return RouterKey
}

func NewMsgAddProperty(from sdk.AccAddress, area uint64, location string) *MsgAddProperty {
	return &MsgAddProperty{
		From:     from,
		Area:     area,
		Location: location,
	}
}

var _ sdk.Msg = (*MsgSellProperty)(nil)

type MsgSellProperty struct {
	From      sdk.AccAddress   `json:"from"`
	PropID    types.PropertyID `json:"prop_id"`
	Shares    uint64           `json:"shares"`
	PerSqCost sdk.Coin         `json:"per_sq_cost"`
}

func (msg MsgSellProperty) Type() string {
	return "sell_property"
}

func (msg MsgSellProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}
	if msg.Shares <= 0 {
		return ErrInvalidShares
	}
	if msg.PerSqCost.Amount.Int64() < 0 {
		return ErrInvalidCost
	}

	return nil
}

func (msg MsgSellProperty) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSellProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgSellProperty) Route() string {
	return RouterKey
}

func NewMsgSellProperty(from sdk.AccAddress, id types.PropertyID, area uint64, cost sdk.Coin) *MsgSellProperty {
	return &MsgSellProperty{
		From:      from,
		Shares:    area,
		PropID:    id,
		PerSqCost: cost,
	}
}

var _ sdk.Msg = (*MsgUpdateMarketProperty)(nil)

type MsgUpdateMarketProperty struct {
	From        sdk.AccAddress   `json:"from"`
	PropID      types.PropertyID `json:"prop_id"`
	PerSqFtCost sdk.Coin         `json:"per_sq_ft_cost"`
}

func (msg MsgUpdateMarketProperty) Type() string {
	return "update_market_property"
}

func (msg MsgUpdateMarketProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}

	if msg.PropID.String() != "" {
		return ErrInvalidField
	}
	if !msg.PerSqFtCost.IsValid() {
		return ErrInvalidField
	}

	return nil
}

func (msg MsgUpdateMarketProperty) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgUpdateMarketProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgUpdateMarketProperty) Route() string {
	return RouterKey
}

func NewMsgUpdateMarketProperty(from sdk.AccAddress, id types.PropertyID, cost sdk.Coin) *MsgUpdateMarketProperty {
	return &MsgUpdateMarketProperty{
		From:        from,
		PropID:      id,
		PerSqFtCost: cost,
	}
}

var _ sdk.Msg = (*MsgBuyProperty)(nil)

type MsgBuyProperty struct {
	From   sdk.AccAddress   `json:"from"`
	PropID types.PropertyID `json:"prop_id"`
	Shares uint64           `json:"shares"`
}

func (msg MsgBuyProperty) Type() string {
	return "buy_property"
}

func (msg MsgBuyProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}

	if msg.Shares <= 0 {
		return ErrInvalidField
	}

	return nil
}

func (msg MsgBuyProperty) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBuyProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgBuyProperty) Route() string {
	return RouterKey
}

func NewMsgBuyProperty(from sdk.AccAddress, id types.PropertyID, area uint64) *MsgBuyProperty {
	return &MsgBuyProperty{
		From:   from,
		PropID: id,
		Shares: area,
	}
}
