package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAdd = "add"

var _ sdk.Msg = &MsgAdd{}

func NewMsgAdd(creator string, rolename string, address string) *MsgAdd {
	return &MsgAdd{
		Creator:  creator,
		Rolename: rolename,
		Address:  address,
	}
}

func (msg *MsgAdd) Route() string {
	return RouterKey
}

func (msg *MsgAdd) Type() string {
	return TypeMsgAdd
}

func (msg *MsgAdd) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAdd) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAdd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
