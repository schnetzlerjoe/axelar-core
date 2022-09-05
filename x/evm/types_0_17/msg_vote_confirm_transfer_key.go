package types_0_17

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/axelarnetwork/axelar-core/utils"
	types "github.com/axelarnetwork/axelar-core/x/evm/types"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported_0_17"
)

// NewVoteConfirmTransferKeyRequest creates a message of type VoteConfirmTransferKeyRequest
func NewVoteConfirmTransferKeyRequest(
	sender sdk.AccAddress,
	chain string,
	key vote.PollKey,
	confirmed bool) *VoteConfirmTransferKeyRequest {
	return &VoteConfirmTransferKeyRequest{
		Sender:    sender,
		Chain:     utils.NormalizeString(chain),
		PollKey:   key,
		Confirmed: confirmed,
	}
}

// Route returns the route for this message
func (m VoteConfirmTransferKeyRequest) Route() string {
	return types.RouterKey
}

// Type returns the type of the message
func (m VoteConfirmTransferKeyRequest) Type() string {
	return "VoteConfirmTransferKey"
}

// ValidateBasic executes a stateless message validation
func (m VoteConfirmTransferKeyRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := utils.ValidateString(m.Chain); err != nil {
		return sdkerrors.Wrap(err, "invalid chain")
	}

	if err := m.PollKey.Validate(); err != nil {
		return err
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m VoteConfirmTransferKeyRequest) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m VoteConfirmTransferKeyRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
