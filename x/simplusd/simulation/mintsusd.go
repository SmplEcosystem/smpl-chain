package simulation

import (
	"math/rand"

	"github.com/Smpl-Finance/smpl-chain/x/simplusd/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/simplusd/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMintsusd(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintsusd{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Mintsusd simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Mintsusd simulation not implemented"), nil, nil
	}
}
