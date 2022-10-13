package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Smpl-Finance/smpl-chain/testutil/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/roles/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RolesKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
