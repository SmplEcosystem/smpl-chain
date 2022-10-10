package keeper

import (
	"context"
	"fmt"

	"github.com/Smpl-Finance/smpl-chain/x/simplusd/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burnsusd(goCtx context.Context, msg *types.MsgBurnsusd) (*types.MsgBurnsusdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Println("---------")
	// TODO: Handling the message
	_ = ctx

	return &types.MsgBurnsusdResponse{}, nil
}
