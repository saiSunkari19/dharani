package dharani

import (
	"github.com/dharani/x/dharani/keeper"
	"github.com/dharani/x/dharani/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	
	GetPropertyID            = types.GetPropertyID
	ModuleCdc                = types.ModuleCdc
	EventTypeMsgAddProperty  = types.EventTypeMsgAddProperty
	EventTypeMsgSellProperty = types.EventTypeMsgSellProperty
	EventTypeMsgBuyProperty  = types.EventTypeMsgBuyProperty
	
	AttributeKeyFromAddress = types.AttributeKeyFromAddress
	AttributeKeyPropertyID  = types.AttributeKeyPropertyID
	PropertyKey             = types.PropertyKey
	
	ErrInvalidArea     = types.ErrInvalidShares
	ErrInvalidLocation = types.ErrInvalidLocation
	ErrInvalidType     = types.ErrInvalidType
	ErrInvalidCost     = types.ErrInvalidCost
	ErrInvalidId       = types.ErrInvalidId
	ErrInvalidField    = types.ErrInvalidField
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
