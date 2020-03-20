// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodestats.proto

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

type ReputationStats struct {
	TotalCount           int64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	SuccessCount         int64    `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	ReputationAlpha      float64  `protobuf:"fixed64,3,opt,name=reputation_alpha,json=reputationAlpha,proto3" json:"reputation_alpha,omitempty"`
	ReputationBeta       float64  `protobuf:"fixed64,4,opt,name=reputation_beta,json=reputationBeta,proto3" json:"reputation_beta,omitempty"`
	ReputationScore      float64  `protobuf:"fixed64,5,opt,name=reputation_score,json=reputationScore,proto3" json:"reputation_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReputationStats) Reset()         { *m = ReputationStats{} }
func (m *ReputationStats) String() string { return proto.CompactTextString(m) }
func (*ReputationStats) ProtoMessage()    {}
func (*ReputationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{0}
}
func (m *ReputationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReputationStats.Unmarshal(m, b)
}
func (m *ReputationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReputationStats.Marshal(b, m, deterministic)
}
func (m *ReputationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReputationStats.Merge(m, src)
}
func (m *ReputationStats) XXX_Size() int {
	return xxx_messageInfo_ReputationStats.Size(m)
}
func (m *ReputationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ReputationStats.DiscardUnknown(m)
}

var xxx_messageInfo_ReputationStats proto.InternalMessageInfo

func (m *ReputationStats) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ReputationStats) GetSuccessCount() int64 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *ReputationStats) GetReputationAlpha() float64 {
	if m != nil {
		return m.ReputationAlpha
	}
	return 0
}

func (m *ReputationStats) GetReputationBeta() float64 {
	if m != nil {
		return m.ReputationBeta
	}
	return 0
}

func (m *ReputationStats) GetReputationScore() float64 {
	if m != nil {
		return m.ReputationScore
	}
	return 0
}

type GetStatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{1}
}
func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

type GetStatsResponse struct {
	UptimeCheck          *ReputationStats `protobuf:"bytes,1,opt,name=uptime_check,json=uptimeCheck,proto3" json:"uptime_check,omitempty"`
	AuditCheck           *ReputationStats `protobuf:"bytes,2,opt,name=audit_check,json=auditCheck,proto3" json:"audit_check,omitempty"`
	Disqualified         *time.Time       `protobuf:"bytes,3,opt,name=disqualified,proto3,stdtime" json:"disqualified,omitempty"`
	Suspended            *time.Time       `protobuf:"bytes,4,opt,name=suspended,proto3,stdtime" json:"suspended,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{2}
}
func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(m, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetUptimeCheck() *ReputationStats {
	if m != nil {
		return m.UptimeCheck
	}
	return nil
}

func (m *GetStatsResponse) GetAuditCheck() *ReputationStats {
	if m != nil {
		return m.AuditCheck
	}
	return nil
}

func (m *GetStatsResponse) GetDisqualified() *time.Time {
	if m != nil {
		return m.Disqualified
	}
	return nil
}

func (m *GetStatsResponse) GetSuspended() *time.Time {
	if m != nil {
		return m.Suspended
	}
	return nil
}

type DailyStorageUsageRequest struct {
	From                 time.Time `protobuf:"bytes,1,opt,name=from,proto3,stdtime" json:"from"`
	To                   time.Time `protobuf:"bytes,2,opt,name=to,proto3,stdtime" json:"to"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DailyStorageUsageRequest) Reset()         { *m = DailyStorageUsageRequest{} }
