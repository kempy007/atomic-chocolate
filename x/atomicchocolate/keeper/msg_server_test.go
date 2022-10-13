package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/kempy007/atomic-chocolate/testutil/keeper"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/keeper"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AtomicchocolateKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
