package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintsusdse = "mintsusdse"

var _ sdk.Msg = &MsgMintsusdse{}

func NewMsgMintsusdse(creator string, amount string) *MsgMintsusdse {
	return &MsgMintsusdse{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMintsusdse) Route() string {
	return RouterKey
}

func (msg *MsgMintsusdse) Type() string {
	return TypeMsgMintsusdse
}

func (msg *MsgMintsusdse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintsusdse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintsusdse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
