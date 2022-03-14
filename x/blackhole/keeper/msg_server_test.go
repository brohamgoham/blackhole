package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/brohamgoham/blackhole/testutil/keeper"
	"github.com/brohamgoham/blackhole/x/blackhole/keeper"
	"github.com/brohamgoham/blackhole/x/blackhole/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlackholeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
