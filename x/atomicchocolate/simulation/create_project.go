package simulation

import (
	"math/rand"

	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/keeper"
	"github.com/kempy007/atomic-chocolate/x/atomicchocolate/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateProject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateProject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateProject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateProject simulation not implemented"), nil, nil
	}
}
