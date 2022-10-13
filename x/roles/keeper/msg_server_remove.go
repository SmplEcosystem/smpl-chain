package keeper

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Smpl-Finance/smpl-chain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Remove(goCtx context.Context, msg *types.MsgRemove) (*types.MsgRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var isRoleExist bool
	isRoleExist = false

	isExist := ctx.MultiStore().GetKVStore(k.storeKey).Get([]byte(msg.Address + msg.Rolename))

	if len(isExist) > 0 {
		json.Unmarshal(isExist, &isRoleExist)

	}

	if !isRoleExist {
		return nil, errors.New("role does not exists")
	}
	ctx.MultiStore().GetKVStore(k.storeKey).Delete([]byte(msg.Address + msg.Rolename))

	return &types.MsgRemoveResponse{}, nil
}
