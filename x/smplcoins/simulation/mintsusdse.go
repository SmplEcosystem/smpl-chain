package simulation

import (
	"math/rand"

	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMintsusdse(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintsusdse{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Mintsusdse simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Mintsusdse simulation not implemented"), nil, nil
	}
}
