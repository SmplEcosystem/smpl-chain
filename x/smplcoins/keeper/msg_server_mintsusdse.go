package keeper

import (
	"context"

	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mintsusdse(goCtx context.Context, msg *types.MsgMintsusdse) (*types.MsgMintsusdseResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := sdk.ParseUint(msg.Amount)
	if err != nil {
		return nil, err
	}

	coin := sdk.NewCoins(sdk.Coin{Denom: "USDSE", Amount: sdk.Int(amount)})

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return nil, err
	// }

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coin)
	if err != nil {
		return nil, err
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, coin)

	return &types.MsgMintsusdseResponse{}, nil
}
