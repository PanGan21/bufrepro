package modb_test

import (
	"testing"

	keepertest "bufrepro/testutil/keeper"
	"bufrepro/testutil/nullify"
	"bufrepro/x/modb/module"
	"bufrepro/x/modb/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ModbKeeper(t)
	modb.InitGenesis(ctx, k, genesisState)
	got := modb.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
