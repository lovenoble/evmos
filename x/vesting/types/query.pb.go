// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/vesting/v2/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryBalancesRequest is the request type for the Query/Balances RPC method.
type QueryBalancesRequest struct {
	// address of the clawback vesting account
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryBalancesRequest) Reset()         { *m = QueryBalancesRequest{} }
func (m *QueryBalancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalancesRequest) ProtoMessage()    {}
func (*QueryBalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e31744b0ce27e85a, []int{0}
}
func (m *QueryBalancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalancesRequest.Merge(m, src)
}
func (m *QueryBalancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalancesRequest proto.InternalMessageInfo

func (m *QueryBalancesRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QueryBalancesResponse is the response type for the Query/Balances RPC
// method.
type QueryBalancesResponse struct {
	// locked defines the current amount of locked tokens
	Locked github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=locked,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"locked"`
	// unvested defines the current amount of unvested tokens
	Unvested github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=unvested,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"unvested"`
	// vested defines the current amount of vested tokens
	Vested github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=vested,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"vested"`
}

func (m *QueryBalancesResponse) Reset()         { *m = QueryBalancesResponse{} }
func (m *QueryBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalancesResponse) ProtoMessage()    {}
func (*QueryBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e31744b0ce27e85a, []int{1}
}
func (m *QueryBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalancesResponse.Merge(m, src)
}
func (m *QueryBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalancesResponse proto.InternalMessageInfo

func (m *QueryBalancesResponse) GetLocked() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Locked
	}
	return nil
}

func (m *QueryBalancesResponse) GetUnvested() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Unvested
	}
	return nil
}

func (m *QueryBalancesResponse) GetVested() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Vested
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryBalancesRequest)(nil), "evmos.vesting.v2.QueryBalancesRequest")
	proto.RegisterType((*QueryBalancesResponse)(nil), "evmos.vesting.v2.QueryBalancesResponse")
}

func init() { proto.RegisterFile("evmos/vesting/v2/query.proto", fileDescriptor_e31744b0ce27e85a) }

var fileDescriptor_e31744b0ce27e85a = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x41, 0x4f, 0xe2, 0x40,
	0x14, 0xee, 0x40, 0x96, 0x65, 0x67, 0x2f, 0x9b, 0x86, 0x4d, 0xba, 0x84, 0x14, 0x42, 0x36, 0x48,
	0x8c, 0xce, 0x40, 0x3d, 0x78, 0x47, 0xff, 0x80, 0x1c, 0xbd, 0x4d, 0xdb, 0x49, 0x6d, 0x80, 0x79,
	0x85, 0x99, 0x36, 0x12, 0xe3, 0xc5, 0x9b, 0x37, 0x13, 0xfe, 0x85, 0xbf, 0x84, 0x78, 0x22, 0xf1,
	0xe2, 0x49, 0x0d, 0xf8, 0x43, 0x4c, 0xdb, 0x81, 0x18, 0x34, 0xf1, 0xa2, 0xa7, 0xbe, 0xce, 0x7b,
	0xdf, 0xf7, 0xde, 0xf7, 0xbe, 0x87, 0x6b, 0x3c, 0x19, 0x81, 0xa4, 0x09, 0x97, 0x2a, 0x14, 0x01,
	0x4d, 0x1c, 0x3a, 0x8e, 0xf9, 0x64, 0x4a, 0xa2, 0x09, 0x28, 0x30, 0xff, 0x64, 0x59, 0xa2, 0xb3,
	0x24, 0x71, 0xaa, 0xb6, 0x07, 0x32, 0x05, 0xb8, 0x4c, 0x72, 0x9a, 0x74, 0x5d, 0xae, 0x58, 0x97,
	0x7a, 0x10, 0x8a, 0x1c, 0x51, 0xad, 0x04, 0x10, 0x40, 0x16, 0xd2, 0x34, 0xd2, 0xaf, 0xb5, 0x00,
	0x20, 0x18, 0x72, 0xca, 0xa2, 0x90, 0x32, 0x21, 0x40, 0x31, 0x15, 0x82, 0x90, 0x79, 0xb6, 0xd9,
	0xc1, 0x95, 0x93, 0xb4, 0x69, 0x8f, 0x0d, 0x99, 0xf0, 0xb8, 0xec, 0xf3, 0x71, 0xcc, 0xa5, 0x32,
	0x2d, 0xfc, 0x93, 0xf9, 0xfe, 0x84, 0x4b, 0x69, 0xa1, 0x06, 0x6a, 0xff, 0xea, 0xaf, 0x7f, 0x9b,
	0x77, 0x05, 0xfc, 0x77, 0x0b, 0x22, 0x23, 0x10, 0x92, 0x9b, 0x1e, 0x2e, 0x0d, 0xc1, 0x1b, 0x70,
	0xdf, 0x42, 0x8d, 0x62, 0xfb, 0xb7, 0xf3, 0x8f, 0xe4, 0x03, 0x93, 0x74, 0x60, 0xa2, 0x07, 0x26,
	0x47, 0x10, 0x8a, 0x5e, 0x67, 0xfe, 0x58, 0x37, 0x6e, 0x9f, 0xea, 0xed, 0x20, 0x54, 0x67, 0xb1,
	0x4b, 0x3c, 0x18, 0x51, 0xad, 0x2e, 0xff, 0xec, 0x4b, 0x7f, 0x40, 0xd5, 0x34, 0xe2, 0x32, 0x03,
	0xc8, 0xbe, 0xa6, 0x36, 0x03, 0x5c, 0x8e, 0x45, 0xba, 0x14, 0xee, 0x5b, 0x85, 0xaf, 0x6f, 0xb3,
	0x21, 0x4f, 0xd5, 0xe8, 0x36, 0xc5, 0x6f, 0x50, 0x93, 0x53, 0x3b, 0x33, 0x84, 0x7f, 0x64, 0xcb,
	0x34, 0xaf, 0x11, 0x2e, 0xaf, 0x37, 0x6a, 0xb6, 0xc8, 0xb6, 0xf9, 0xe4, 0x23, 0x97, 0xaa, 0x3b,
	0x9f, 0xd6, 0xe5, 0xd6, 0x34, 0xf7, 0xae, 0xee, 0x5f, 0x66, 0x85, 0x96, 0xf9, 0x9f, 0xbe, 0xbb,
	0x39, 0x57, 0xd7, 0xd2, 0x0b, 0xed, 0xf0, 0x65, 0xef, 0x78, 0xbe, 0xb4, 0xd1, 0x62, 0x69, 0xa3,
	0xe7, 0xa5, 0x8d, 0x6e, 0x56, 0xb6, 0xb1, 0x58, 0xd9, 0xc6, 0xc3, 0xca, 0x36, 0x4e, 0x77, 0xdf,
	0x28, 0xcc, 0x99, 0x34, 0x5f, 0xf7, 0x90, 0x9e, 0x6f, 0x58, 0x33, 0xa5, 0x6e, 0x29, 0xbb, 0xb0,
	0x83, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x6d, 0xdd, 0x49, 0xe7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Balances retrieves the unvested, vested and locked tokens for a vesting account
	Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error) {
	out := new(QueryBalancesResponse)
	err := c.cc.Invoke(ctx, "/evmos.vesting.v2.Query/Balances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Balances retrieves the unvested, vested and locked tokens for a vesting account
	Balances(context.Context, *QueryBalancesRequest) (*QueryBalancesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Balances(ctx context.Context, req *QueryBalancesRequest) (*QueryBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balances not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Balances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.vesting.v2.Query/Balances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balances(ctx, req.(*QueryBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evmos.vesting.v2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balances",
			Handler:    _Query_Balances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evmos/vesting/v2/query.proto",
}

func (m *QueryBalancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vested) > 0 {
		for iNdEx := len(m.Vested) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Vested[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Unvested) > 0 {
		for iNdEx := len(m.Unvested) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Unvested[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Locked) > 0 {
		for iNdEx := len(m.Locked) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Locked[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryBalancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Locked) > 0 {
		for _, e := range m.Locked {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Unvested) > 0 {
		for _, e := range m.Unvested {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Vested) > 0 {
		for _, e := range m.Vested {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBalancesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBalancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locked = append(m.Locked, types.Coin{})
			if err := m.Locked[len(m.Locked)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unvested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unvested = append(m.Unvested, types.Coin{})
			if err := m.Unvested[len(m.Unvested)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vested = append(m.Vested, types.Coin{})
			if err := m.Vested[len(m.Vested)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
