package keeper_test

import (
	"testing"

	testkeeper "github.com/brohamgoham/blackhole/testutil/keeper"
	"github.com/brohamgoham/blackhole/x/blackhole/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlackholeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
