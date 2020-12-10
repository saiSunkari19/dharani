package dharani

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/dharani/x/dharani/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgAddProperty:
			return handlerAddProperty(ctx, k, msg)
		case types.MsgSellProperty:
			return handlerSellProperty(ctx, k, msg)
		case types.MsgUpdateMarketProperty:
			return handlerUpdateMarketProperty(ctx, k, msg)
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
		types.TypeOwn, sdk.Coin{})

	key := types.GetPropertyKey(property.Owner, []byte(id))
	k.SetProperty(ctx, key, property)
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
	key := types.GetPropertyKey(msg.From, []byte(msg.PropID.String()))
	property := k.GetProperty(ctx, key)

	if property == nil {
		return nil, sdkerrors.Wrap(ErrInvalidId, "property doesn't exist")
	}
	if !property.Owner.Equals(msg.From) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "unauthorised")
	}
	if property.RemainingShares < msg.Shares {
		return nil, sdkerrors.Wrap(ErrInvalidArea, "cannot sell property more than you have")
	}

	pc := k.GetMarketPlacePropertyCount(ctx)
	uniqueID := GetPropertyID(pc)

	sellProperty := property.NewMarketPlaceProperty(uniqueID, msg.Shares, types.TypeBought, msg.PerSqCost)

	property.RemainingShares = property.RemainingShares - msg.Shares

	k.SetProperty(ctx, key, *property)

	key = types.GetMarketPlacePropertyKey([]byte(uniqueID))

	k.SetProperty(ctx, key, sellProperty)
	k.SetMarketPalcePropertyCount(ctx, pc+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgSellProperty,
			sdk.NewAttribute(AttributeKeyFromAddress, property.Owner.String()),
			sdk.NewAttribute(AttributeKeyPropertyID, uniqueID),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handlerUpdateMarketProperty(ctx sdk.Context, k Keeper, msg types.MsgUpdateMarketProperty) (*sdk.Result, error) {
	key := types.GetMarketPlacePropertyKey(msg.PropID)
	property := k.GetProperty(ctx, key)

	if property == nil {
		return nil, sdkerrors.Wrap(ErrInvalidId, "property doesn't exist")
	}

	property.PerSqFtCost = msg.PerSqFtCost

	k.SetProperty(ctx, key, *property)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgSellProperty,
			sdk.NewAttribute(AttributeKeyFromAddress, property.Owner.String()),
			sdk.NewAttribute(AttributeKeyPropertyID, property.ID),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

}

func handlerBuyProperty(ctx sdk.Context, k Keeper, msg types.MsgBuyProperty) (*sdk.Result, error) {
	key := types.GetMarketPlacePropertyKey(msg.PropID)
	property := k.GetProperty(ctx, key)

	if property == nil {
		return nil, sdkerrors.Wrap(ErrInvalidId, "property doesn't exist")
	}

	if msg.Shares > property.RemainingShares {
		return nil, sdkerrors.Wrap(ErrInvalidArea, "buying property exceeds the limit")
	}

	amount := msg.Shares * property.PerSqFtCost.Amount.Uint64()
	deductAmount := sdk.NewInt64Coin(property.PerSqFtCost.Denom, int64(amount))

	bal := k.CoinKeeper.GetBalance(ctx, msg.From, property.PerSqFtCost.Denom)
	if bal.IsLT(deductAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "to buy property")
	}

	_, err := k.CoinKeeper.SubtractCoins(ctx, msg.From, []sdk.Coin{deductAmount})
	if err != nil {
		return nil, err
	}

	_, err = k.CoinKeeper.AddCoins(ctx, property.Owner, []sdk.Coin{deductAmount})
	if err != nil {
		return nil, err
	}

	property.RemainingShares = property.RemainingShares - msg.Shares
	k.SetProperty(ctx, key, *property)

	pc := k.GetPropertyCount(ctx)
	id := GetPropertyID(pc)

	buyProperty := types.NewProperty(id, msg.Shares, msg.From, property.Location,
		types.TypeOwn, sdk.Coin{})
	buyProperty.CostAtBought = property.PerSqFtCost

	key = types.GetPropertyKey(msg.From, []byte(id))
	k.SetProperty(ctx, key, buyProperty)
	k.SetPropertyCount(ctx, pc+1)

	key = types.GetPropertyBoughtKey([]byte(property.ID))
	k.SetProperty(ctx, key, buyProperty)

	key = types.GetPropertyKey(msg.From, []byte(property.ID))
	property = k.GetProperty(ctx, key)
	property.SoldShares = msg.Shares

	k.SetProperty(ctx, key, *property)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeMsgBuyProperty,
			sdk.NewAttribute(AttributeKeyFromAddress, buyProperty.Owner.String()),
			sdk.NewAttribute(AttributeKeyPropertyID, buyProperty.ID),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
