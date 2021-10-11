package smplchain_test

import (
	"testing"

	keepertest "github.com/Smpl-Finance/smpl-chain/testutil/keeper"
	"github.com/Smpl-Finance/smpl-chain/x/smplchain"
	"github.com/Smpl-Finance/smpl-chain/x/smplchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplchainKeeper(t)
	smplchain.InitGenesis(ctx, *k, genesisState)
	got := smplchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
