package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dharani/types"
)

var _ sdk.Msg = (*MsgAddProperty)(nil)

type MsgAddProperty struct {
	From     sdk.AccAddress
	Area     int64
	Location string
}

func (msg MsgAddProperty) Type() string {
	return "add_property"
}

func (msg MsgAddProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}
	if msg.Area < 0 {
		return ErrInvalidArea
	}
	if msg.Location == "" {
		return ErrInvalidLocation
	}

	return nil
}

func (msg MsgAddProperty) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgAddProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgAddProperty) Route() string {
	return RouterKey
}

func NewMsgAddProperty(from sdk.AccAddress, area int64, location string) *MsgAddProperty {
	return &MsgAddProperty{
		From:     from,
		Area:     area,
		Location: location,
	}
}

var _ sdk.Msg = (*MsgSellProperty)(nil)

type MsgSellProperty struct {
	From   sdk.AccAddress
	PropID types.PropertyID
	Area   int64
	Cost   sdk.Coin
}

func (msg MsgSellProperty) Type() string {
	return "sell_property"
}

func (msg MsgSellProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}
	if msg.Area < 0 {
		return ErrInvalidArea
	}
	if msg.Cost.Amount.Int64() < 0 {
		return ErrInvalidCost
	}

	return nil
}

func (msg MsgSellProperty) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgSellProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgSellProperty) Route() string {
	return RouterKey
}

func NewMsgSellProperty(from sdk.AccAddress, id types.PropertyID, area int64, cost sdk.Coin) *MsgSellProperty {
	return &MsgSellProperty{
		From:   from,
		Area:   area,
		PropID: id,
		Cost:   cost,
	}
}

var _ sdk.Msg = (*MsgBuyProperty)(nil)

type MsgBuyProperty struct {
	From   sdk.AccAddress
	PropID types.PropertyID
	Cost   sdk.Coin
}

func (msg MsgBuyProperty) Type() string {
	return "buy_property"
}

func (msg MsgBuyProperty) ValidateBasic() error {
	if msg.From == nil || msg.From.Empty() {
		return ErrInvalidFromAddress
	}
	if msg.Cost.Amount.Int64() < 0 {
		return ErrInvalidCost
	}

	return nil
}

func (msg MsgBuyProperty) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgBuyProperty) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgBuyProperty) Route() string {
	return RouterKey
}

func NewMsgBuyProperty(from sdk.AccAddress, id types.PropertyID, cost sdk.Coin) *MsgBuyProperty {
	return &MsgBuyProperty{
		From:   from,
		PropID: id,
		Cost:   cost,
	}
}
