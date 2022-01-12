// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Fee implements the ics29 Fee interface
// See Fee Payment Middleware spec:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-029-fee-payment#fee-middleware-contract
type Fee struct {
	ReceiveFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=receive_fee,json=receiveFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"receive_fee" yaml:"receive_fee"`
	AckFee     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=ack_fee,json=ackFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"ack_fee" yaml:"ack_fee"`
	TimeoutFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=timeout_fee,json=timeoutFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"timeout_fee" yaml:"timeout_fee"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetReceiveFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ReceiveFee
	}
	return nil
}

func (m *Fee) GetAckFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AckFee
	}
	return nil
}

func (m *Fee) GetTimeoutFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TimeoutFee
	}
	return nil
}

// IdentifiedPacketFee contains the relayer fee along with the associated metadata needed to process it.
// This includes the PacketId identifying the packet the fee is paying for,
// the refund address to which any unused funds are refunded,
// and an optional list of relayers that are permitted to receive the fee.
type IdentifiedPacketFee struct {
	PacketId      *types1.PacketId `protobuf:"bytes,1,opt,name=packet_id,json=packetId,proto3" json:"packet_id,omitempty" yaml:"packet_id"`
	Fee           Fee              `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee"`
	RefundAddress string           `protobuf:"bytes,3,opt,name=refund_address,json=refundAddress,proto3" json:"refund_address,omitempty" yaml:"refund_address"`
	Relayers      []string         `protobuf:"bytes,4,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *IdentifiedPacketFee) Reset()         { *m = IdentifiedPacketFee{} }
func (m *IdentifiedPacketFee) String() string { return proto.CompactTextString(m) }
func (*IdentifiedPacketFee) ProtoMessage()    {}
func (*IdentifiedPacketFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{1}
}
func (m *IdentifiedPacketFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedPacketFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedPacketFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedPacketFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedPacketFee.Merge(m, src)
}
func (m *IdentifiedPacketFee) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedPacketFee) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedPacketFee.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedPacketFee proto.InternalMessageInfo

func (m *IdentifiedPacketFee) GetPacketId() *types1.PacketId {
	if m != nil {
		return m.PacketId
	}
	return nil
}

func (m *IdentifiedPacketFee) GetFee() Fee {
	if m != nil {
		return m.Fee
	}
	return Fee{}
}

func (m *IdentifiedPacketFee) GetRefundAddress() string {
	if m != nil {
		return m.RefundAddress
	}
	return ""
}

func (m *IdentifiedPacketFee) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

func init() {
	proto.RegisterType((*Fee)(nil), "ibc.applications.fee.v1.Fee")
	proto.RegisterType((*IdentifiedPacketFee)(nil), "ibc.applications.fee.v1.IdentifiedPacketFee")
}

func init() { proto.RegisterFile("ibc/applications/fee/v1/fee.proto", fileDescriptor_cb3319f1af2a53e5) }

var fileDescriptor_cb3319f1af2a53e5 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0x72, 0x3a, 0xae, 0xae, 0x38, 0xa1, 0x70, 0x88, 0x5e, 0x05, 0x69, 0xc9, 0x94,
	0xa5, 0xb6, 0x5a, 0x58, 0x60, 0x82, 0x20, 0x55, 0x3a, 0x89, 0xe1, 0x94, 0x91, 0xa5, 0x72, 0xec,
	0xd7, 0x9e, 0xd5, 0x24, 0x8e, 0xe2, 0x34, 0x52, 0x57, 0x16, 0x56, 0x3e, 0x07, 0x9f, 0xe4, 0xc6,
	0x1b, 0x99, 0x0a, 0x6a, 0xbf, 0x41, 0x57, 0x16, 0xe4, 0xd8, 0x3d, 0x55, 0x42, 0xe8, 0xd4, 0xc9,
	0xf6, 0x7b, 0xef, 0xef, 0xdf, 0x7b, 0xcf, 0xcf, 0xe8, 0xb5, 0x48, 0x18, 0xa1, 0x45, 0x91, 0x0a,
	0x46, 0x2b, 0x21, 0x73, 0x45, 0x66, 0x00, 0xa4, 0x1e, 0xe9, 0x05, 0x17, 0xa5, 0xac, 0xa4, 0xf7,
	0x42, 0x24, 0x0c, 0x1f, 0x86, 0x60, 0xed, 0xab, 0x47, 0x3d, 0x9f, 0x49, 0x95, 0x49, 0x45, 0x12,
	0xaa, 0xb4, 0x24, 0x81, 0x8a, 0x8e, 0x08, 0x93, 0x22, 0x37, 0xc2, 0xde, 0xc5, 0x5c, 0xce, 0x65,
	0xb3, 0x25, 0x7a, 0x67, 0xad, 0x0d, 0x91, 0xc9, 0x12, 0x08, 0xbb, 0xa1, 0x79, 0x0e, 0xa9, 0xa6,
	0xd9, 0xad, 0x09, 0x09, 0xbe, 0xb9, 0xc8, 0x9d, 0x00, 0x78, 0x5f, 0x1d, 0xd4, 0x29, 0x81, 0x81,
	0xa8, 0x61, 0x3a, 0x03, 0xe8, 0x3a, 0x03, 0x37, 0xec, 0x8c, 0x2f, 0xb1, 0xe1, 0x62, 0xcd, 0xc5,
	0x96, 0x8b, 0x3f, 0x49, 0x91, 0x47, 0x93, 0xdb, 0x75, 0xbf, 0xb5, 0x5b, 0xf7, 0xbd, 0x15, 0xcd,
	0xd2, 0xf7, 0xc1, 0x81, 0x36, 0xf8, 0xf1, 0xab, 0x1f, 0xce, 0x45, 0x75, 0xb3, 0x4c, 0x30, 0x93,
	0x19, 0xb1, 0xa9, 0x9b, 0x65, 0xa8, 0xf8, 0x82, 0x54, 0xab, 0x02, 0x54, 0x73, 0x8d, 0x8a, 0x91,
	0x55, 0xea, 0x24, 0x6a, 0xf4, 0x98, 0xb2, 0x45, 0xc3, 0x7f, 0xf4, 0x10, 0x3f, 0xb2, 0xfc, 0x73,
	0xc3, 0xb7, 0xba, 0xe3, 0xd8, 0xa7, 0x94, 0x2d, 0xf6, 0xc5, 0x57, 0x22, 0x03, 0xb9, 0xac, 0x1a,
	0xb8, 0x7b, 0x64, 0xf1, 0x07, 0xda, 0x23, 0x8b, 0xb7, 0xca, 0x09, 0x40, 0xf0, 0xc7, 0x41, 0xcf,
	0xae, 0x38, 0xe4, 0x95, 0x98, 0x09, 0xe0, 0xd7, 0x94, 0x2d, 0x40, 0xdb, 0xbd, 0x6b, 0xd4, 0x2e,
	0x9a, 0xc3, 0x54, 0xf0, 0xae, 0x33, 0x70, 0xc2, 0xce, 0xf8, 0x15, 0xd6, 0x73, 0xa2, 0x1f, 0x16,
	0xef, 0x5f, 0xb3, 0x1e, 0x61, 0x23, 0xb9, 0xe2, 0xd1, 0xc5, 0x6e, 0xdd, 0x7f, 0x6a, 0x32, 0xbb,
	0x57, 0x06, 0xf1, 0x59, 0x61, 0xfd, 0xde, 0x5b, 0xe4, 0x9a, 0x16, 0xeb, 0xbb, 0x5e, 0xe2, 0xff,
	0xcc, 0x1c, 0x9e, 0x00, 0x44, 0x27, 0xba, 0xd0, 0x58, 0x87, 0x7b, 0x1f, 0xd0, 0x79, 0x09, 0xb3,
	0x65, 0xce, 0xa7, 0x94, 0xf3, 0x12, 0x94, 0xea, 0xba, 0x03, 0x27, 0x6c, 0x47, 0x97, 0xbb, 0x75,
	0xff, 0xf9, 0x7e, 0x08, 0x0e, 0xfd, 0x41, 0xfc, 0xc4, 0x18, 0x3e, 0x9a, 0xb3, 0xd7, 0x43, 0x67,
	0x25, 0xa4, 0x74, 0x05, 0xa5, 0xea, 0x9e, 0x0c, 0xdc, 0xb0, 0x1d, 0xdf, 0x9f, 0xa3, 0xcf, 0xb7,
	0x1b, 0xdf, 0xb9, 0xdb, 0xf8, 0xce, 0xef, 0x8d, 0xef, 0x7c, 0xdf, 0xfa, 0xad, 0xbb, 0xad, 0xdf,
	0xfa, 0xb9, 0xf5, 0x5b, 0x5f, 0xc6, 0xff, 0x76, 0x53, 0x24, 0x6c, 0x38, 0x97, 0x24, 0x93, 0x7c,
	0x99, 0x82, 0xd2, 0x7f, 0x4a, 0x91, 0xf1, 0xbb, 0xa1, 0xfe, 0x4e, 0x4d, 0x77, 0x93, 0xd3, 0x66,
	0xb8, 0xdf, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x7c, 0x4e, 0x14, 0x73, 0x03, 0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TimeoutFee) > 0 {
		for iNdEx := len(m.TimeoutFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AckFee) > 0 {
		for iNdEx := len(m.AckFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ReceiveFee) > 0 {
		for iNdEx := len(m.ReceiveFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReceiveFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedPacketFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedPacketFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedPacketFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintFee(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RefundAddress) > 0 {
		i -= len(m.RefundAddress)
		copy(dAtA[i:], m.RefundAddress)
		i = encodeVarintFee(dAtA, i, uint64(len(m.RefundAddress)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PacketId != nil {
		{
			size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFee(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ReceiveFee) > 0 {
		for _, e := range m.ReceiveFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.AckFee) > 0 {
		for _, e := range m.AckFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.TimeoutFee) > 0 {
		for _, e := range m.TimeoutFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *IdentifiedPacketFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketId != nil {
		l = m.PacketId.Size()
		n += 1 + l + sovFee(uint64(l))
	}
	l = m.Fee.Size()
	n += 1 + l + sovFee(uint64(l))
	l = len(m.RefundAddress)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiveFee = append(m.ReceiveFee, types.Coin{})
			if err := m.ReceiveFee[len(m.ReceiveFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckFee = append(m.AckFee, types.Coin{})
			if err := m.AckFee[len(m.AckFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutFee = append(m.TimeoutFee, types.Coin{})
			if err := m.TimeoutFee[len(m.TimeoutFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IdentifiedPacketFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IdentifiedPacketFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedPacketFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PacketId == nil {
				m.PacketId = &types1.PacketId{}
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFee
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFee
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)