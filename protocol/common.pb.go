// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

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

// Role : 用户角色（学生/老师...）
type Role int32

const (
	Role_STUDENT Role = 0
	Role_TEACHER Role = 1
	Role_MANAGER Role = 2
)

var Role_name = map[int32]string{
	0: "STUDENT",
	1: "TEACHER",
	2: "MANAGER",
}

var Role_value = map[string]int32{
	"STUDENT": 0,
	"TEACHER": 1,
	"MANAGER": 2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

// ProblemDifficluty : 题目难度
type ProblemDifficluty int32

const (
	ProblemDifficluty_EASY   ProblemDifficluty = 0
	ProblemDifficluty_MEDIUM ProblemDifficluty = 1
	ProblemDifficluty_HARD   ProblemDifficluty = 2
)

var ProblemDifficluty_name = map[int32]string{
	0: "EASY",
	1: "MEDIUM",
	2: "HARD",
}

var ProblemDifficluty_value = map[string]int32{
	"EASY":   0,
	"MEDIUM": 1,
	"HARD":   2,
}

func (x ProblemDifficluty) String() string {
	return proto.EnumName(ProblemDifficluty_name, int32(x))
}

func (ProblemDifficluty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

// UserInfo : 用户基本信息
type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 Role     `protobuf:"varint,4,opt,name=role,proto3,enum=protocol.Role" json:"role,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Sex                  bool     `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	School               string   `protobuf:"bytes,9,opt,name=school,proto3" json:"school,omitempty"`
	LastLogin            int64    `protobuf:"varint,10,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	Create               int64    `protobuf:"varint,11,opt,name=create,proto3" json:"create,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_STUDENT
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *UserInfo) GetLastLogin() int64 {
	if m != nil {
		return m.LastLogin
	}
	return 0
}

func (m *UserInfo) GetCreate() int64 {
	if m != nil {
		return m.Create
	}
	return 0
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// ProblemExample : 题目输入输出样例
type ProblemExample struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output               string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemExample) Reset()         { *m = ProblemExample{} }
func (m *ProblemExample) String() string { return proto.CompactTextString(m) }
func (*ProblemExample) ProtoMessage()    {}
func (*ProblemExample) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *ProblemExample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemExample.Unmarshal(m, b)
}
func (m *ProblemExample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemExample.Marshal(b, m, deterministic)
}
func (m *ProblemExample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemExample.Merge(m, src)
}
func (m *ProblemExample) XXX_Size() int {
	return xxx_messageInfo_ProblemExample.Size(m)
}
func (m *ProblemExample) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemExample.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemExample proto.InternalMessageInfo

func (m *ProblemExample) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ProblemExample) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// Problem : 题目
type Problem struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	In                   string            `protobuf:"bytes,4,opt,name=in,proto3" json:"in,omitempty"`
	Out                  string            `protobuf:"bytes,5,opt,name=out,proto3" json:"out,omitempty"`
	Hint                 string            `protobuf:"bytes,6,opt,name=hint,proto3" json:"hint,omitempty"`
	InOutExamples        []*ProblemExample `protobuf:"bytes,7,rep,name=in_out_examples,json=inOutExamples,proto3" json:"in_out_examples,omitempty"`
	JudgeLimitTime       int64             `protobuf:"varint,8,opt,name=judge_limit_time,json=judgeLimitTime,proto3" json:"judge_limit_time,omitempty"`
	JudgeLimitMem        int64             `protobuf:"varint,9,opt,name=judge_limit_mem,json=judgeLimitMem,proto3" json:"judge_limit_mem,omitempty"`
	Tags                 []int64           `protobuf:"varint,10,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	Difficluty           ProblemDifficluty `protobuf:"varint,11,opt,name=difficluty,proto3,enum=protocol.ProblemDifficluty" json:"difficluty,omitempty"`
	SubmitTime           int64             `protobuf:"varint,12,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	AcceptTime           int64             `protobuf:"varint,13,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Problem) Reset()         { *m = Problem{} }
