package exported_0_17

import (
	fmt "fmt"

	"github.com/axelarnetwork/axelar-core/utils"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewPollKey constructor for PollKey without nonce
func NewPollKey(module string, id string) PollKey {
	return PollKey{
		Module: module,
		ID:     utils.NormalizeString(id),
	}
}

func (m PollKey) String() string {
	return fmt.Sprintf("%s_%s", m.Module, m.ID)
}

// Validate performs a stateless validity check to ensure PollKey has been properly initialized
func (m PollKey) Validate() error {
	if m.Module == "" {
		return fmt.Errorf("missing module")
	}

	if err := utils.ValidateString(m.ID, ""); err != nil {
		return sdkerrors.Wrap(err, "invalid poll key ID")
	}

	return nil
}
