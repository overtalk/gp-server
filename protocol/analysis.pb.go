// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/analysis.proto

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

// 根据难度进行数据分析
type AnalysisByDifficultyReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64    `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalysisByDifficultyReq) Reset()         { *m = AnalysisByDifficultyReq{} }
func (m *AnalysisByDifficultyReq) String() string { return proto.CompactTextString(m) }
func (*AnalysisByDifficultyReq) ProtoMessage()    {}
func (*AnalysisByDifficultyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e46a166d5b3183, []int{0}
}

func (m *AnalysisByDifficultyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisByDifficultyReq.Unmarshal(m, b)
}
func (m *AnalysisByDifficultyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisByDifficultyReq.Marshal(b, m, deterministic)
}
func (m *AnalysisByDifficultyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisByDifficultyReq.Merge(m, src)
}
func (m *AnalysisByDifficultyReq) XXX_Size() int {
	return xxx_messageInfo_AnalysisByDifficultyReq.Size(m)
}
func (m *AnalysisByDifficultyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisByDifficultyReq.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisByDifficultyReq proto.InternalMessageInfo

func (m *AnalysisByDifficultyReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AnalysisByDifficultyReq) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AnalysisByDifficultyReq) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type AnalysisByDifficultyResp struct {
	Status               *Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Line                 []float64 `protobuf:"fixed64,2,rep,packed,name=line,proto3" json:"line,omitempty"`
	Pie                  []int64   `protobuf:"varint,3,rep,packed,name=pie,proto3" json:"pie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AnalysisByDifficultyResp) Reset()         { *m = AnalysisByDifficultyResp{} }
func (m *AnalysisByDifficultyResp) String() string { return proto.CompactTextString(m) }
func (*AnalysisByDifficultyResp) ProtoMessage()    {}
func (*AnalysisByDifficultyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e46a166d5b3183, []int{1}
}

func (m *AnalysisByDifficultyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisByDifficultyResp.Unmarshal(m, b)
}
func (m *AnalysisByDifficultyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisByDifficultyResp.Marshal(b, m, deterministic)
}
func (m *AnalysisByDifficultyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisByDifficultyResp.Merge(m, src)
}
func (m *AnalysisByDifficultyResp) XXX_Size() int {
	return xxx_messageInfo_AnalysisByDifficultyResp.Size(m)
}
func (m *AnalysisByDifficultyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisByDifficultyResp.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisByDifficultyResp proto.InternalMessageInfo

func (m *AnalysisByDifficultyResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AnalysisByDifficultyResp) GetLine() []float64 {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *AnalysisByDifficultyResp) GetPie() []int64 {
	if m != nil {
		return m.Pie
	}
	return nil
}

// 根据tags进行数据分析
type AnalysisByTagsReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64    `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Tags                 []int64  `protobuf:"varint,4,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalysisByTagsReq) Reset()         { *m = AnalysisByTagsReq{} }
func (m *AnalysisByTagsReq) String() string { return proto.CompactTextString(m) }
func (*AnalysisByTagsReq) ProtoMessage()    {}
func (*AnalysisByTagsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e46a166d5b3183, []int{2}
}

func (m *AnalysisByTagsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisByTagsReq.Unmarshal(m, b)
}
func (m *AnalysisByTagsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisByTagsReq.Marshal(b, m, deterministic)
}
func (m *AnalysisByTagsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisByTagsReq.Merge(m, src)
}
func (m *AnalysisByTagsReq) XXX_Size() int {
	return xxx_messageInfo_AnalysisByTagsReq.Size(m)
}
func (m *AnalysisByTagsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisByTagsReq.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisByTagsReq proto.InternalMessageInfo

func (m *AnalysisByTagsReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AnalysisByTagsReq) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AnalysisByTagsReq) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *AnalysisByTagsReq) GetTags() []int64 {
	if m != nil {
		return m.Tags
	}
	return nil
}

