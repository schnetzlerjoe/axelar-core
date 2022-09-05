package types_0_17

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/axelarnetwork/axelar-core/utils"
	types "github.com/axelarnetwork/axelar-core/x/evm/types"
)

// NewConfirmChainRequest creates a message of type ConfirmTokenRequest
func NewConfirmChainRequest(sender sdk.AccAddress, name string) *ConfirmChainRequest {
	return &ConfirmChainRequest{
		Sender: sender,
		Name:   utils.NormalizeString(name),
	}
}

// Route implements sdk.Msg
func (m ConfirmChainRequest) Route() string {
	return types.RouterKey
}

// Type implements sdk.Msg
func (m ConfirmChainRequest) Type() string {
	return "ConfirmChain"
}

// ValidateBasic implements sdk.Msg
func (m ConfirmChainRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := utils.ValidateString(m.Name); err != nil {
		return sdkerrors.Wrap(err, "invalid chain name")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (m ConfirmChainRequest) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m ConfirmChainRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
