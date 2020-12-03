package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidFromAddress = sdkerrors.Register(ModuleName, 11, "invalid from address")
	ErrInvalidShares      = sdkerrors.Register(ModuleName, 12, "invalid shares")
	ErrInvalidLocation    = sdkerrors.Register(ModuleName, 13, "invalid location")
	ErrInvalidType        = sdkerrors.Register(ModuleName, 14, "invalid type")
	ErrInvalidCost        = sdkerrors.Register(ModuleName, 15, "invalid cost")
	ErrInvalidId          = sdkerrors.Register(ModuleName, 16, "invalid id")
	ErrInvalidField       = sdkerrors.Register(ModuleName, 17, "invalid filed")
)
