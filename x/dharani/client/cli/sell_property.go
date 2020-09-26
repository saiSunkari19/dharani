package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	types2 "github.com/dharani/types"
	"github.com/dharani/x/dharani/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CommandSellProperty(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell",
		Short: "to sell property value",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI(nil).WithTxEncoder(client.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			_propertyID := viper.GetString(flagPropertyID)
			area := viper.GetInt64(flagArea)
			coin := viper.GetString(flagCost)
			propertyID, err := types2.NewPropertyIDFromString(_propertyID)
			if err != nil {
				return err
			}
			cost, err := sdk.ParseCoin(coin)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellProperty(ctx.FromAddress, string(propertyID), area, cost)

			return client.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagPropertyID, "", "property-id")
	cmd.Flags().Int64(flagArea, 0, "area")
	cmd.Flags().String(flagCost, "", "cost")

	_ = cmd.MarkFlagRequired(flagPropertyID)
	_ = cmd.MarkFlagRequired(flagArea)
	_ = cmd.MarkFlagRequired(flagCost)

	return cmd
}
