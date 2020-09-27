package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	types2 "github.com/dharani/types"
	"github.com/dharani/x/dharani/types"
)

func CommandBuyProperty(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy",
		Short: "to buy property value",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI(nil).WithTxEncoder(client.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			_propertyID := viper.GetString(flagPropertyID)
			_noOfSqFt := viper.GetInt64(flagNoOfSqFt)
			propertyID, err := types2.NewPropertyIDFromString(_propertyID)
			if err != nil {
				return err
			}
			
			msg := types.NewMsgBuyProperty(ctx.FromAddress, propertyID, _noOfSqFt)
			
			return client.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	
	cmd.Flags().String(flagPropertyID, "", "property-id")
	cmd.Flags().Int64(flagNoOfSqFt, 1, "no of square feet to buy ")
	
	_ = cmd.MarkFlagRequired(flagPropertyID)
	_ = cmd.MarkFlagRequired(flagNoOfSqFt)
	
	return cmd
}
