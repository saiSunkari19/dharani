package types

// GenesisState - all dharani state that must be provided at genesis
type GenesisState struct {
	Properties []Property
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */) GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state, these values will be initialized but empty
	}
}

// ValidateGenesis validates the dharani genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	return nil
}

