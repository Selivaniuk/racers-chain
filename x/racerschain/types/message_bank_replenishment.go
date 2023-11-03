package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBankReplenishment = "bank_replenishment"

var _ sdk.Msg = &MsgBankReplenishment{}

func NewMsgBankReplenishment(creator string, amount string) *MsgBankReplenishment {
	return &MsgBankReplenishment{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBankReplenishment) Route() string {
	return RouterKey
}

func (msg *MsgBankReplenishment) Type() string {
	return TypeMsgBankReplenishment
}

func (msg *MsgBankReplenishment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBankReplenishment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBankReplenishment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
