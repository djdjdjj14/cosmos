package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LogList: []Log{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in log
	logIdMap := make(map[uint64]bool)
	logCount := gs.GetLogCount()
	for _, elem := range gs.LogList {
		if _, ok := logIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for log")
		}
		if elem.Id >= logCount {
			return fmt.Errorf("log id should be lower or equal than the last id")
		}
		logIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
