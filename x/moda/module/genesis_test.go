package moda_test

import (
	"testing"

	keepertest "bufrepro/testutil/keeper"
	"bufrepro/testutil/nullify"
	"bufrepro/x/moda/module"
	"bufrepro/x/moda/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ModaKeeper(t)
	moda.InitGenesis(ctx, k, genesisState)
	got := moda.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
