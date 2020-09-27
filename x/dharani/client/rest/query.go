package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/dharani/x/dharani/types"
	"github.com/gorilla/mux"
	"net/http"
)

func getPropertyByIDHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		bz, _, err := ctx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s",
			types.QuerierRoute, types.QueryProperty, vars["id"]), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, bz)
		return
	}
}

func getPropertyByAddressHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		bz, _, err := ctx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s", types.QuerierRoute,
			types.QueryPropertyByAddr, vars["address"]), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var properties []types.Property
		ctx.Codec.MustUnmarshalJSON(bz, &properties)

		rest.PostProcessResponse(w, ctx, properties)
		return
	}
}

func getAllPropertiesHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := ctx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s", types.QuerierRoute, types.QueryAllProperties, nil), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var properties []types.Property
		ctx.Codec.MustUnmarshalJSON(bz, &properties)

		rest.PostProcessResponse(w, ctx, properties)
		return
	}
}