func (m *Problem) String() string { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()    {}
func (*Problem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{2}
}

func (m *Problem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Problem.Unmarshal(m, b)
}
func (m *Problem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Problem.Marshal(b, m, deterministic)
}
func (m *Problem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Problem.Merge(m, src)
}
func (m *Problem) XXX_Size() int {
	return xxx_messageInfo_Problem.Size(m)
}
func (m *Problem) XXX_DiscardUnknown() {
	xxx_messageInfo_Problem.DiscardUnknown(m)
}

var xxx_messageInfo_Problem proto.InternalMessageInfo

func (m *Problem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Problem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Problem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Problem) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Problem) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *Problem) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *Problem) GetInOutExamples() []*ProblemExample {
	if m != nil {
		return m.InOutExamples
	}
	return nil
}

func (m *Problem) GetJudgeLimitTime() int64 {
	if m != nil {
		return m.JudgeLimitTime
	}
	return 0
}

func (m *Problem) GetJudgeLimitMem() int64 {
	if m != nil {
		return m.JudgeLimitMem
	}
	return 0
}

func (m *Problem) GetTags() []int64 {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Problem) GetDifficluty() ProblemDifficluty {
	if m != nil {
		return m.Difficluty
	}
	return ProblemDifficluty_EASY
}

func (m *Problem) GetSubmitTime() int64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *Problem) GetAcceptTime() int64 {
	if m != nil {
		return m.AcceptTime
	}
	return 0
}

// SubmitRecord : 提交情况
type SubmitRecord struct {
	Problem              *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	SubmitTime           int64    `protobuf:"varint,2,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	IsPass               bool     `protobuf:"varint,3,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitRecord) Reset()         { *m = SubmitRecord{} }
func (m *SubmitRecord) String() string { return proto.CompactTextString(m) }
func (*SubmitRecord) ProtoMessage()    {}
func (*SubmitRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{3}
}

func (m *SubmitRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRecord.Unmarshal(m, b)
}
func (m *SubmitRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRecord.Marshal(b, m, deterministic)
}
func (m *SubmitRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRecord.Merge(m, src)
}
func (m *SubmitRecord) XXX_Size() int {
	return xxx_messageInfo_SubmitRecord.Size(m)
}
func (m *SubmitRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRecord proto.InternalMessageInfo

func (m *SubmitRecord) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

func (m *SubmitRecord) GetSubmitTime() int64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *SubmitRecord) GetIsPass() bool {
	if m != nil {
		return m.IsPass
	}
	return false
}

// Announcement : 公告，包括班级公告和全局公告
type Announcement struct {
	Publisher            string   `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	DisableTime          int64    `protobuf:"varint,4,opt,name=disable_time,json=disableTime,proto3" json:"disable_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announcement) Reset()         { *m = Announcement{} }
func (m *Announcement) String() string { return proto.CompactTextString(m) }
func (*Announcement) ProtoMessage()    {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{4}
}

func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Announcement.Unmarshal(m, b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
}
func (m *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(m, src)
}
func (m *Announcement) XXX_Size() int {
	return xxx_messageInfo_Announcement.Size(m)
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Announcement) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Announcement) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Announcement) GetDisableTime() int64 {
	if m != nil {
		return m.DisableTime
	}
	return 0
}

