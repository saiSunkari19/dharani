package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers dharani-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/faucet", faucetHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/property", addPropertyHandlerFunc(cliCtx)).Methods("POST")
}
