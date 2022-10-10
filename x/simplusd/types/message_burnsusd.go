package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnsusd = "burnsusd"

var _ sdk.Msg = &MsgBurnsusd{}

func NewMsgBurnsusd(creator string, amount string) *MsgBurnsusd {
	return &MsgBurnsusd{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnsusd) Route() string {
	return RouterKey
}

func (msg *MsgBurnsusd) Type() string {
	return TypeMsgBurnsusd
}

func (msg *MsgBurnsusd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnsusd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnsusd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
