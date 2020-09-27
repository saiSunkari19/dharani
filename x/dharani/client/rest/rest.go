package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers dharani-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/faucet", faucetHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/property/add", addPropertyHandlerFunc(cliCtx)).Methods("POST")
	r.HandleFunc("/property/sell", sellPropertyHandlerFunc(cliCtx)).Methods("PUT")
	r.HandleFunc("/property/buy", buyPropertyHandlerFunc(cliCtx)).Methods("PUT")
	r.HandleFunc("/property/id/{id}", getPropertyByIDHandlerFunc(cliCtx)).Methods("GET")
	r.HandleFunc("/property/address/{address}", getPropertyByAddressHandlerFunc(cliCtx)).Methods("GET")
	r.HandleFunc("/properties", getAllPropertiesHandlerFunc(cliCtx)).Methods("GET")
}
