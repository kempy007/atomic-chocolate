package keeper_test

import (
	"testing"

	testkeeper "github.com/kempy007/atomic-chocolate/testutil/keeper"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AtomicchocolateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