type AnalysisByTagsResp struct {
	Status               *Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Line                 []float64 `protobuf:"fixed64,2,rep,packed,name=line,proto3" json:"line,omitempty"`
	Pie                  []int64   `protobuf:"varint,3,rep,packed,name=pie,proto3" json:"pie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AnalysisByTagsResp) Reset()         { *m = AnalysisByTagsResp{} }
func (m *AnalysisByTagsResp) String() string { return proto.CompactTextString(m) }
func (*AnalysisByTagsResp) ProtoMessage()    {}
func (*AnalysisByTagsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e46a166d5b3183, []int{3}
}

func (m *AnalysisByTagsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisByTagsResp.Unmarshal(m, b)
}
func (m *AnalysisByTagsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisByTagsResp.Marshal(b, m, deterministic)
}
func (m *AnalysisByTagsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisByTagsResp.Merge(m, src)
}
func (m *AnalysisByTagsResp) XXX_Size() int {
	return xxx_messageInfo_AnalysisByTagsResp.Size(m)
}
func (m *AnalysisByTagsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisByTagsResp.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisByTagsResp proto.InternalMessageInfo

func (m *AnalysisByTagsResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AnalysisByTagsResp) GetLine() []float64 {
	if m != nil {
		return m.Line
	}
	return nil
}

func (m *AnalysisByTagsResp) GetPie() []int64 {
	if m != nil {
		return m.Pie
	}
	return nil
}

func init() {
	proto.RegisterType((*AnalysisByDifficultyReq)(nil), "protocol.AnalysisByDifficultyReq")
	proto.RegisterType((*AnalysisByDifficultyResp)(nil), "protocol.AnalysisByDifficultyResp")
	proto.RegisterType((*AnalysisByTagsReq)(nil), "protocol.AnalysisByTagsReq")
	proto.RegisterType((*AnalysisByTagsResp)(nil), "protocol.AnalysisByTagsResp")
}

func init() { proto.RegisterFile("proto/analysis.proto", fileDescriptor_99e46a166d5b3183) }

var fileDescriptor_99e46a166d5b3183 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x5c, 0xa5, 0x70, 0x2c, 0xed, 0x81, 0x54, 0x0b, 0x31, 0xa0, 0x4c, 0x99, 0x82,
	0x04, 0x23, 0x13, 0x15, 0x0b, 0xab, 0xe9, 0xc4, 0x82, 0x4c, 0xed, 0x56, 0xae, 0x4c, 0x12, 0x72,
	0x8e, 0x44, 0xfe, 0x3d, 0xf2, 0x35, 0x56, 0x24, 0xc4, 0x08, 0x53, 0xde, 0xdd, 0x17, 0xf9, 0x7b,
	0x07, 0x97, 0x6d, 0xd7, 0x84, 0xe6, 0x56, 0xd7, 0xda, 0x0f, 0xe4, 0xa8, 0xe2, 0x11, 0x4f, 0xf9,
	0xb3, 0x6d, 0xfc, 0x15, 0x1e, 0x39, 0x05, 0x1d, 0xfa, 0x91, 0x16, 0x07, 0x58, 0x3d, 0x8e, 0xff,
	0xaf, 0x87, 0x27, 0xb7, 0xdb, 0xb9, 0x6d, 0xef, 0xc3, 0xa0, 0xec, 0x27, 0xae, 0x60, 0xde, 0x93,
	0xed, 0xde, 0x9c, 0x91, 0xd9, 0x4d, 0x56, 0x0a, 0x95, 0xc7, 0xf1, 0xd9, 0xe0, 0x35, 0x9c, 0x51,
	0xd0, 0x5d, 0xd8, 0xb8, 0x0f, 0x2b, 0x4f, 0x18, 0x4d, 0x0b, 0x94, 0x30, 0xb7, 0xb5, 0x61, 0x26,
	0x98, 0xa5, 0xb1, 0x38, 0x80, 0xfc, 0xdd, 0x45, 0x2d, 0x96, 0x90, 0x1f, 0x7b, 0xb1, 0xeb, 0xfc,
	0x6e, 0x51, 0xa5, 0xda, 0xd5, 0x0b, 0xef, 0xd5, 0xc8, 0x11, 0x61, 0xe6, 0x5d, 0x1d, 0xc5, 0xa2,
	0xcc, 0x14, 0x67, 0x5c, 0x80, 0x68, 0x5d, 0xf4, 0x89, 0x52, 0xa8, 0x18, 0x8b, 0x2f, 0x58, 0x4e,
	0xae, 0x8d, 0xde, 0xd3, 0x7f, 0x5c, 0x14, 0xbb, 0x04, 0xbd, 0x27, 0x39, 0x63, 0x31, 0xe7, 0xc2,
	0x00, 0xfe, 0x34, 0xff, 0xfd, 0x7d, 0xeb, 0x8b, 0xd7, 0x65, 0x7a, 0xe0, 0x21, 0x85, 0xf7, 0x9c,
	0xd3, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa1, 0x6c, 0x69, 0x09, 0x02, 0x00, 0x00,
}
