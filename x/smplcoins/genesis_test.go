package smplcoins_test

import (
	"testing"

	keepertest "github.com/Smpl-Finance/smpl-chain/testutil/keeper"
	"github.com/Smpl-Finance/smpl-chain/testutil/nullify"
	"github.com/Smpl-Finance/smpl-chain/x/smplcoins"
	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplcoinsKeeper(t)
	smplcoins.InitGenesis(ctx, *k, genesisState)
	got := smplcoins.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
