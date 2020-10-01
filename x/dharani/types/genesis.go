package types

type GenesisState struct {
	Properties []Property
}

func NewGenesisState() GenesisState {
	return GenesisState{
	}
}
func DefaultGenesisState() GenesisState {
	return GenesisState{
	}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}
