package keeper

import (
	"context"
	"errors"

	"github.com/Smpl-Finance/smpl-chain/x/smplcoins/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burnsusdse(goCtx context.Context, msg *types.MsgBurnsusdse) (*types.MsgBurnsusdseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	amount, err := sdk.ParseUint(msg.Amount)
	if err != nil {
		return nil, err
	}

	coin := sdk.NewCoins(sdk.Coin{Denom: "USDSE", Amount: sdk.Int(amount)})

	if k.bankKeeper.SpendableCoins(ctx, user).IsAllGTE(coin) {
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, coin)
		err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coin)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("not enough balance")
	}
	_ = ctx

	return &types.MsgBurnsusdseResponse{}, nil
}
