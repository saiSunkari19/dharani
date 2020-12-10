package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/context"

	// "strings"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dharani/x/dharani/types"
)

func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	dharaniQueryCmd := &cobra.Command{
		Use:   "property [address] [id]",
		Short: "property query sub commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			add, _ := sdk.AccAddressFromBech32(args[0])
			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s/%s", types.QuerierRoute, types.QueryProperty, add, args[1]), nil)
			if err != nil {
				return err
			}

			var property types.Property
			cdc.MustUnmarshalJSON(bz, &property)
			return cliCtx.PrintOutput(property)
		},
	}

	return flags.GetCommands(dharaniQueryCmd)[0]
}

func GetQueryMarketPlaceCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	dharaniQueryCmd := &cobra.Command{
		Use:   "property-market-place",
		Short: "property market-place query sub commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			add, _ := sdk.AccAddressFromBech32(args[0])
			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s/%s", types.QuerierRoute, types.QueryProperty, add, args[1]), nil)
			if err != nil {
				return err
			}

			var property types.Property
			cdc.MustUnmarshalJSON(bz, &property)
			return cliCtx.PrintOutput(property)
		},
	}

	return flags.GetCommands(dharaniQueryCmd)[0]
}

func GetPropertyByID(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id [address] [propertyID] ",
		Short: "query property by its ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s", types.QuerierRoute, types.QueryProperty, args[0], args[1]), nil)
			if err != nil {
				return err
			}

			var property types.Property
			cdc.MustUnmarshalJSON(bz, &property)
			return cliCtx.PrintOutput(property)
		},
	}

	return flags.GetCommands(cmd)[0]
}

func GetAllProperties(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "query all properties",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s", types.QuerierRoute, types.QueryAllProperties, nil), nil)
			if err != nil {
				return err
			}

			var properties []types.Property
			cdc.MustUnmarshalJSON(bz, &properties)
			return cliCtx.PrintOutput(properties)
		},
	}

	return flags.GetCommands(cmd)[0]
}

func GetPropertiesByAddress(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address [address]",
		Short: "query properties by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s", types.QuerierRoute, types.QueryPropertyByAddr, args[0]), nil)
			if err != nil {
				return err
			}

			var properties []types.Property
			cdc.MustUnmarshalJSON(bz, &properties)
			return cliCtx.PrintOutput(properties)
		},
	}

	return flags.GetCommands(cmd)[0]
}
