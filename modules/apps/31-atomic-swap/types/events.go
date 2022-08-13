package types

// IBC transfer events
const (
	EventTypeTimeout      = "timeout"
	EventTypePacket       = "atomic_swap_packet"
	EventTypeTransfer     = "ibc_atomic_swap"
	EventTypeChannelClose = "channel_closed"
	EventTypeDenomTrace   = "denomination_trace"

	AttributeKeyReceiver        = "receiver"
	AttributeKeySendingDenom    = "sending_token_denom"
	AttributeKeySendingAmount   = "sending_token_amount"
	AttributeKeyReceivingDenom  = "receiving_token_denom"
	AttributeKeyReceivingAmount = "receiving_token_amount"
	AttributeKeyCounterParty    = "expected_counterparty_address"
	AttributeKeyRefundReceiver  = "refund_receiver"
	AttributeKeyRefundDenom     = "refund_denom"
	AttributeKeyRefundAmount    = "refund_amount"
	AttributeKeyAckSuccess      = "success"
	AttributeKeyAck             = "acknowledgement"
	AttributeKeyAckError        = "error"
	AttributeKeyTraceHash       = "trace_hash"
)
