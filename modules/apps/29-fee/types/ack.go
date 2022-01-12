package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Success implements the Acknowledgement interface. The acknowledgement is
// considered successful if the forward relayer address is empty. Otherwise it is
// considered a failed acknowledgement.
func (ack IncentivizedAcknowledgement) Success() bool {
	return ack.ForwardRelayerAddress != ""
}

// Acknowledgement implements the Acknowledgement interface. It returns the
// acknowledgement serialised using JSON.
func (ack IncentivizedAcknowledgement) Acknowledgement() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&ack))
}