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

	user, _ := sdk.AccAddressFromBech32(msg.Creator)

	isAdmin, _ := k.IsAdmin(ctx, types.ModuleName, user)

	if !isAdmin {
		return nil, errors.New("only admin can do this action")
	}
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