func (m *DailyStorageUsageRequest) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageRequest) ProtoMessage()    {}
func (*DailyStorageUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{3}
}
func (m *DailyStorageUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageRequest.Unmarshal(m, b)
}
func (m *DailyStorageUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageRequest.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageRequest.Merge(m, src)
}
func (m *DailyStorageUsageRequest) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageRequest.Size(m)
}
func (m *DailyStorageUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageRequest proto.InternalMessageInfo

func (m *DailyStorageUsageRequest) GetFrom() time.Time {
	if m != nil {
		return m.From
	}
	return time.Time{}
}

func (m *DailyStorageUsageRequest) GetTo() time.Time {
	if m != nil {
		return m.To
	}
	return time.Time{}
}

type DailyStorageUsageResponse struct {
	NodeId               NodeID                                    `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	DailyStorageUsage    []*DailyStorageUsageResponse_StorageUsage `protobuf:"bytes,2,rep,name=daily_storage_usage,json=dailyStorageUsage,proto3" json:"daily_storage_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *DailyStorageUsageResponse) Reset()         { *m = DailyStorageUsageResponse{} }
func (m *DailyStorageUsageResponse) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageResponse) ProtoMessage()    {}
func (*DailyStorageUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{4}
}
func (m *DailyStorageUsageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageResponse.Unmarshal(m, b)
}
func (m *DailyStorageUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageResponse.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageResponse.Merge(m, src)
}
func (m *DailyStorageUsageResponse) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageResponse.Size(m)
}
func (m *DailyStorageUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageResponse proto.InternalMessageInfo

func (m *DailyStorageUsageResponse) GetDailyStorageUsage() []*DailyStorageUsageResponse_StorageUsage {
	if m != nil {
		return m.DailyStorageUsage
	}
	return nil
}

type DailyStorageUsageResponse_StorageUsage struct {
	AtRestTotal          float64   `protobuf:"fixed64,1,opt,name=at_rest_total,json=atRestTotal,proto3" json:"at_rest_total,omitempty"`
	Timestamp            time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DailyStorageUsageResponse_StorageUsage) Reset() {
	*m = DailyStorageUsageResponse_StorageUsage{}
}
func (m *DailyStorageUsageResponse_StorageUsage) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageResponse_StorageUsage) ProtoMessage()    {}
func (*DailyStorageUsageResponse_StorageUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{4, 0}
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Unmarshal(m, b)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Merge(m, src)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Size(m)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageResponse_StorageUsage proto.InternalMessageInfo

func (m *DailyStorageUsageResponse_StorageUsage) GetAtRestTotal() float64 {
	if m != nil {
		return m.AtRestTotal
	}
	return 0
}

func (m *DailyStorageUsageResponse_StorageUsage) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*ReputationStats)(nil), "nodestats.ReputationStats")
	proto.RegisterType((*GetStatsRequest)(nil), "nodestats.GetStatsRequest")
	proto.RegisterType((*GetStatsResponse)(nil), "nodestats.GetStatsResponse")
	proto.RegisterType((*DailyStorageUsageRequest)(nil), "nodestats.DailyStorageUsageRequest")
	proto.RegisterType((*DailyStorageUsageResponse)(nil), "nodestats.DailyStorageUsageResponse")
	proto.RegisterType((*DailyStorageUsageResponse_StorageUsage)(nil), "nodestats.DailyStorageUsageResponse.StorageUsage")
}

func init() { proto.RegisterFile("nodestats.proto", fileDescriptor_e0b184ee117142aa) }

