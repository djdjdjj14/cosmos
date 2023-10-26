package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateLog = "create_log"
	TypeMsgUpdateLog = "update_log"
	TypeMsgDeleteLog = "delete_log"
)

var _ sdk.Msg = &MsgCreateLog{}

func NewMsgCreateLog(creator string, title string, body string, time string) *MsgCreateLog {
	return &MsgCreateLog{
		Creator: creator,
		Title:   title,
		Body:    body,
		Time:    time,
	}
}

func (msg *MsgCreateLog) Route() string {
	return RouterKey
}

func (msg *MsgCreateLog) Type() string {
	return TypeMsgCreateLog
}

func (msg *MsgCreateLog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLog{}

func NewMsgUpdateLog(creator string, id uint64, title string, body string, time string) *MsgUpdateLog {
	return &MsgUpdateLog{
		Id:      id,
		Creator: creator,
		Title:   title,
		Body:    body,
		Time:    time,
	}
}

func (msg *MsgUpdateLog) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLog) Type() string {
	return TypeMsgUpdateLog
}

func (msg *MsgUpdateLog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLog{}

func NewMsgDeleteLog(creator string, id uint64) *MsgDeleteLog {
	return &MsgDeleteLog{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteLog) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLog) Type() string {
	return TypeMsgDeleteLog
}

func (msg *MsgDeleteLog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
