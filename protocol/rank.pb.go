// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rank.proto

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

// 排行榜信息
type RankListResq struct {
	PageIndex            int64    `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankListResq) Reset()         { *m = RankListResq{} }
func (m *RankListResq) String() string { return proto.CompactTextString(m) }
func (*RankListResq) ProtoMessage()    {}
func (*RankListResq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c4dbbb5cfa23ee, []int{0}
}

func (m *RankListResq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankListResq.Unmarshal(m, b)
}
func (m *RankListResq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankListResq.Marshal(b, m, deterministic)
}
func (m *RankListResq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankListResq.Merge(m, src)
}
func (m *RankListResq) XXX_Size() int {
	return xxx_messageInfo_RankListResq.Size(m)
}
func (m *RankListResq) XXX_DiscardUnknown() {
	xxx_messageInfo_RankListResq.DiscardUnknown(m)
}

var xxx_messageInfo_RankListResq proto.InternalMessageInfo

func (m *RankListResq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *RankListResq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type RankListResp struct {
	Status               *Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Items                []*RankItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	PageIndex            int64       `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64       `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	TotalNum             int64       `protobuf:"varint,5,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	Pos                  int64       `protobuf:"varint,6,opt,name=pos,proto3" json:"pos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RankListResp) Reset()         { *m = RankListResp{} }
func (m *RankListResp) String() string { return proto.CompactTextString(m) }
func (*RankListResp) ProtoMessage()    {}
func (*RankListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12c4dbbb5cfa23ee, []int{1}
}

func (m *RankListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankListResp.Unmarshal(m, b)
}
func (m *RankListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankListResp.Marshal(b, m, deterministic)
}
func (m *RankListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankListResp.Merge(m, src)
}
func (m *RankListResp) XXX_Size() int {
	return xxx_messageInfo_RankListResp.Size(m)
}
func (m *RankListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RankListResp.DiscardUnknown(m)
}

var xxx_messageInfo_RankListResp proto.InternalMessageInfo

func (m *RankListResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RankListResp) GetItems() []*RankItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RankListResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *RankListResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *RankListResp) GetTotalNum() int64 {
	if m != nil {
		return m.TotalNum
	}
	return 0
}

func (m *RankListResp) GetPos() int64 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func init() {
	proto.RegisterType((*RankListResq)(nil), "protocol.RankListResq")
	proto.RegisterType((*RankListResp)(nil), "protocol.RankListResp")
}

func init() { proto.RegisterFile("proto/rank.proto", fileDescriptor_12c4dbbb5cfa23ee) }

var fileDescriptor_12c4dbbb5cfa23ee = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x63, 0x63, 0x3a, 0x51, 0x08, 0x73, 0x8a, 0x8a, 0x50, 0x7a, 0xca, 0x29, 0x42,
	0x3d, 0x7a, 0xf3, 0x64, 0x41, 0x3c, 0xac, 0x37, 0x2f, 0xb2, 0xd6, 0xa5, 0x84, 0x76, 0x77, 0xd6,
	0xee, 0x04, 0xfc, 0x8f, 0xfe, 0x29, 0xd9, 0x59, 0x83, 0xc4, 0x83, 0xa7, 0xbc, 0x7c, 0xdf, 0xf0,
	0xf6, 0x41, 0xed, 0x8f, 0xc4, 0x74, 0x73, 0xd4, 0x6e, 0xdf, 0x49, 0xc4, 0x52, 0x3e, 0x5b, 0x3a,
	0x5c, 0x62, 0x72, 0x5b, 0xb2, 0x96, 0x5c, 0xb2, 0x23, 0x0b, 0xac, 0x79, 0x08, 0x89, 0xad, 0x1e,
	0xe0, 0x4c, 0x69, 0xb7, 0x7f, 0xec, 0x03, 0x2b, 0x13, 0x3e, 0xf0, 0x1a, 0xc0, 0xeb, 0x9d, 0x79,
	0xed, 0xdd, 0xbb, 0xf9, 0x6c, 0xb2, 0x65, 0xd6, 0xe6, 0x6a, 0x11, 0xc9, 0x26, 0x02, 0xbc, 0x80,
	0x52, 0xb4, 0x1b, 0x6c, 0x33, 0x13, 0x79, 0x1a, 0xff, 0x9f, 0x06, 0xbb, 0xfa, 0xca, 0x26, 0x55,
	0x1e, 0x5b, 0x28, 0xd2, 0x53, 0x52, 0x53, 0xad, 0xeb, 0x6e, 0x5c, 0xd7, 0x3d, 0x0b, 0x57, 0x3f,
	0x1e, 0x5b, 0x98, 0xf7, 0x6c, 0x6c, 0x68, 0x66, 0xcb, 0xbc, 0xad, 0xd6, 0xf8, 0x7b, 0x18, 0x0b,
	0x37, 0x6c, 0xac, 0x4a, 0x07, 0x7f, 0xe6, 0xe5, 0xff, 0xcd, 0x3b, 0x99, 0xcc, 0xc3, 0x2b, 0x58,
	0x30, 0xb1, 0x3e, 0x88, 0x9b, 0x8b, 0x2b, 0x05, 0x44, 0x59, 0x43, 0xee, 0x29, 0x34, 0x85, 0xe0,
	0x18, 0xef, 0xcf, 0x5f, 0xaa, 0x1d, 0xdd, 0x8d, 0x3b, 0xde, 0x0a, 0x49, 0xb7, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0x73, 0xad, 0x47, 0x73, 0x01, 0x00, 0x00,
}