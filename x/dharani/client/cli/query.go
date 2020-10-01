package cli

import (
	"fmt"
	
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
		Use:   "property",
		Short: "property query sub commands",
	}
	
	dharaniQueryCmd.AddCommand(
		GetPropertyByID(cdc),
		GetAllProperties(cdc),
		GetPropertiesByAddress(cdc),
	)
	
	return dharaniQueryCmd
}

func GetPropertyByID(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id [propertyID]",
		Short: "query property by its ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			
			bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s/%s", types.QuerierRoute, types.QueryProperty, args[0]), nil)
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
