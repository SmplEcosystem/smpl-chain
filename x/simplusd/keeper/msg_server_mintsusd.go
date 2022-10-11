package keeper

import (
	"context"

	"github.com/Smpl-Finance/smpl-chain/x/simplusd/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mintsusd(goCtx context.Context, msg *types.MsgMintsusd) (*types.MsgMintsusdResponse, error) {

	amount, err := sdk.ParseUint(msg.Amount)
	if err != nil {
		return nil, err
	}
	coin := sdk.NewCoins(sdk.Coin{Denom: "SUSD", Amount: sdk.Int(amount)})

	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coin)
	if err != nil {
		return nil, err
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, coin)

	return &types.MsgMintsusdResponse{}, nil
}
