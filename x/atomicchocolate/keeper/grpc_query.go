package keeper

import (
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
)

var _ types.QueryServer = Keeper{}
