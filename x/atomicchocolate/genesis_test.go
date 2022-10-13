package atomicchocolate_test

import (
	"testing"

	keepertest "github.com/kempy007/atomic-chocolate/testutil/keeper"
	"github.com/kempy007/atomic-chocolate/testutil/nullify"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AtomicchocolateKeeper(t)
	atomicchocolate.InitGenesis(ctx, *k, genesisState)
	got := atomicchocolate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
