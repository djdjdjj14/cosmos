package microservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"microservice/x/microservice/keeper"
	"microservice/x/microservice/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the log
	for _, elem := range genState.LogList {
		k.SetLog(ctx, elem)
	}

	// Set log count
	k.SetLogCount(ctx, genState.LogCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LogList = k.GetAllLog(ctx)
	genesis.LogCount = k.GetLogCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
