package keeper

import (
	"github.com/tendermint/go-amino"

	"github.com/dharani/x/dharani/types"

	// this line is used by starport scaffolding
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for dharani clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryProperty:
			{
				return queryProperty(ctx, path[1:], k)
			}
		case types.QueryAllProperties:
			{
				return queryAllProperties(ctx, k)
			}
		case types.QueryPropertyByAddr:
			{
				return queryPropertyByAddress(ctx, path[1:], k)
			}
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown dharani query endpoint")
		}
	}
}

func queryProperty(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	property := k.GetProperty(ctx, path[0])
	if property == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "No property found with propertyID")
	}

	bz, err := amino.MarshalJSONIndent(property, "", " ")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryAllProperties(ctx sdk.Context, k Keeper) ([]byte, error) {
	properties := k.GetAllProperties(ctx)
	if properties == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no properties found")
	}

	bz, err := amino.MarshalJSONIndent(properties, "", " ")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryPropertyByAddress(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	addr, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "error while converting account address")
	}

	properties := k.GetPropertyByAddress(ctx, addr)
	bz, err := amino.MarshalJSONIndent(properties, "", " ")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
