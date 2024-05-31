package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bufrepro/testutil/keeper"
	"bufrepro/x/modb/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ModbKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
