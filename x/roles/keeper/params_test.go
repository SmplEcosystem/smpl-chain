package keeper_test

import (
	"testing"

	testkeeper "github.com/Smpl-Finance/smpl-chain/testutil/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/roles/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RolesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
