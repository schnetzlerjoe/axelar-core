package types_0_17

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	vote "github.com/axelarnetwork/axelar-core/x/vote/exported_0_17"
	"github.com/axelarnetwork/axelar-core/x/vote/types"
)

// NewVoteRequest creates a message of type VoteMsgRequest
func NewVoteRequest(sender sdk.AccAddress, pollKey vote.PollKey, vote vote.Vote) *VoteRequest {
	return &VoteRequest{
		Sender:  sender,
		PollKey: pollKey,
		Vote:    vote,
	}
}

// Route implements sdk.Msg
func (m VoteRequest) Route() string {
	return types.RouterKey
}

// Type implements sdk.Msg
func (m VoteRequest) Type() string {
	return "Vote"
}

// ValidateBasic implements sdk.Msg
func (m VoteRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.PollKey.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid poll key")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (m VoteRequest) GetSignBytes() []byte {
	bz := types.ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m VoteRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
