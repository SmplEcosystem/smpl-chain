package keeper

import (
	"github.com/Smpl-Finance/smpl-chain/x/simplusd/types"
)

var _ types.QueryServer = Keeper{}
