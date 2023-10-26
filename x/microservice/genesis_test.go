package microservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "microservice/testutil/keeper"
	"microservice/testutil/nullify"
	"microservice/x/microservice"
	"microservice/x/microservice/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LogList: []types.Log{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LogCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MicroserviceKeeper(t)
	microservice.InitGenesis(ctx, *k, genesisState)
	got := microservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LogList, got.LogList)
	require.Equal(t, genesisState.LogCount, got.LogCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