var fileDescriptor_e0b184ee117142aa = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xb1, 0x1b, 0x42, 0xf3, 0x9c, 0x36, 0xcd, 0xb0, 0x31, 0x61, 0x91, 0xc8, 0x45, 0x6a,
	0xd8, 0xa4, 0x22, 0xb0, 0x40, 0x42, 0x2c, 0x48, 0x2a, 0x41, 0x37, 0x2c, 0x9c, 0xb2, 0x61, 0x81,
	0x35, 0xf1, 0xbc, 0xb8, 0x16, 0x49, 0xc6, 0xf5, 0x3c, 0x23, 0x71, 0x05, 0x56, 0x1c, 0x00, 0x89,
	0x2b, 0x70, 0x0c, 0x4e, 0xc0, 0x82, 0x45, 0xb9, 0x0a, 0x9a, 0xb1, 0x13, 0xa7, 0xa6, 0x94, 0xb0,
	0xf4, 0xef, 0xff, 0xff, 0x3d, 0xf3, 0x3d, 0x3f, 0x68, 0x2d, 0xa5, 0x40, 0x45, 0x9c, 0xd4, 0x20,
	0x49, 0x25, 0x49, 0xd6, 0x58, 0x0b, 0x1d, 0x88, 0x64, 0x24, 0x73, 0xb9, 0xd3, 0x8d, 0xa4, 0x8c,
	0xe6, 0x78, 0x6c, 0x9e, 0xa6, 0xd9, 0xec, 0x98, 0xe2, 0x85, 0xb6, 0x2d, 0x92, 0xdc, 0xe0, 0xfd,
	0xb0, 0xa0, 0xe5, 0x63, 0x92, 0x11, 0xa7, 0x58, 0x2e, 0x27, 0xba, 0x80, 0x75, 0xc1, 0x21, 0x49,
	0x7c, 0x1e, 0x84, 0x32, 0x5b, 0x92, 0x6b, 0xf5, 0xac, 0xfe, 0x8e, 0x0f, 0x46, 0x1a, 0x6b, 0x85,
	0x1d, 0xc2, 0x9e, 0xca, 0xc2, 0x10, 0x95, 0x2a, 0x2c, 0xb6, 0xb1, 0x34, 0x0b, 0x31, 0x37, 0x3d,
	0x84, 0x83, 0x74, 0x5d, 0x1c, 0xf0, 0x79, 0x72, 0xce, 0xdd, 0x9d, 0x9e, 0xd5, 0xb7, 0xfc, 0x56,
	0xa9, 0xbf, 0xd0, 0x32, 0x3b, 0x82, 0x0d, 0x29, 0x98, 0x22, 0x71, 0xb7, 0x66, 0x9c, 0xfb, 0xa5,
	0x3c, 0x42, 0xe2, 0x95, 0x4e, 0x15, 0xca, 0x14, 0xdd, 0xdb, 0xd5, 0xce, 0x89, 0x96, 0xbd, 0x36,
	0xb4, 0x5e, 0x22, 0x99, 0x0b, 0xf9, 0x78, 0x91, 0xa1, 0x22, 0xef, 0x8b, 0x0d, 0x07, 0xa5, 0xa6,
	0x12, 0xb9, 0x54, 0xc8, 0x9e, 0x43, 0x33, 0x4b, 0x34, 0x95, 0x20, 0x3c, 0xc7, 0xf0, 0xbd, 0xb9,
	0xad, 0x33, 0xec, 0x0c, 0x4a, 0xc0, 0x15, 0x3c, 0xbe, 0x93, 0xfb, 0xc7, 0xda, 0xce, 0x9e, 0x81,
	0xc3, 0x33, 0x11, 0x53, 0x91, 0xb6, 0xff, 0x99, 0x06, 0x63, 0xcf, 0xc3, 0xaf, 0xa0, 0x29, 0x62,
	0x75, 0x91, 0xf1, 0x79, 0x3c, 0x8b, 0x51, 0x18, 0x3c, 0x3a, 0x9d, 0x0f, 0x6d, 0xb0, 0x1a, 0xda,
	0xe0, 0x6c, 0x35, 0xb4, 0xd1, 0xee, 0xf7, 0xcb, 0xae, 0xf5, 0xf9, 0x57, 0xd7, 0xf2, 0xaf, 0x24,
	0xd9, 0x08, 0x1a, 0x2a, 0x53, 0x09, 0x2e, 0x05, 0x0a, 0xc3, 0x6e, 0xdb, 0x9a, 0x32, 0xe6, 0x7d,
	0xb2, 0xc0, 0x3d, 0xe1, 0xf1, 0xfc, 0xe3, 0x84, 0x64, 0xca, 0x23, 0x7c, 0xa3, 0x78, 0x84, 0x05,
	0x3b, 0xf6, 0x14, 0x6a, 0xb3, 0x54, 0x2e, 0xd6, 0x78, 0x6e, 0xee, 0xbe, 0x65, 0xba, 0x4d, 0x82,
	0x3d, 0x01, 0x9b, 0xe4, 0x1a, 0xcc, 0x36, 0x39, 0x9b, 0xa4, 0xf7, 0xd5, 0x86, 0x7b, 0xd7, 0x1c,
	0xa6, 0x18, 0xda, 0x11, 0xdc, 0xd1, 0x84, 0x83, 0x58, 0x98, 0x03, 0x35, 0x47, 0xfb, 0x3a, 0xfc,
	0xf3, 0xb2, 0x5b, 0x7f, 0x2d, 0x05, 0x9e, 0x9e, 0xf8, 0x75, 0xfd, 0xfa, 0x54, 0x30, 0x0e, 0x77,
	0x85, 0x6e, 0x09, 0x54, 0x5e, 0x13, 0x64, 0xba, 0xc7, 0xb5, 0x7b, 0x3b, 0x7d, 0x67, 0xf8, 0x68,
	0x63, 0x4c, 0x7f, 0xfd, 0xd6, 0xe0, 0x8a, 0xd8, 0x16, 0x55, 0x5f, 0xe7, 0x03, 0x34, 0x37, 0x9f,
	0x99, 0x07, 0x7b, 0x9c, 0x82, 0x14, 0x15, 0x05, 0x66, 0x65, 0xcc, 0x09, 0x2d, 0xdf, 0xe1, 0xe4,
	0xa3, 0xa2, 0x33, 0x2d, 0xe9, 0x71, 0xad, 0x17, 0xf1, 0xbf, 0xd0, 0x94, 0xb1, 0xe1, 0x37, 0x0b,
	0x1a, 0xfa, 0xb6, 0xf9, 0xce, 0x8e, 0x61, 0x77, 0xf5, 0x6b, 0xb3, 0xcd, 0xdf, 0xaf, 0xb2, 0x03,
	0x9d, 0xfb, 0xd7, 0xbe, 0x2b, 0xb0, 0xbe, 0x83, 0xf6, 0x1f, 0x1c, 0xd8, 0xe1, 0xcd, 0x94, 0xf2,
	0xda, 0x07, 0xdb, 0xa0, 0x1c, 0xd5, 0xde, 0xda, 0xc9, 0x74, 0x5a, 0x37, 0x37, 0x7c, 0xfc, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xb9, 0xdd, 0x37, 0x5c, 0xc4, 0x04, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCNodeStatsClient interface {
	DRPCConn() drpc.Conn

	GetStats(ctx context.Context, in *GetStatsRequest) (*GetStatsResponse, error)
	DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error)
}

