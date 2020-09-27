package rest

import (
	types2 "github.com/dharani/types"
	"net/http"

	rest2 "github.com/dharani/client/rest"
	"github.com/dharani/x/dharani/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type msgBuyProperty struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	ID       string       `json:"id"`
	Area     int64        `json:"area"`
	Name     string       `json:"name"`
	Password string       `json:"password"`
	Mode     string       `json:"mode"`
}

func buyPropertyHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgBuyProperty

		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := types2.NewPropertyIDFromString(req.ID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		msg := types.NewMsgBuyProperty(fromAddress, id, req.Area)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := rest2.SignAndBroadcast(req.BaseReq, ctx, req.Mode, req.Name, req.Password, []sdk.Msg{msg})

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
		return
	}
}
