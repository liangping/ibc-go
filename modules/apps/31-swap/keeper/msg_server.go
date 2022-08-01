package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v4/modules/apps/31-swap/types"
)

var _ types.MsgServer = Keeper{}

// See createOutgoingPacket in spec:https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#packet-relay

// Transfer defines a rpc handler method for MsgTransfer.
func (k Keeper) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.SendAddress)
	if err != nil {
		return nil, err
	}
	if err := k.SendSwap(
		ctx, msg.SourcePort, msg.SourceChannel,
		*msg.SupplyToken, *msg.DemandToken, sender, msg.ReceiveAddress, msg.CounterpartyAddress,
		*msg.TimeoutHeight, msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	k.Logger(ctx).Info("IBC fungible token swap", "token", msg.SupplyToken.Denom, "amount", msg.SupplyToken.Amount,
		" <=> token", msg.DemandToken.Denom, "amount", msg.DemandToken.Amount,
		"sender", msg.SendAddress, "receiver", msg.ReceiveAddress, "counterpaty: ", msg.CounterpartyAddress)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSwap,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SendAddress),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.ReceiveAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgSwapResponse{}, nil
}
