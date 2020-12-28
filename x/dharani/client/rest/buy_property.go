package rest

import (
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
	Area     uint64       `json:"area"`
	Name     string       `json:"name"`
	Password string       `json:"password"`
	Mode     string       `json:"mode"`
}

func buyPropertyHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgBuyProperty

		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "error while validating req body")
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "error while validating req body")
			return
		}
		if req.Area <= 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "invalid area")
			return
		}
		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgBuyProperty(fromAddress, req.ID, req.Area)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest2.BuildSignBroadcast(w, req.BaseReq, ctx, req.Mode, req.Name, req.Password, []sdk.Msg{msg})
	}
}
