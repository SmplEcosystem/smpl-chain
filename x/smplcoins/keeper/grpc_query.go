package keeper

import (
	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
)

var _ types.QueryServer = Keeper{}
