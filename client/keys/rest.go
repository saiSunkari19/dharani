package keys

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/keys", AddKeysRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/keys/{name}", GetKeyRequestHandler(cliCtx)).Methods("GET")
}