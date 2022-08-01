package types

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// DefaultRelativePacketTimeoutHeight is the default packet timeout height (in blocks) relative
	// to the current block height of the counterparty chain provided by the client state. The
	// timeout is disabled when set to 0.
	DefaultRelativePacketTimeoutHeight = "0-1000"

	// DefaultRelativePacketTimeoutTimestamp is the default packet timeout timestamp (in nanoseconds)
	// relative to the current block timestamp of the counterparty chain provided by the client
	// state. The timeout is disabled when set to 0. The default is currently set to a 10 minute
	// timeout.
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// NewFungibleTokenPacketData contructs a new FungibleTokenPacketData instance
func NewFungibleTokenSwapPacketData(
	supplyToken, demandCoin sdk.Coin,
	sender, receiver, counterparty string,
) FungibleTokenSwapPacketData {
	return FungibleTokenSwapPacketData{
		SupplyToken:                 supplyToken.String(),
		DemandToken:                 demandCoin.String(),
		SendAddress:                 sender,
		AddressToReceiveDemandToken: receiver,
		CounterPartyAddress:         &counterparty,
	}
}

// ValidateBasic is used for validating the token transfer.
// NOTE: The addresses formats are not validated as the sender and recipient can have different
// formats defined by their corresponding chains that are not known to IBC.
func (ftpd FungibleTokenSwapPacketData) ValidateBasic() error {
	supplyCoin, err := sdk.ParseCoinNormalized(ftpd.SupplyToken)
	if err != nil {
		return err
	}
	if !supplyCoin.IsPositive() {
		return sdkerrors.Wrapf(ErrInvalidAmount, "amount must be strictly positive: got %d", &ftpd.SupplyToken)
	}
	demandCoin, err2 := sdk.ParseCoinNormalized(ftpd.DemandToken)
	if err2 != nil {
		return err2
	}
	if !demandCoin.IsPositive() {
		return sdkerrors.Wrapf(ErrInvalidAmount, "amount must be strictly positive: got %d", &ftpd.DemandToken)
	}
	if strings.TrimSpace(ftpd.SendAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be blank")
	}
	if strings.TrimSpace(ftpd.AddressToReceiveDemandToken) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "receiver address cannot be blank")
	}
	return nil
}

// GetBytes is a helper for serialising
func (ftpd FungibleTokenSwapPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&ftpd))
}
