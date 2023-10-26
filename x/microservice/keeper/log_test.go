package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "microservice/testutil/keeper"
	"microservice/testutil/nullify"
	"microservice/x/microservice/keeper"
	"microservice/x/microservice/types"
)

func createNLog(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Log {
	items := make([]types.Log, n)
	for i := range items {
		items[i].Id = keeper.AppendLog(ctx, items[i])
	}
	return items
}

func TestLogGet(t *testing.T) {
	keeper, ctx := keepertest.MicroserviceKeeper(t)
	items := createNLog(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLog(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLogRemove(t *testing.T) {
	keeper, ctx := keepertest.MicroserviceKeeper(t)
	items := createNLog(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLog(ctx, item.Id)
		_, found := keeper.GetLog(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLogGetAll(t *testing.T) {
	keeper, ctx := keepertest.MicroserviceKeeper(t)
	items := createNLog(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLog(ctx)),
	)
}

func TestLogCount(t *testing.T) {
	keeper, ctx := keepertest.MicroserviceKeeper(t)
	items := createNLog(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLogCount(ctx))
}
