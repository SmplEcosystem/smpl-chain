package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintsusd = "mintsusd"

var _ sdk.Msg = &MsgMintsusd{}

func NewMsgMintsusd(creator string, amount string) *MsgMintsusd {
	return &MsgMintsusd{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMintsusd) Route() string {
	return RouterKey
}

func (msg *MsgMintsusd) Type() string {
	return TypeMsgMintsusd
}

func (msg *MsgMintsusd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintsusd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintsusd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
