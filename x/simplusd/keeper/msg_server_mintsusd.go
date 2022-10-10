package keeper

import (
	"context"

	"github.com/Smpl-Finance/smpl-chain/x/simplusd/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mintsusd(goCtx context.Context, msg *types.MsgMintsusd) (*types.MsgMintsusdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintsusdResponse{}, nil
}
