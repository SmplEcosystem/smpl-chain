package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Smpl-Finance/smpl-chain/testutil/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SmplcoinsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
