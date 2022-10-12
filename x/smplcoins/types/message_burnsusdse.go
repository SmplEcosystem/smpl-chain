package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnsusdse = "burnsusdse"

var _ sdk.Msg = &MsgBurnsusdse{}

func NewMsgBurnsusdse(creator string, amount string) *MsgBurnsusdse {
	return &MsgBurnsusdse{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnsusdse) Route() string {
	return RouterKey
}

func (msg *MsgBurnsusdse) Type() string {
	return TypeMsgBurnsusdse
}

func (msg *MsgBurnsusdse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnsusdse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnsusdse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
