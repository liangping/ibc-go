package types

// IBC transfer events
const (
	EventTypeTimeout      = "timeout"
	EventTypePacket       = "fungible_token_packet"
	EventTypeSwap         = "ibc_swap"
	EventTypeChannelClose = "channel_closed"
	EventTypeDenomTrace   = "denomination_trace"

	AttributeKeyReceiver     = "receiver"
	AttributeKeyDenom        = "denom"
	AttributeKeyAmount       = "amount"
	AttributeKeyCounterParty = "counterparty"
	AttributeKeyAckSuccess   = "success"
	AttributeKeyAck          = "acknowledgement"
	AttributeKeyAckError     = "error"
)
