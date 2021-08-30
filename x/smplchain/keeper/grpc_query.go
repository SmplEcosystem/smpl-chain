package keeper

import (
	"github.com/Smpl-Finance/smpl-chain/x/smplchain/types"
)

var _ types.QueryServer = Keeper{}
