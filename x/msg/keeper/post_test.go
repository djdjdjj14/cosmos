package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "microservice/testutil/keeper"
	"microservice/testutil/nullify"
	"microservice/x/msg/keeper"
	"microservice/x/msg/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Post {
	items := make([]types.Post, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPost(ctx, items[i])
	}
	return items
}

func TestPostGet(t *testing.T) {
	keeper, ctx := keepertest.MsgKeeper(t)
	items := createNPost(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPost(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPostRemove(t *testing.T) {
	keeper, ctx := keepertest.MsgKeeper(t)
	items := createNPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePost(ctx,
			item.Index,
		)
		_, found := keeper.GetPost(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.MsgKeeper(t)
	items := createNPost(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPost(ctx)),
	)
}
