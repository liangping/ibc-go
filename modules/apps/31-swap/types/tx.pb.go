// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ibc/applications/swap/v1/tx.proto

package types

import (
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgTransfer defines a msg to transfer fungible tokens (i.e Coins) between
// ICS20 enabled chains. See ICS Spec here:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#data-structures
type MsgSwap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	// the tokens to be transferred
	SupplyToken *types.Coin `protobuf:"bytes,3,opt,name=supply_token,json=supplyToken,proto3" json:"supply_token,omitempty"`
	DemandToken *types.Coin `protobuf:"bytes,4,opt,name=demand_token,json=demandToken,proto3" json:"demand_token,omitempty"`
	// the sender address
	SendAddress string `protobuf:"bytes,5,opt,name=send_address,json=sendAddress,proto3" json:"send_address,omitempty"`
	// the recipient address on the destination chain
	ReceiveAddress      string `protobuf:"bytes,6,opt,name=receive_address,json=receiveAddress,proto3" json:"receive_address,omitempty"`
	CounterpartyAddress string `protobuf:"bytes,7,opt,name=counterparty_address,json=counterpartyAddress,proto3" json:"counterparty_address,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight *types1.Height `protobuf:"bytes,8,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	// Timeout timestamp in absolute nanoseconds since unix epoch.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,9,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty"`
}

func (x *MsgSwap) Reset() {
	*x = MsgSwap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_swap_v1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSwap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSwap) ProtoMessage() {}

func (x *MsgSwap) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_swap_v1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSwap.ProtoReflect.Descriptor instead.
func (*MsgSwap) Descriptor() ([]byte, []int) {
	return file_ibc_applications_swap_v1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSwap) GetSourcePort() string {
	if x != nil {
		return x.SourcePort
	}
	return ""
}

func (x *MsgSwap) GetSourceChannel() string {
	if x != nil {
		return x.SourceChannel
	}
	return ""
}

func (x *MsgSwap) GetSupplyToken() *types.Coin {
	if x != nil {
		return x.SupplyToken
	}
	return nil
}

func (x *MsgSwap) GetDemandToken() *types.Coin {
	if x != nil {
		return x.DemandToken
	}
	return nil
}

func (x *MsgSwap) GetSendAddress() string {
	if x != nil {
		return x.SendAddress
	}
	return ""
}

func (x *MsgSwap) GetReceiveAddress() string {
	if x != nil {
		return x.ReceiveAddress
	}
	return ""
}

func (x *MsgSwap) GetCounterpartyAddress() string {
	if x != nil {
		return x.CounterpartyAddress
	}
	return ""
}

func (x *MsgSwap) GetTimeoutHeight() *types1.Height {
	if x != nil {
		return x.TimeoutHeight
	}
	return nil
}

func (x *MsgSwap) GetTimeoutTimestamp() uint64 {
	if x != nil {
		return x.TimeoutTimestamp
	}
	return 0
}

// MsgSwapResponse defines the Msg/Swap response type.
type MsgSwapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSwapResponse) Reset() {
	*x = MsgSwapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_swap_v1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSwapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSwapResponse) ProtoMessage() {}

func (x *MsgSwapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_swap_v1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSwapResponse.ProtoReflect.Descriptor instead.
func (*MsgSwapResponse) Descriptor() ([]byte, []int) {
	return file_ibc_applications_swap_v1_tx_proto_rawDescGZIP(), []int{1}
}

var File_ibc_applications_swap_v1_tx_proto protoreflect.FileDescriptor

var file_ibc_applications_swap_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x69, 0x62, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x77, 0x61, 0x70,
	0x12, 0x37, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xf2, 0xde, 0x1f, 0x12, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0xf2, 0xde, 0x1f, 0x15, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x52, 0x0d, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x42, 0x0a, 0x0c, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x42, 0x0a, 0x0c, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x31, 0x0a, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x61, 0x72, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x60, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x62, 0x63,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x1d, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde, 0x1f, 0x15,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1c, 0xf2, 0xde, 0x1f, 0x18, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x52, 0x10, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a,
	0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73, 0x67,
	0x53, 0x77, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5b, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x54, 0x0a, 0x04, 0x53, 0x77, 0x61, 0x70, 0x12, 0x21, 0x2e, 0x69, 0x62,
	0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x77, 0x61, 0x70, 0x1a, 0x29,
	0x2e, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x77, 0x61,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69,
	0x62, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x33, 0x31, 0x2d, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_applications_swap_v1_tx_proto_rawDescOnce sync.Once
	file_ibc_applications_swap_v1_tx_proto_rawDescData = file_ibc_applications_swap_v1_tx_proto_rawDesc
)

func file_ibc_applications_swap_v1_tx_proto_rawDescGZIP() []byte {
	file_ibc_applications_swap_v1_tx_proto_rawDescOnce.Do(func() {
		file_ibc_applications_swap_v1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_applications_swap_v1_tx_proto_rawDescData)
	})
	return file_ibc_applications_swap_v1_tx_proto_rawDescData
}

var file_ibc_applications_swap_v1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ibc_applications_swap_v1_tx_proto_goTypes = []interface{}{
	(*MsgSwap)(nil),         // 0: ibc.applications.swap.v1.MsgSwap
	(*MsgSwapResponse)(nil), // 1: ibc.applications.swap.v1.MsgSwapResponse
	(*types.Coin)(nil),      // 2: cosmos.base.v1beta1.Coin
	(*types1.Height)(nil),   // 3: ibc.core.client.v1.Height
}
var file_ibc_applications_swap_v1_tx_proto_depIdxs = []int32{
	2, // 0: ibc.applications.swap.v1.MsgSwap.supply_token:type_name -> cosmos.base.v1beta1.Coin
	2, // 1: ibc.applications.swap.v1.MsgSwap.demand_token:type_name -> cosmos.base.v1beta1.Coin
	3, // 2: ibc.applications.swap.v1.MsgSwap.timeout_height:type_name -> ibc.core.client.v1.Height
	0, // 3: ibc.applications.swap.v1.Msg.Swap:input_type -> ibc.applications.swap.v1.MsgSwap
	1, // 4: ibc.applications.swap.v1.Msg.Swap:output_type -> ibc.applications.swap.v1.MsgSwapResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ibc_applications_swap_v1_tx_proto_init() }
func file_ibc_applications_swap_v1_tx_proto_init() {
	if File_ibc_applications_swap_v1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ibc_applications_swap_v1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSwap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ibc_applications_swap_v1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSwapResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ibc_applications_swap_v1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ibc_applications_swap_v1_tx_proto_goTypes,
		DependencyIndexes: file_ibc_applications_swap_v1_tx_proto_depIdxs,
		MessageInfos:      file_ibc_applications_swap_v1_tx_proto_msgTypes,
	}.Build()
	File_ibc_applications_swap_v1_tx_proto = out.File
	file_ibc_applications_swap_v1_tx_proto_rawDesc = nil
	file_ibc_applications_swap_v1_tx_proto_goTypes = nil
	file_ibc_applications_swap_v1_tx_proto_depIdxs = nil
}
