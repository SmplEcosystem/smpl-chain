package keeper

import (
	"context"
	"errors"

	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mintsusdse(goCtx context.Context, msg *types.MsgMintsusdse) (*types.MsgMintsusdseResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := sdk.ParseUint(msg.Amount)
	if err != nil {
		return nil, err
	}
	coin := sdk.NewCoins(sdk.Coin{Denom: "usdse", Amount: sdk.Int(amount)})

	user, err := sdk.AccAddressFromBech32(msg.Creator)

	isBanker, err := k.rolesKeeper.HasRole(ctx, types.ModuleName, user, "bank")
	if !isBanker {
		return &types.MsgMintsusdseResponse{}, errors.New("bank role missing")
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coin)
	if err != nil {
		return nil, err
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, coin)

	return &types.MsgMintsusdseResponse{}, nil
}
