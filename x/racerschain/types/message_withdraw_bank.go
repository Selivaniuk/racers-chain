package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawBank = "withdraw_bank"

var _ sdk.Msg = &MsgWithdrawBank{}

func NewMsgWithdrawBank(creator string, amount string) *MsgWithdrawBank {
	return &MsgWithdrawBank{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgWithdrawBank) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawBank) Type() string {
	return TypeMsgWithdrawBank
}

func (msg *MsgWithdrawBank) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawBank) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawBank) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
