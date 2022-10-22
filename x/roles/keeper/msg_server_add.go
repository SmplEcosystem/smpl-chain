package keeper

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Smpl-Finance/smpl-chain/x/roles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Add(goCtx context.Context, msg *types.MsgAdd) (*types.MsgAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, _ := sdk.AccAddressFromBech32(msg.Creator)

	isAdmin, _ := k.IsAdmin(ctx, types.ModuleName, user)

	if !isAdmin {
		return nil, errors.New("only admin can do this action")
	}

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
