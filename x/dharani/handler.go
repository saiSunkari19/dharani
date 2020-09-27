package dharani

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dharani/x/dharani/types"
	"strings"
)

// NewHandler ...
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgAddProperty:
			return handlerAddProperty(ctx, k, msg)
		case types.MsgSellProperty:
			return handlerSellProperty(ctx, k, msg)
		case types.MsgBuyProperty:
			return handlerBuyProperty(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handlerAddProperty(ctx sdk.Context, k Keeper, msg types.MsgAddProperty) (*sdk.Result, error) {
	pc := k.GetPropertyCount(ctx)

	id := GetPropertyID(pc)
	property := types.NewProperty(id, msg.Area, msg.From, msg.Location,
		types.TypeOwn, "", sdk.Coin{})

	k.SetProperty(ctx, id, property)
	k.SetPropertyCount(ctx, pc+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgAddProperty,
			sdk.NewAttribute(AttributeKeyFromAddress, property.Owner.String()),
			sdk.NewAttribute(AttributeKeyPropertyID, property.ID),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handlerSellProperty(ctx sdk.Context, k Keeper, msg types.MsgSellProperty) (*sdk.Result, error) {
	property := k.GetProperty(ctx, msg.PropID.String())

	if property == nil {
		return nil, errors.New("invalid property")
	}
	if !property.Owner.Equals(msg.From) {
		return nil, errors.New("invalid property owner")
	}
	if property.Area < msg.Area {
		return nil, errors.New("invalid area")
	}

	pc := k.GetPropertyCount(ctx)
	id := GetPropertyID(pc)

	sellProperty := types.NewProperty(id, property.Area, property.Owner,
		property.Location, types.TypeSell, property.ID, msg.Cost)

	property.Area = property.Area - msg.Area

	k.SetProperty(ctx, property.ID, *property)
	k.SetProperty(ctx, id, sellProperty)
	k.SetPropertyCount(ctx, pc+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgSellProperty,
			sdk.NewAttribute(AttributeKeyFromAddress, property.Owner.String()),
			sdk.NewAttribute(AttributeKeyPropertyID, id),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handlerBuyProperty(ctx sdk.Context, k Keeper, msg types.MsgBuyProperty) (*sdk.Result, error) {
	property := k.GetProperty(ctx, msg.PropID.String())

	if property == nil {
		return nil, errors.New("invalid property")
	}
	if strings.Compare(property.Type, types.TypeSell) != 0 {
		return nil, errors.New("invalid property")
	}

	bal := k.CoinKeeper.GetBalance(ctx, msg.From, property.Cost.Denom)
	if bal.IsLT(property.Cost) {
		return nil, errors.New("invalid account balance")
	}

	_, err := k.CoinKeeper.SubtractCoins(ctx, msg.From, []sdk.Coin{property.Cost})
	if err != nil {
		return nil, err
	}

	_, err = k.CoinKeeper.AddCoins(ctx, property.Owner, []sdk.Coin{property.Cost})
	if err != nil {
		return nil, err
	}

	pc := k.GetPropertyCount(ctx)
	id := GetPropertyID(pc)

	buyProperty := types.NewProperty(id, property.Area, msg.From, property.Location,
		types.TypeOwn, property.ID, sdk.Coin{})

	property.Area = 0
	k.SetProperty(ctx, property.ID, *property)
	k.SetProperty(ctx, id, buyProperty)
	k.SetPropertyCount(ctx, pc+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgBuyProperty,
			sdk.NewAttribute(AttributeKeyFromAddress, buyProperty.Owner.String()),
			sdk.NewAttribute(AttributeKeyPropertyID, buyProperty.ID),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
