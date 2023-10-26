package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "microservice/testutil/keeper"
	"microservice/x/msg/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MsgKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