// Class : 班级信息
type Class struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tutor                string          `protobuf:"bytes,2,opt,name=tutor,proto3" json:"tutor,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string          `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Number               int64           `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	IsCheck              bool            `protobuf:"varint,6,opt,name=is_check,json=isCheck,proto3" json:"is_check,omitempty"`
	CreateTime           int64           `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Announcements        []*Announcement `protobuf:"bytes,8,rep,name=announcements,proto3" json:"announcements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{5}
}

func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Class) GetTutor() string {
	if m != nil {
		return m.Tutor
	}
	return ""
}

func (m *Class) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Class) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *Class) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Class) GetIsCheck() bool {
	if m != nil {
		return m.IsCheck
	}
	return false
}

func (m *Class) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Class) GetAnnouncements() []*Announcement {
	if m != nil {
		return m.Announcements
	}
	return nil
}

// RankItem : 排行榜信息
type RankItem struct {
	Ranking              int64    `protobuf:"varint,1,opt,name=ranking,proto3" json:"ranking,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PassNum              int64    `protobuf:"varint,3,opt,name=pass_num,json=passNum,proto3" json:"pass_num,omitempty"`
	SubmitNum            int64    `protobuf:"varint,4,opt,name=submit_num,json=submitNum,proto3" json:"submit_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankItem) Reset()         { *m = RankItem{} }
func (m *RankItem) String() string { return proto.CompactTextString(m) }
func (*RankItem) ProtoMessage()    {}
func (*RankItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{6}
}

func (m *RankItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankItem.Unmarshal(m, b)
}
func (m *RankItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankItem.Marshal(b, m, deterministic)
}
func (m *RankItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankItem.Merge(m, src)
}
func (m *RankItem) XXX_Size() int {
	return xxx_messageInfo_RankItem.Size(m)
}
func (m *RankItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RankItem.DiscardUnknown(m)
}

var xxx_messageInfo_RankItem proto.InternalMessageInfo

func (m *RankItem) GetRanking() int64 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func (m *RankItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankItem) GetPassNum() int64 {
	if m != nil {
		return m.PassNum
	}
	return 0
}

func (m *RankItem) GetSubmitNum() int64 {
	if m != nil {
		return m.SubmitNum
	}
	return 0
}

func init() {
	proto.RegisterEnum("protocol.Role", Role_name, Role_value)
	proto.RegisterEnum("protocol.ProblemDifficluty", ProblemDifficluty_name, ProblemDifficluty_value)
	proto.RegisterType((*UserInfo)(nil), "protocol.UserInfo")
	proto.RegisterType((*ProblemExample)(nil), "protocol.ProblemExample")
	proto.RegisterType((*Problem)(nil), "protocol.Problem")
	proto.RegisterType((*SubmitRecord)(nil), "protocol.SubmitRecord")
	proto.RegisterType((*Announcement)(nil), "protocol.Announcement")
	proto.RegisterType((*Class)(nil), "protocol.Class")
	proto.RegisterType((*RankItem)(nil), "protocol.RankItem")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0x25, 0xc7, 0x92, 0x8f, 0x6c, 0xd7, 0x25, 0x8a, 0x4e, 0xdd, 0x0f, 0xea, 0xf9, 0x62,
	0x30, 0x3a, 0x2c, 0x03, 0xd2, 0xcb, 0x0e, 0xc3, 0xbc, 0xc4, 0x58, 0x03, 0x34, 0x59, 0xc1, 0x24,
	0x17, 0xdb, 0x8d, 0x21, 0x4b, 0x8c, 0xcd, 0x45, 0x24, 0x05, 0x91, 0xc4, 0xba, 0x37, 0xd8, 0x0b,
	0xec, 0x35, 0xf6, 0x7a, 0xbb, 0x1d, 0x78, 0x48, 0x3b, 0x4e, 0xdc, 0x2b, 0x9d, 0xef, 0xe3, 0x27,
	0xf2, 0xfc, 0x7c, 0x24, 0x90, 0xa6, 0x55, 0x46, 0x7d, 0x5f, 0x2a, 0x21, 0x94, 0x3c, 0x46, 0x40,
	0x52, 0xfc, 0x94, 0xaa, 0x9e, 0xfe, 0x13, 0x41, 0x7a, 0xa3, 0x59, 0x7b, 0x2e, 0x6f, 0x15, 0x19,
	0x41, 0xc4, 0xab, 0xbc, 0x33, 0xe9, 0xcc, 0x62, 0x1a, 0xf1, 0x8a, 0x4c, 0xa1, 0xdb, 0xaa, 0x9a,
	0xe5, 0xdd, 0x49, 0x67, 0x36, 0x3a, 0x19, 0x1d, 0x6f, 0xff, 0x3a, 0xa6, 0xaa, 0x66, 0x14, 0xd7,
	0x08, 0x81, 0xae, 0x2c, 0x04, 0xcb, 0x8f, 0x26, 0x9d, 0x59, 0x9f, 0x62, 0x4c, 0xc6, 0x10, 0x6b,
	0xf6, 0x31, 0xef, 0x4d, 0x3a, 0xb3, 0x94, 0xba, 0x90, 0x3c, 0x87, 0xa3, 0x66, 0xa3, 0x24, 0xcb,
	0x13, 0x94, 0x79, 0xe0, 0x58, 0x26, 0x0a, 0x5e, 0xe7, 0xa9, 0x67, 0x11, 0x90, 0x17, 0xd0, 0xd3,
	0xe5, 0x46, 0xa9, 0x3a, 0xef, 0x23, 0x1d, 0x10, 0xf9, 0x0a, 0xa0, 0x2e, 0xb4, 0x59, 0xd6, 0x6a,
	0xcd, 0x65, 0x0e, 0x98, 0x65, 0xdf, 0x31, 0xef, 0x1d, 0xe1, 0x7e, 0x2b, 0x5b, 0x56, 0x18, 0x96,
	0x67, 0xb8, 0x14, 0x10, 0xc9, 0x21, 0x29, 0xca, 0x52, 0x59, 0x69, 0xf2, 0x08, 0xf7, 0xdb, 0x42,
	0xf2, 0x39, 0xa4, 0x4d, 0xa1, 0xf5, 0x9f, 0xaa, 0xad, 0xf2, 0x18, 0x97, 0x76, 0x78, 0xfa, 0x23,
	0x8c, 0x3e, 0xb4, 0x6a, 0x55, 0x33, 0xb1, 0xf8, 0x58, 0x88, 0xa6, 0xc6, 0x64, 0xb9, 0x6c, 0xac,
	0xc1, 0xfe, 0xf4, 0xa9, 0x07, 0xee, 0x54, 0x65, 0x8d, 0xa3, 0xfd, 0xe6, 0x01, 0x4d, 0xff, 0x8d,
	0x21, 0x09, 0x1b, 0x1c, 0xb4, 0xf5, 0x39, 0x1c, 0x19, 0x6e, 0x6a, 0x16, 0x7e, 0xf1, 0x80, 0x4c,
	0x20, 0xab, 0x98, 0x2e, 0x5b, 0xde, 0x18, 0xae, 0x64, 0x48, 0x68, 0x9f, 0xc2, 0x7d, 0x24, 0x0e,
	0xa3, 0x4f, 0x23, 0x2e, 0x5d, 0x9b, 0x95, 0x35, 0xa1, 0xf3, 0x2e, 0x74, 0xc3, 0xd8, 0x70, 0x69,
	0xb0, 0xf3, 0x7d, 0x8a, 0x31, 0xf9, 0x09, 0x9e, 0x72, 0xb9, 0x54, 0xd6, 0x2c, 0x99, 0xaf, 0x44,
	0xe7, 0xc9, 0x24, 0x9e, 0x65, 0x27, 0xf9, 0xfd, 0x3c, 0x1f, 0x96, 0x4a, 0x87, 0x5c, 0xfe, 0x6a,
	0x4d, 0x40, 0x9a, 0xcc, 0x60, 0xfc, 0x87, 0xad, 0xd6, 0x6c, 0x59, 0x73, 0xc1, 0xcd, 0xd2, 0x70,
	0xc1, 0x70, 0x62, 0x31, 0x1d, 0x21, 0xff, 0xde, 0xd1, 0xd7, 0x5c, 0x30, 0xf2, 0x0d, 0x3c, 0xdd,
	0x57, 0x0a, 0x26, 0x70, 0x86, 0x31, 0x1d, 0xde, 0x0b, 0x2f, 0x98, 0x70, 0x79, 0x9a, 0x62, 0xad,
	0x73, 0x98, 0xc4, 0xb3, 0x98, 0x62, 0x4c, 0xde, 0x02, 0x54, 0xfc, 0xf6, 0x96, 0x97, 0xb5, 0x35,
	0x7f, 0xe1, 0x0c, 0x47, 0x27, 0x5f, 0x1c, 0xa4, 0x78, 0xb6, 0x93, 0xd0, 0x3d, 0x39, 0x79, 0x05,
	0x99, 0xb6, 0xab, 0x5d, 0x76, 0x03, 0x3c, 0x14, 0x3c, 0x85, 0x99, 0xbd, 0x82, 0xac, 0x28, 0x4b,
	0xd6, 0x04, 0xc1, 0xd0, 0x0b, 0x3c, 0xe5, 0x04, 0x53, 0x0b, 0x83, 0x2b, 0x94, 0x53, 0x56, 0xaa,
	0xb6, 0x22, 0xdf, 0x42, 0xd2, 0xf8, 0x23, 0x71, 0x72, 0xd9, 0xc9, 0xb3, 0x83, 0x5c, 0xe8, 0x56,
	0xf1, 0xf8, 0xf8, 0xe8, 0xe0, 0xf8, 0xcf, 0x20, 0xe1, 0x7a, 0xe9, 0xdc, 0x85, 0x83, 0x4d, 0x69,
	0x8f, 0xeb, 0x0f, 0x85, 0xd6, 0xd3, 0xbf, 0x3b, 0x30, 0x98, 0x4b, 0xa9, 0xac, 0x2c, 0x99, 0x60,
	0xd2, 0x90, 0x2f, 0xa1, 0xdf, 0xd8, 0x55, 0xcd, 0xf5, 0x86, 0xb5, 0xc1, 0x6a, 0xf7, 0x84, 0xb3,
	0x5b, 0xc5, 0x8c, 0xbb, 0x32, 0xc1, 0x6e, 0x1e, 0xb9, 0x04, 0xbc, 0xdd, 0x7d, 0x02, 0xb1, 0x4f,
	0xc0, 0x53, 0x98, 0xc0, 0xd7, 0x30, 0xa8, 0xb8, 0x2e, 0x56, 0x75, 0x50, 0x74, 0x51, 0x91, 0x05,
	0x0e, 0x3b, 0xf0, 0x5f, 0x07, 0x8e, 0x4e, 0xeb, 0x42, 0xeb, 0x4f, 0x1a, 0xd6, 0x1a, 0xd5, 0xee,
	0x0c, 0xeb, 0xc0, 0xee, 0xe6, 0xc7, 0x7b, 0x37, 0x7f, 0x0a, 0x03, 0x2e, 0x4d, 0xab, 0x2a, 0x5b,
	0xa2, 0x8b, 0xbd, 0x59, 0x1f, 0x70, 0xae, 0x06, 0x69, 0xc5, 0x8a, 0xb5, 0xe8, 0xdc, 0x98, 0x06,
	0x44, 0x5e, 0x42, 0xca, 0xf5, 0xb2, 0xdc, 0xb0, 0xf2, 0x2e, 0x3c, 0x1d, 0x09, 0xd7, 0xa7, 0x0e,
	0x3e, 0x2e, 0x2f, 0x39, 0x28, 0xef, 0x07, 0x18, 0x16, 0x7b, 0x5d, 0xd4, 0x79, 0x8a, 0x16, 0x7f,
	0x71, 0x3f, 0xb3, 0xfd, 0x26, 0xd3, 0x87, 0xe2, 0x69, 0x0b, 0x29, 0x2d, 0xe4, 0xdd, 0xb9, 0x61,
	0xc2, 0x3d, 0x17, 0x6d, 0x21, 0xef, 0xb8, 0x5c, 0x87, 0x06, 0x6c, 0xe1, 0xae, 0xde, 0x68, 0xaf,
	0xde, 0x97, 0xfe, 0x09, 0x59, 0x4a, 0x2b, 0x42, 0xd3, 0x13, 0x87, 0x2f, 0xad, 0x70, 0xcf, 0x55,
	0xf0, 0x84, 0x5b, 0xf4, 0xfd, 0xee, 0x7b, 0xe6, 0xd2, 0x8a, 0xd7, 0xdf, 0x41, 0xd7, 0xbd, 0xa2,
	0x24, 0x83, 0xe4, 0xea, 0xfa, 0xe6, 0x6c, 0x71, 0x79, 0x3d, 0x7e, 0xe2, 0xc0, 0xf5, 0x62, 0x7e,
	0xfa, 0x6e, 0x41, 0xc7, 0x1d, 0x07, 0x2e, 0xe6, 0x97, 0xf3, 0x5f, 0x16, 0x74, 0x1c, 0xbd, 0x7e,
	0x03, 0xcf, 0x0e, 0x6e, 0x00, 0x49, 0xa1, 0xbb, 0x98, 0x5f, 0xfd, 0x36, 0x7e, 0x42, 0x00, 0x7a,
	0x17, 0x8b, 0xb3, 0xf3, 0x9b, 0x8b, 0x71, 0xc7, 0xb1, 0xef, 0xe6, 0xf4, 0x6c, 0x1c, 0xfd, 0x3c,
	0xfc, 0x3d, 0x5b, 0xab, 0xb7, 0xdb, 0x16, 0xac, 0x7a, 0x18, 0xbd, 0xf9, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xb7, 0x98, 0x40, 0x12, 0x06, 0x00, 0x00,
}
