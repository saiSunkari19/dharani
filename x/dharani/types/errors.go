package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidFromAddress = sdkerrors.Register(ModuleName, 1, "invalid from address")
	ErrInvalidArea        = sdkerrors.Register(ModuleName, 2, "invalid area")
	ErrInvalidLocation    = sdkerrors.Register(ModuleName, 3, "invalid location")
	ErrInvalidType        = sdkerrors.Register(ModuleName, 4, "invalid  type")
	ErrInvalidCost        = sdkerrors.Register(ModuleName, 5, "invalid  cost")
)
