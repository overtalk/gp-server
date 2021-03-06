// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/paper.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ManualModifyPaperReq_ModifyType int32

const (
	ManualModifyPaperReq_ADD ManualModifyPaperReq_ModifyType = 0
	ManualModifyPaperReq_DEL ManualModifyPaperReq_ModifyType = 1
)

var ManualModifyPaperReq_ModifyType_name = map[int32]string{
	0: "ADD",
	1: "DEL",
}

var ManualModifyPaperReq_ModifyType_value = map[string]int32{
	"ADD": 0,
	"DEL": 1,
}

func (x ManualModifyPaperReq_ModifyType) String() string {
	return proto.EnumName(ManualModifyPaperReq_ModifyType_name, int32(x))
}

func (ManualModifyPaperReq_ModifyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d655f297bc51c60, []int{2, 0}
}

// 创建试卷
type NewPaperReq struct {
	Paper                *Paper   `protobuf:"bytes,1,opt,name=paper,proto3" json:"paper,omitempty"`
	Algorithm            int64    `protobuf:"varint,2,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPaperReq) Reset()         { *m = NewPaperReq{} }
func (m *NewPaperReq) String() string { return proto.CompactTextString(m) }
func (*NewPaperReq) ProtoMessage()    {}
func (*NewPaperReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d655f297bc51c60, []int{0}
}

func (m *NewPaperReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPaperReq.Unmarshal(m, b)
}
func (m *NewPaperReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPaperReq.Marshal(b, m, deterministic)
}
func (m *NewPaperReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPaperReq.Merge(m, src)
}
func (m *NewPaperReq) XXX_Size() int {
	return xxx_messageInfo_NewPaperReq.Size(m)
}
func (m *NewPaperReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPaperReq.DiscardUnknown(m)
}

var xxx_messageInfo_NewPaperReq proto.InternalMessageInfo

func (m *NewPaperReq) GetPaper() *Paper {
	if m != nil {
		return m.Paper
	}
	return nil
}

func (m *NewPaperReq) GetAlgorithm() int64 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

type NewPaperResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Paper                *Paper   `protobuf:"bytes,3,opt,name=paper,proto3" json:"paper,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPaperResp) Reset()         { *m = NewPaperResp{} }
func (m *NewPaperResp) String() string { return proto.CompactTextString(m) }
func (*NewPaperResp) ProtoMessage()    {}
func (*NewPaperResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d655f297bc51c60, []int{1}
}

func (m *NewPaperResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPaperResp.Unmarshal(m, b)
}
func (m *NewPaperResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPaperResp.Marshal(b, m, deterministic)
}
func (m *NewPaperResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPaperResp.Merge(m, src)
}
func (m *NewPaperResp) XXX_Size() int {
	return xxx_messageInfo_NewPaperResp.Size(m)
}
func (m *NewPaperResp) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPaperResp.DiscardUnknown(m)
}

var xxx_messageInfo_NewPaperResp proto.InternalMessageInfo

func (m *NewPaperResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *NewPaperResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *NewPaperResp) GetPaper() *Paper {
	if m != nil {
		return m.Paper
	}
	return nil
}

// 手动修改试卷
type ManualModifyPaperReq struct {
	ModifyType           ManualModifyPaperReq_ModifyType `protobuf:"varint,1,opt,name=modifyType,proto3,enum=protocol.ManualModifyPaperReq_ModifyType" json:"modifyType,omitempty"`
	ProblemId            int64                           `protobuf:"varint,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	PaperId              int64                           `protobuf:"varint,3,opt,name=paper_id,json=paperId,proto3" json:"paper_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ManualModifyPaperReq) Reset()         { *m = ManualModifyPaperReq{} }
func (m *ManualModifyPaperReq) String() string { return proto.CompactTextString(m) }
func (*ManualModifyPaperReq) ProtoMessage()    {}
func (*ManualModifyPaperReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d655f297bc51c60, []int{2}
}

func (m *ManualModifyPaperReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualModifyPaperReq.Unmarshal(m, b)
}
func (m *ManualModifyPaperReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualModifyPaperReq.Marshal(b, m, deterministic)
}
func (m *ManualModifyPaperReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualModifyPaperReq.Merge(m, src)
}
func (m *ManualModifyPaperReq) XXX_Size() int {
	return xxx_messageInfo_ManualModifyPaperReq.Size(m)
}
func (m *ManualModifyPaperReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualModifyPaperReq.DiscardUnknown(m)
}

var xxx_messageInfo_ManualModifyPaperReq proto.InternalMessageInfo

func (m *ManualModifyPaperReq) GetModifyType() ManualModifyPaperReq_ModifyType {
	if m != nil {
		return m.ModifyType
	}
	return ManualModifyPaperReq_ADD
}

func (m *ManualModifyPaperReq) GetProblemId() int64 {
	if m != nil {
		return m.ProblemId
	}
	return 0
}

func (m *ManualModifyPaperReq) GetPaperId() int64 {
	if m != nil {
		return m.PaperId
	}
	return 0
}

type ManualModifyPaperResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualModifyPaperResp) Reset()         { *m = ManualModifyPaperResp{} }
func (m *ManualModifyPaperResp) String() string { return proto.CompactTextString(m) }
func (*ManualModifyPaperResp) ProtoMessage()    {}
func (*ManualModifyPaperResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d655f297bc51c60, []int{3}
}

