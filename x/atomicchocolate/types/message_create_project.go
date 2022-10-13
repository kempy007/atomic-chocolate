package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateProject = "create_project"

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(creator string, title string, description string, literature string) *MsgCreateProject {
	return &MsgCreateProject{
		Creator:     creator,
		Title:       title,
		Description: description,
		Literature:  literature,
	}
}

func (msg *MsgCreateProject) Route() string {
	return RouterKey
}

func (msg *MsgCreateProject) Type() string {
	return TypeMsgCreateProject
}

func (msg *MsgCreateProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
