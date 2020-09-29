package rest

import (
	"net/http"
	"os"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/viper"
)

func BuildSignBroadcast(w http.ResponseWriter, br rest.BaseReq, cliCtx context.CLIContext,
	mode, name, password string, msgs []sdk.Msg) {
	
	gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(http.ResponseWriter(nil), br.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		rest.WriteErrorResponse(w, http.StatusBadRequest, "gas adjustment failed")
		return
	}
	
	simAndExec, gas, err := flags.ParseGas(br.Gas)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	kr, err := keyring.New("dharani", keyring.BackendOS, app.DefaultNodeHome,, os.Stdin)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	_, err = kr.Key(name)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	cliCtx = cliCtx.WithFromName(name)
	cliCtx = cliCtx.WithBroadcastMode(mode)
	
	addr, err := sdk.AccAddressFromBech32(br.From)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	cliCtx = cliCtx.WithFromAddress(addr)
	
	account, err := auth.NewAccountRetriever(client.Codec, cliCtx).GetAccount(cliCtx.GetFromAddress())
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	txBldr := auth.NewTxBuilder(
		client.GetTxEncoder(cliCtx.Codec), account.GetAccountNumber(), account.GetSequence(), gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)
	txBldr = txBldr.WithKeybase(kr)
	
	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "invalid gas adjustment")
			return
		}
		
		txBldr, err = client.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "invalid gas adjustment")
			return
		}
		
		if br.Simulate {
			rest.WriteSimulationResponse(w, cliCtx.Codec, txBldr.Gas())
			return
		}
	}
	
	txBytes, err := txBldr.BuildAndSign(cliCtx.GetFromName(), password, msgs)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		
	}
	
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	rest.PostProcessResponseBare(w, cliCtx, res)
}
