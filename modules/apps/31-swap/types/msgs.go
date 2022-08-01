package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

// msg types
const (
	TypeMsgTransfer = "swap"
)

// NewMsgSwap creates a new MsgSwap instance
//nolint:interfacer
func NewMsgSwap(
	sourcePort, sourceChannel string,
	supply, demand sdk.Coin,
	sender, receiver string,
	counterPartyAddress string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
) *MsgSwap {
	return &MsgSwap{
		SourcePort:          sourcePort,
		SourceChannel:       sourceChannel,
		SupplyToken:         &supply,
		DemandToken:         &demand,
		SendAddress:         sender,
		ReceiveAddress:      receiver,
		CounterpartyAddress: counterPartyAddress,
		TimeoutHeight:       &timeoutHeight,
		TimeoutTimestamp:    timeoutTimestamp,
	}
}

// Route implements sdk.Msg
func (MsgSwap) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgSwap) Type() string {
	return TypeMsgTransfer
}

// ValidateBasic performs a basic check of the MsgTransfer fields.
// NOTE: timeout height or timestamp values can be 0 to disable the timeout.
// NOTE: The recipient addresses format is not validated as the format defined by
// the chain is not known to IBC.
func (msg MsgSwap) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.SourcePort); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.SourceChannel); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	if !msg.SupplyToken.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.SupplyToken.String())
	}
	if !msg.SupplyToken.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, msg.SupplyToken.String())
	}
	if !msg.DemandToken.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.DemandToken.String())
	}
	if !msg.DemandToken.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, msg.DemandToken.String())
	}
	// NOTE: sender format must be validated as it is required by the GetSigners function.
	_, err := sdk.AccAddressFromBech32(msg.SendAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	if strings.TrimSpace(msg.ReceiveAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	return nil
}

// GetSignBytes implements sdk.Msg.
func (msg MsgSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgSwap) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.SendAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
