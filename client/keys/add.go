package keys

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/go-bip39"
	"github.com/gorilla/mux"
)

type AddKeysReqBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Seed     string `json:"seed"`
}

const mnemonicEntropySize = 256

func AddKeysRequestHandler(clientCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req AddKeysReqBody
		body, err := ioutil.ReadAll(r.Body)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		kr, err := keyring.New("dharani", keyring.BackendTest,
			viper.GetString(flags.FlagHome), os.Stdin)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		_, err = kr.Key(req.Name)
		if err == nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "key already existed")
			return
		}

		if len(req.Seed) == 0 {
			entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			mnemonic, err := bip39.NewMnemonic(entropySeed)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			coinType := sdk.GetConfig().GetCoinType()
			hdPath := hd.CreateHDPath(coinType, 0, 0).String()

			info, err := kr.NewAccount(req.Name, mnemonic, keyring.DefaultBIP39Passphrase, hdPath, hd.Secp256k1)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			out, err := keyring.Bech32KeyOutput(info)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			out.Mnemonic = mnemonic
			jsonString, err := keys.MarshalJSON(out)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, clientCtx, jsonString)
		}
	}
}

func GetKeyRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		name := vars["name"]

		kr, err := keyring.New("dharani", keyring.BackendTest,
			viper.GetString(flags.FlagHome), os.Stdin)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		info, err := kr.Key(name)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		out, err := keyring.Bech32KeyOutput(info)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		bz, err := keys.MarshalJSON(out)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, bz)
	}
}
