package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/dharani/x/dharani/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CommandAddProperty(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "to adding property value",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI(nil).WithTxEncoder(client.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			area := viper.GetInt64(flagArea)
			location := viper.GetString(flagLocation)

			msg := types.NewMsgAddProperty(ctx.FromAddress, area, location)

			return client.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().Int64(flagArea, 0, "area")
	cmd.Flags().String(flagLocation, "", "location")

	_ = cmd.MarkFlagRequired(flagArea)
	_ = cmd.MarkFlagRequired(flagLocation)

	return cmd
}