func (m *ManualModifyPaperResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualModifyPaperResp.Unmarshal(m, b)
}
func (m *ManualModifyPaperResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualModifyPaperResp.Marshal(b, m, deterministic)
}
func (m *ManualModifyPaperResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualModifyPaperResp.Merge(m, src)
}
func (m *ManualModifyPaperResp) XXX_Size() int {
	return xxx_messageInfo_ManualModifyPaperResp.Size(m)
}
func (m *ManualModifyPaperResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualModifyPaperResp.DiscardUnknown(m)
}

var xxx_messageInfo_ManualModifyPaperResp proto.InternalMessageInfo

func (m *ManualModifyPaperResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ManualModifyPaperResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterEnum("protocol.ManualModifyPaperReq_ModifyType", ManualModifyPaperReq_ModifyType_name, ManualModifyPaperReq_ModifyType_value)
	proto.RegisterType((*NewPaperReq)(nil), "protocol.NewPaperReq")
	proto.RegisterType((*NewPaperResp)(nil), "protocol.NewPaperResp")
	proto.RegisterType((*ManualModifyPaperReq)(nil), "protocol.ManualModifyPaperReq")
	proto.RegisterType((*ManualModifyPaperResp)(nil), "protocol.ManualModifyPaperResp")
}

func init() { proto.RegisterFile("proto/paper.proto", fileDescriptor_4d655f297bc51c60) }

var fileDescriptor_4d655f297bc51c60 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x41, 0x4f, 0xb3, 0x40,
	0x18, 0x84, 0x3f, 0x3e, 0x62, 0x4b, 0xdf, 0x1a, 0xa5, 0xab, 0x26, 0xd5, 0xa8, 0x69, 0x48, 0x4c,
	0xf0, 0x82, 0x49, 0x3d, 0x7a, 0xd2, 0xe0, 0x81, 0x44, 0x8c, 0xd9, 0x7a, 0xf2, 0x52, 0x29, 0xac,
	0x4a, 0x02, 0xdd, 0x95, 0x85, 0x98, 0x5e, 0xfc, 0x63, 0xfe, 0x39, 0xc3, 0xbb, 0x6c, 0xd7, 0x43,
	0xe3, 0xc9, 0x13, 0xc3, 0x0c, 0xf3, 0x30, 0xbb, 0x30, 0x12, 0x15, 0xaf, 0xf9, 0x85, 0x48, 0x04,
	0xab, 0x02, 0xd4, 0xc4, 0xc1, 0x47, 0xca, 0x8b, 0x23, 0xa2, 0xc2, 0x94, 0x97, 0x25, 0x5f, 0xaa,
	0x54, 0x7b, 0xb2, 0x4e, 0xea, 0x46, 0x2a, 0xcf, 0xa3, 0x30, 0xbc, 0x67, 0x1f, 0x0f, 0x2d, 0x83,
	0xb2, 0x77, 0x72, 0x06, 0x5b, 0xc8, 0x1b, 0x5b, 0x13, 0xcb, 0x1f, 0x4e, 0x77, 0x03, 0x0d, 0x0c,
	0xd4, 0x27, 0x2a, 0x25, 0xc7, 0x30, 0x48, 0x8a, 0x57, 0x5e, 0xe5, 0xf5, 0x5b, 0x39, 0xfe, 0x3f,
	0xb1, 0x7c, 0x9b, 0x1a, 0xc3, 0xfb, 0x84, 0x6d, 0xc3, 0x94, 0x82, 0xf8, 0xd0, 0x53, 0xff, 0xec,
	0xa8, 0xae, 0xa1, 0xce, 0xd0, 0xa7, 0x5d, 0x4e, 0x4e, 0x00, 0x72, 0x39, 0x97, 0x4d, 0x9a, 0x32,
	0x29, 0x11, 0xec, 0xd0, 0x41, 0x2e, 0x67, 0xca, 0x30, 0xeb, 0xec, 0xdf, 0xd6, 0x79, 0x5f, 0x16,
	0xec, 0xc7, 0xc9, 0xb2, 0x49, 0x8a, 0x98, 0x67, 0xf9, 0xcb, 0x6a, 0x7d, 0xba, 0x08, 0xa0, 0x44,
	0xe7, 0x71, 0x25, 0x18, 0x8e, 0xd9, 0x99, 0x9e, 0x1b, 0xc8, 0xa6, 0x4e, 0x10, 0xaf, 0x0b, 0xf4,
	0x47, 0xb9, 0x5d, 0x2a, 0x2a, 0xbe, 0x28, 0x58, 0x39, 0xcf, 0x33, 0x7d, 0x05, 0x9d, 0x13, 0x65,
	0xe4, 0x10, 0x1c, 0xdc, 0xd2, 0x86, 0x36, 0x86, 0x7d, 0x7c, 0x8f, 0x32, 0xef, 0x14, 0xc0, 0x30,
	0x49, 0x1f, 0xec, 0xeb, 0x30, 0x74, 0xff, 0xb5, 0x22, 0xbc, 0xbd, 0x73, 0x2d, 0xef, 0x19, 0x0e,
	0x36, 0x0c, 0xf9, 0xc3, 0x6b, 0xbc, 0xd9, 0x7b, 0x1a, 0xe9, 0xe6, 0x95, 0x16, 0x8b, 0x1e, 0xaa,
	0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x5f, 0xa9, 0x0f, 0x56, 0x02, 0x00, 0x00,
}
