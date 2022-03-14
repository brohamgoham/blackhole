package blackhole_test

import (
	"testing"

	keepertest "github.com/brohamgoham/blackhole/testutil/keeper"
	"github.com/brohamgoham/blackhole/testutil/nullify"
	"github.com/brohamgoham/blackhole/x/blackhole"
	"github.com/brohamgoham/blackhole/x/blackhole/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlackholeKeeper(t)
	blackhole.InitGenesis(ctx, *k, genesisState)
	got := blackhole.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
