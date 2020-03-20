// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contact.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"

	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CheckInRequest struct {
	Address              string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Version              *NodeVersion  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Capacity             *NodeCapacity `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Operator             *NodeOperator `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CheckInRequest) Reset()         { *m = CheckInRequest{} }
func (m *CheckInRequest) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest) ProtoMessage()    {}
func (*CheckInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}
func (m *CheckInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInRequest.Unmarshal(m, b)
}
func (m *CheckInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInRequest.Marshal(b, m, deterministic)
}
func (m *CheckInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInRequest.Merge(m, src)
}
func (m *CheckInRequest) XXX_Size() int {
	return xxx_messageInfo_CheckInRequest.Size(m)
}
func (m *CheckInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInRequest proto.InternalMessageInfo

func (m *CheckInRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CheckInRequest) GetVersion() *NodeVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *CheckInRequest) GetCapacity() *NodeCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *CheckInRequest) GetOperator() *NodeOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

type CheckInResponse struct {
	PingNodeSuccess      bool     `protobuf:"varint,1,opt,name=ping_node_success,json=pingNodeSuccess,proto3" json:"ping_node_success,omitempty"`
	PingErrorMessage     string   `protobuf:"bytes,2,opt,name=ping_error_message,json=pingErrorMessage,proto3" json:"ping_error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInResponse) Reset()         { *m = CheckInResponse{} }
func (m *CheckInResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInResponse) ProtoMessage()    {}
func (*CheckInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}
func (m *CheckInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInResponse.Unmarshal(m, b)
}
func (m *CheckInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInResponse.Marshal(b, m, deterministic)
}
func (m *CheckInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInResponse.Merge(m, src)
}
func (m *CheckInResponse) XXX_Size() int {
	return xxx_messageInfo_CheckInResponse.Size(m)
}
func (m *CheckInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInResponse proto.InternalMessageInfo

func (m *CheckInResponse) GetPingNodeSuccess() bool {
	if m != nil {
		return m.PingNodeSuccess
	}
	return false
}

func (m *CheckInResponse) GetPingErrorMessage() string {
	if m != nil {
		return m.PingErrorMessage
	}
	return ""
}

type GetTimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}
func (m *GetTimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeRequest.Unmarshal(m, b)
}
func (m *GetTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeRequest.Marshal(b, m, deterministic)
}
func (m *GetTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeRequest.Merge(m, src)
}
func (m *GetTimeRequest) XXX_Size() int {
	return xxx_messageInfo_GetTimeRequest.Size(m)
}
func (m *GetTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeRequest proto.InternalMessageInfo

type GetTimeResponse struct {
	Timestamp            time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTimeResponse) Reset()         { *m = GetTimeResponse{} }
func (m *GetTimeResponse) String() string { return proto.CompactTextString(m) }
func (*GetTimeResponse) ProtoMessage()    {}
func (*GetTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}
func (m *GetTimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeResponse.Unmarshal(m, b)
}
func (m *GetTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeResponse.Marshal(b, m, deterministic)
}
func (m *GetTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeResponse.Merge(m, src)
}
func (m *GetTimeResponse) XXX_Size() int {
	return xxx_messageInfo_GetTimeResponse.Size(m)
}
func (m *GetTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeResponse proto.InternalMessageInfo

func (m *GetTimeResponse) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type ContactPingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingRequest) Reset()         { *m = ContactPingRequest{} }
func (m *ContactPingRequest) String() string { return proto.CompactTextString(m) }
func (*ContactPingRequest) ProtoMessage()    {}
func (*ContactPingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}
func (m *ContactPingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingRequest.Unmarshal(m, b)
}
func (m *ContactPingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingRequest.Marshal(b, m, deterministic)
}
func (m *ContactPingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingRequest.Merge(m, src)
}
func (m *ContactPingRequest) XXX_Size() int {
	return xxx_messageInfo_ContactPingRequest.Size(m)
}
func (m *ContactPingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingRequest proto.InternalMessageInfo

type ContactPingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingResponse) Reset()         { *m = ContactPingResponse{} }
func (m *ContactPingResponse) String() string { return proto.CompactTextString(m) }
func (*ContactPingResponse) ProtoMessage()    {}
func (*ContactPingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{5}
}
func (m *ContactPingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingResponse.Unmarshal(m, b)
}
func (m *ContactPingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingResponse.Marshal(b, m, deterministic)
}
func (m *ContactPingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingResponse.Merge(m, src)
}
func (m *ContactPingResponse) XXX_Size() int {
	return xxx_messageInfo_ContactPingResponse.Size(m)
}
func (m *ContactPingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckInRequest)(nil), "contact.CheckInRequest")
	proto.RegisterType((*CheckInResponse)(nil), "contact.CheckInResponse")
	proto.RegisterType((*GetTimeRequest)(nil), "contact.GetTimeRequest")
	proto.RegisterType((*GetTimeResponse)(nil), "contact.GetTimeResponse")
	proto.RegisterType((*ContactPingRequest)(nil), "contact.ContactPingRequest")
	proto.RegisterType((*ContactPingResponse)(nil), "contact.ContactPingResponse")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0xc5, 0x21, 0xc2, 0x4e, 0x45, 0xe4, 0xd3, 0x80, 0xb0, 0x0c, 0x52, 0x22, 0xaf, 0x22, 0x40,
	0x8e, 0x14, 0xb6, 0x59, 0x25, 0x8a, 0x10, 0x0b, 0x20, 0x6a, 0x02, 0x0b, 0x36, 0x91, 0x63, 0x17,
	0xc6, 0x0a, 0x76, 0x7b, 0xba, 0x3b, 0x23, 0xcd, 0x76, 0x4e, 0x30, 0xb7, 0x99, 0x2b, 0xcc, 0x29,
	0x66, 0xae, 0x32, 0x72, 0x77, 0xdb, 0xf9, 0xcd, 0xce, 0xf5, 0xde, 0xab, 0x57, 0xae, 0x57, 0x0d,
	0x2f, 0x23, 0x96, 0xcb, 0x30, 0x92, 0x41, 0xc1, 0x99, 0x64, 0xc4, 0x36, 0xa5, 0x07, 0x09, 0x4b,
	0x98, 0x06, 0xbd, 0x41, 0xc2, 0x58, 0xf2, 0x1f, 0xc7, 0xaa, 0xda, 0xec, 0xfe, 0x8e, 0x65, 0x9a,
	0xa1, 0x90, 0x61, 0x56, 0x18, 0x01, 0xe4, 0x2c, 0x46, 0xfd, 0xed, 0xdf, 0x5a, 0xd0, 0x99, 0xff,
	0xc3, 0x68, 0xfb, 0x35, 0xa7, 0x78, 0xb1, 0x43, 0x21, 0x89, 0x0b, 0x76, 0x18, 0xc7, 0x1c, 0x85,
	0x70, 0xad, 0xa1, 0x35, 0x6a, 0xd1, 0xaa, 0x24, 0x1f, 0xc1, 0xbe, 0x44, 0x2e, 0x52, 0x96, 0xbb,
	0x8d, 0xa1, 0x35, 0x6a, 0x4f, 0xfa, 0x81, 0xb2, 0xfa, 0xce, 0x62, 0xfc, 0xad, 0x09, 0x5a, 0x29,
	0x48, 0x00, 0x4e, 0x14, 0x16, 0x61, 0x94, 0xca, 0x2b, 0xf7, 0xb9, 0x52, 0x93, 0xbd, 0x7a, 0x6e,
	0x18, 0x5a, 0x6b, 0x4a, 0x3d, 0x2b, 0x90, 0x87, 0x92, 0x71, 0xb7, 0x79, 0xaa, 0xff, 0x61, 0x18,
	0x5a, 0x6b, 0xfc, 0x2d, 0x74, 0xeb, 0x1f, 0x17, 0x05, 0xcb, 0x05, 0x92, 0x0f, 0xd0, 0x2f, 0xd2,
	0x3c, 0x59, 0x97, 0x6d, 0x6b, 0xb1, 0x8b, 0xa2, 0x6a, 0x07, 0x87, 0x76, 0x4b, 0xa2, 0x74, 0xfa,
	0xa9, 0x61, 0xf2, 0x09, 0x88, 0xd2, 0x22, 0xe7, 0x8c, 0xaf, 0x33, 0x14, 0x22, 0x4c, 0x50, 0xad,
	0xd5, 0xa2, 0xbd, 0x92, 0x59, 0x94, 0xc4, 0x37, 0x8d, 0xfb, 0x3d, 0xe8, 0x7c, 0x41, 0xb9, 0x4a,
	0x33, 0x34, 0x29, 0xf9, 0xbf, 0xa0, 0x5b, 0x23, 0x66, 0xfc, 0x0c, 0x5a, 0x75, 0xd4, 0x6a, 0x6c,
	0x7b, 0xe2, 0x05, 0xfa, 0x18, 0x41, 0x75, 0x8c, 0x60, 0x55, 0x29, 0x66, 0xce, 0xdd, 0xfd, 0xe0,
	0xd9, 0xcd, 0xc3, 0xc0, 0xa2, 0xfb, 0x36, 0xff, 0x35, 0x90, 0xb9, 0xbe, 0xe9, 0x32, 0xcd, 0x93,
	0x6a, 0xd8, 0x1b, 0x78, 0x75, 0x84, 0xea, 0x81, 0x93, 0x25, 0xd8, 0x06, 0x26, 0x0b, 0x70, 0x96,
	0x66, 0x43, 0xf2, 0x2e, 0xa8, 0x5e, 0xc9, 0xb9, 0x95, 0xf7, 0xfe, 0x69, 0xd2, 0x38, 0x5e, 0x5b,
	0xd0, 0x54, 0x1e, 0x53, 0xb0, 0x4d, 0xba, 0xe4, 0xed, 0xbe, 0xe3, 0xe8, 0xa1, 0x78, 0xee, 0x39,
	0x61, 0x92, 0x98, 0x82, 0x6d, 0xc2, 0x39, 0xe8, 0x3e, 0x0e, 0xf0, 0xa0, 0xfb, 0x24, 0xc7, 0x59,
	0xf3, 0x4f, 0xa3, 0xd8, 0x6c, 0x5e, 0xa8, 0xc8, 0x3e, 0x3f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x0f, 0x79, 0xa3, 0xf3, 0x02, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCContactClient interface {
	DRPCConn() drpc.Conn

	PingNode(ctx context.Context, in *ContactPingRequest) (*ContactPingResponse, error)
}

type drpcContactClient struct {
	cc drpc.Conn
}

func NewDRPCContactClient(cc drpc.Conn) DRPCContactClient {
	return &drpcContactClient{cc}
}

func (c *drpcContactClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcContactClient) PingNode(ctx context.Context, in *ContactPingRequest) (*ContactPingResponse, error) {
	out := new(ContactPingResponse)
	err := c.cc.Invoke(ctx, "/contact.Contact/PingNode", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCContactServer interface {
	PingNode(context.Context, *ContactPingRequest) (*ContactPingResponse, error)
}

type DRPCContactDescription struct{}

func (DRPCContactDescription) NumMethods() int { return 1 }

func (DRPCContactDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/contact.Contact/PingNode",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCContactServer).
					PingNode(
						ctx,
						in1.(*ContactPingRequest),
					)
			}, DRPCContactServer.PingNode, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterContact(mux drpc.Mux, impl DRPCContactServer) error {
	return mux.Register(impl, DRPCContactDescription{})
}

type DRPCContact_PingNodeStream interface {
	drpc.Stream
	SendAndClose(*ContactPingResponse) error
}

type drpcContactPingNodeStream struct {
	drpc.Stream
}

func (x *drpcContactPingNodeStream) SendAndClose(m *ContactPingResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeClient interface {
	DRPCConn() drpc.Conn

	CheckIn(ctx context.Context, in *CheckInRequest) (*CheckInResponse, error)
	GetTime(ctx context.Context, in *GetTimeRequest) (*GetTimeResponse, error)
}

type drpcNodeClient struct {
	cc drpc.Conn
}

func NewDRPCNodeClient(cc drpc.Conn) DRPCNodeClient {
	return &drpcNodeClient{cc}
}

func (c *drpcNodeClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodeClient) CheckIn(ctx context.Context, in *CheckInRequest) (*CheckInResponse, error) {
	out := new(CheckInResponse)
	err := c.cc.Invoke(ctx, "/contact.Node/CheckIn", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeClient) GetTime(ctx context.Context, in *GetTimeRequest) (*GetTimeResponse, error) {
	out := new(GetTimeResponse)
	err := c.cc.Invoke(ctx, "/contact.Node/GetTime", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCNodeServer interface {
	CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error)
	GetTime(context.Context, *GetTimeRequest) (*GetTimeResponse, error)
}

type DRPCNodeDescription struct{}

func (DRPCNodeDescription) NumMethods() int { return 2 }

func (DRPCNodeDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/contact.Node/CheckIn",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeServer).
					CheckIn(
						ctx,
						in1.(*CheckInRequest),
					)
			}, DRPCNodeServer.CheckIn, true
	case 1:
		return "/contact.Node/GetTime",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeServer).
					GetTime(
						ctx,
						in1.(*GetTimeRequest),
					)
			}, DRPCNodeServer.GetTime, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterNode(mux drpc.Mux, impl DRPCNodeServer) error {
	return mux.Register(impl, DRPCNodeDescription{})
}

type DRPCNode_CheckInStream interface {
	drpc.Stream
	SendAndClose(*CheckInResponse) error
}

type drpcNodeCheckInStream struct {
	drpc.Stream
}

func (x *drpcNodeCheckInStream) SendAndClose(m *CheckInResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNode_GetTimeStream interface {
	drpc.Stream
	SendAndClose(*GetTimeResponse) error
}

type drpcNodeGetTimeStream struct {
	drpc.Stream
}

func (x *drpcNodeGetTimeStream) SendAndClose(m *GetTimeResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---