type drpcNodeStatsClient struct {
	cc drpc.Conn
}

func NewDRPCNodeStatsClient(cc drpc.Conn) DRPCNodeStatsClient {
	return &drpcNodeStatsClient{cc}
}

func (c *drpcNodeStatsClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodeStatsClient) GetStats(ctx context.Context, in *GetStatsRequest) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/GetStats", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeStatsClient) DailyStorageUsage(ctx context.Context, in *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error) {
	out := new(DailyStorageUsageResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/DailyStorageUsage", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCNodeStatsServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	DailyStorageUsage(context.Context, *DailyStorageUsageRequest) (*DailyStorageUsageResponse, error)
}

type DRPCNodeStatsDescription struct{}

func (DRPCNodeStatsDescription) NumMethods() int { return 2 }

func (DRPCNodeStatsDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/nodestats.NodeStats/GetStats",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeStatsServer).
					GetStats(
						ctx,
						in1.(*GetStatsRequest),
					)
			}, DRPCNodeStatsServer.GetStats, true
	case 1:
		return "/nodestats.NodeStats/DailyStorageUsage",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeStatsServer).
					DailyStorageUsage(
						ctx,
						in1.(*DailyStorageUsageRequest),
					)
			}, DRPCNodeStatsServer.DailyStorageUsage, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterNodeStats(mux drpc.Mux, impl DRPCNodeStatsServer) error {
	return mux.Register(impl, DRPCNodeStatsDescription{})
}

type DRPCNodeStats_GetStatsStream interface {
	drpc.Stream
	SendAndClose(*GetStatsResponse) error
}

type drpcNodeStatsGetStatsStream struct {
	drpc.Stream
}

func (x *drpcNodeStatsGetStatsStream) SendAndClose(m *GetStatsResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeStats_DailyStorageUsageStream interface {
	drpc.Stream
	SendAndClose(*DailyStorageUsageResponse) error
}

type drpcNodeStatsDailyStorageUsageStream struct {
	drpc.Stream
}

func (x *drpcNodeStatsDailyStorageUsageStream) SendAndClose(m *DailyStorageUsageResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---
