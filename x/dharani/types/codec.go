package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddProperty{}, "x/dharani/MsgAddProperty", nil)
	cdc.RegisterConcrete(MsgSellProperty{}, "x/dharani/MsgSellProperty", nil)
	cdc.RegisterConcrete(MsgBuyProperty{}, "x/dharani/MsgBuyProperty", nil)
	
	cdc.RegisterConcrete(Property{}, "property", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
