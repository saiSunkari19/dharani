package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func SignAndBroadcast(br rest.BaseReq, cliCtx context.CLIContext,
	mode, name, password string, msgs []sdk.Msg) ([]byte, error) {

	cdc := cliCtx.Codec
	gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(http.ResponseWriter(nil), br.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		return nil, errors.New("")
	}

	simAndExec, gas, err := flags.ParseGas(br.Gas)
	if err != nil {
		return nil, err
	}

	kr, err := keyring.New(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), os.Stdin)
	if err != nil {
		return nil, err
	}

	_, err = kr.Key(name)
	if err == nil {
		return nil, errors.New("No key found")
	}

	txBldr := auth.NewTxBuilder(
		client.GetTxEncoder(cdc), br.AccountNumber, br.Sequence, gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)
	txBldr = txBldr.WithKeybase(kr)

	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			return nil, errors.New("Error invalid gas adjustment")
		}

		txBldr, err = client.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			return nil, err
		}

		if br.Simulate {
			return nil, errors.New("Base req simulation failed")
		}
	}

	stdMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, err
	}

	stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

	stdTx, err = SignStdTxFromRest(txBldr, cliCtx, cliCtx.GetFromName(), stdTx, true, false, password)
	if err != nil {
		return nil, err
	}

	return BroadcastRest(cliCtx, cdc, stdTx, mode)

}

func SignStdTxFromRest(txBldr auth.TxBuilder, cliCtx context.CLIContext, name string, stdTx auth.StdTx,
	appendSig bool, offline bool, password string) (auth.StdTx, error) {

	var signedStdTx auth.StdTx

	info, err := txBldr.Keybase().Key(name)
	if err != nil {
		return signedStdTx, err
	}

	addr := info.GetPubKey().Address()

	if !isTxSigner(sdk.AccAddress(addr), stdTx.GetSigners()) {
		return signedStdTx, fmt.Errorf("%s: %s", "Error invalid signer", name)
	}

	ad := sdk.AccAddress(addr)
	if !offline {
		txBldr, err = populateAccountFromState(txBldr, cliCtx, ad)
		if err != nil {
			return signedStdTx, err
		}
	}

	return txBldr.SignStdTx(name, password, stdTx, appendSig)
}

func BroadcastRest(cliCtx context.CLIContext, cdc *codec.Codec, stdTx auth.StdTx, mode string) ([]byte, error) {

	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, err
	}
	cliCtx = cliCtx.WithBroadcastMode(mode)

	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}

	return PostProcessResponse(cliCtx, res)
}

func isTxSigner(user sdk.AccAddress, signers []sdk.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}

	return false
}

func populateAccountFromState(txBldr auth.TxBuilder, cliCtx context.CLIContext, addr sdk.AccAddress) (auth.TxBuilder, error) {
	accGetter := types.NewAccountRetriever(client.Codec, cliCtx)
	num, seq, err := accGetter.GetAccountNumberSequence(addr)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithAccountNumber(num).WithSequence(seq + txBldr.Sequence()), nil
}

func PostProcessResponse(cliCtx context.CLIContext, response interface{}) ([]byte, error) {
	var output []byte

	if cliCtx.Height < 0 {
		return nil, errors.New("Negative height in response")
	}

	switch response.(type) {
	case []byte:
		output = response.([]byte)

	default:
		var err error
		if cliCtx.Indent {
			output, err = cliCtx.Codec.MarshalJSONIndent(response, "", "  ")
		} else {
			output, err = cliCtx.Codec.MarshalJSON(response)
		}

		if err != nil {
			return nil, err
		}
	}

	if cliCtx.Height > 0 {
		m := make(map[string]interface{})
		err := json.Unmarshal(output, &m)
		if err != nil {
			return nil, err
		}

		m["height"] = cliCtx.Height

		if cliCtx.Indent {
			output, err = json.MarshalIndent(m, "", "  ")
		} else {
			output, err = json.Marshal(m)
		}
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}
