package msg_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "microservice/testutil/keeper"
	"microservice/testutil/nullify"
	"microservice/x/msg"
	"microservice/x/msg/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PostList: []types.Post{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MsgKeeper(t)
	msg.InitGenesis(ctx, *k, genesisState)
	got := msg.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	// this line is used by starport scaffolding # genesis/test/assert
}
