package rest

import (
	"errors"
	"fmt"
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	
	"github.com/dharani/x/dharani/types"
)

func getPropertyHandlerFunc(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var properties []types.Property
		
		id := r.URL.Query().Get("id")
		addr := r.URL.Query().Get("address")
		if len(addr) > 0 {
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPropertyByAddr, addr), nil)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			if len(res) == 0 {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New("response bytes are empty").Error())
				return
			}
			
			if err := cliCtx.Codec.UnmarshalJSON(res, &properties); err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New("error occurs while unmarshal").Error())
				return
			}
			
		} else if len(id) > 1 {
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryProperty, id), nil)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			if len(res) == 0 {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New("response bytes are empty").Error())
				return
			}
			
			var property types.Property
			if err := cliCtx.Codec.UnmarshalJSON(res, &property); err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New("error occurs while unmarshal").Error())
				return
			}
			properties = append(properties, property)
		} else {
			
			resKVs, _, err := cliCtx.QuerySubspace([]byte{0x01}, types.ModuleName)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			
			for _, pair := range resKVs {
				var property types.Property
				cliCtx.Codec.MustUnmarshalBinaryBare(pair.Value, &property)
				properties = append(properties, property)
			}
		}
		
		rest.PostProcessResponse(w, cliCtx, properties)
	}
}
