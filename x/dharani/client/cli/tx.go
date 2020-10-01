package cli

import (
	"github.com/spf13/cobra"
	
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "property",
		Short: "property transactions subcommands",
	}
	
	cmd.AddCommand(flags.PostCommands(
		CommandAddProperty(cdc),
		CommandSellProperty(cdc),
		CommandBuyProperty(cdc))...)
	
	return cmd
}
