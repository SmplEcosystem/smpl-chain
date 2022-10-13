package keeper

import (
	"context"
	"encoding/json"

	"github.com/Smpl-Finance/smpl-chain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Add(goCtx context.Context, msg *types.MsgAdd) (*types.MsgAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var isRoleExist bool
	isRoleExist = false

	isExist := ctx.MultiStore().GetKVStore(k.storeKey).Get([]byte(msg.Address + msg.Rolename))

	if len(isExist) > 0 {
		json.Unmarshal(isExist, &isRoleExist)

	}

	if !isRoleExist {
		ctx.MultiStore().GetKVStore(k.storeKey).Set([]byte(msg.Address+msg.Rolename), []byte("true"))
	}

	return &types.MsgAddResponse{}, nil
}
