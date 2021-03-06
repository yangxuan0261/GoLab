// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/go-lab/test_protobuf/test.proto

package goprotobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 所有的接口
type PBMessageType int32

const (
	// 学生相关
	PBMessageType_getStudentList PBMessageType = 0
)

var PBMessageType_name = map[int32]string{
	0: "getStudentList",
}
var PBMessageType_value = map[string]int32{
	"getStudentList": 0,
}

func (x PBMessageType) Enum() *PBMessageType {
	p := new(PBMessageType)
	*p = x
	return p
}
func (x PBMessageType) String() string {
	return proto.EnumName(PBMessageType_name, int32(x))
}
func (x *PBMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBMessageType_value, data, "PBMessageType")
	if err != nil {
		return err
	}
	*x = PBMessageType(value)
	return nil
}
func (PBMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_a02d4af750960a73, []int{0}
}

type HelloWorld struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Str                  *string  `protobuf:"bytes,2,req,name=str" json:"str,omitempty"`
	Opt                  *int32   `protobuf:"varint,3,opt,name=opt" json:"opt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorld) Reset()         { *m = HelloWorld{} }
func (m *HelloWorld) String() string { return proto.CompactTextString(m) }
func (*HelloWorld) ProtoMessage()    {}
func (*HelloWorld) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_a02d4af750960a73, []int{0}
}
func (m *HelloWorld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorld.Unmarshal(m, b)
}
func (m *HelloWorld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorld.Marshal(b, m, deterministic)
}
func (dst *HelloWorld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorld.Merge(dst, src)
}
func (m *HelloWorld) XXX_Size() int {
	return xxx_messageInfo_HelloWorld.Size(m)
}
func (m *HelloWorld) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorld.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorld proto.InternalMessageInfo

func (m *HelloWorld) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HelloWorld) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

func (m *HelloWorld) GetOpt() int32 {
	if m != nil && m.Opt != nil {
		return *m.Opt
	}
	return 0
}

// 公共请求体
type PBMessageRequest struct {
	Type                 *uint32  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	MessageData          []byte   `protobuf:"bytes,2,opt,name=messageData" json:"messageData,omitempty"`
	Timestamp            *uint64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Version              *string  `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Token                *string  `protobuf:"bytes,14,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBMessageRequest) Reset()         { *m = PBMessageRequest{} }
func (m *PBMessageRequest) String() string { return proto.CompactTextString(m) }
func (*PBMessageRequest) ProtoMessage()    {}
func (*PBMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_a02d4af750960a73, []int{1}
}
func (m *PBMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBMessageRequest.Unmarshal(m, b)
}
func (m *PBMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBMessageRequest.Marshal(b, m, deterministic)
}
func (dst *PBMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBMessageRequest.Merge(dst, src)
}
func (m *PBMessageRequest) XXX_Size() int {
	return xxx_messageInfo_PBMessageRequest.Size(m)
}
func (m *PBMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PBMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PBMessageRequest proto.InternalMessageInfo

func (m *PBMessageRequest) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PBMessageRequest) GetMessageData() []byte {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *PBMessageRequest) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *PBMessageRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *PBMessageRequest) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// 消息响应包
type PBMessageResponse struct {
	Type2                *uint32  `protobuf:"varint,3,opt,name=type2" json:"type2,omitempty"`
	MessageData          []byte   `protobuf:"bytes,4,opt,name=messageData" json:"messageData,omitempty"`
	ResultCode           *uint32  `protobuf:"varint,6,opt,name=resultCode" json:"resultCode,omitempty"`
	ResultInfo           *string  `protobuf:"bytes,7,opt,name=resultInfo" json:"resultInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBMessageResponse) Reset()         { *m = PBMessageResponse{} }
func (m *PBMessageResponse) String() string { return proto.CompactTextString(m) }
func (*PBMessageResponse) ProtoMessage()    {}
func (*PBMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_a02d4af750960a73, []int{2}
}
func (m *PBMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBMessageResponse.Unmarshal(m, b)
}
func (m *PBMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBMessageResponse.Marshal(b, m, deterministic)
}
func (dst *PBMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBMessageResponse.Merge(dst, src)
}
func (m *PBMessageResponse) XXX_Size() int {
	return xxx_messageInfo_PBMessageResponse.Size(m)
}
func (m *PBMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PBMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PBMessageResponse proto.InternalMessageInfo

func (m *PBMessageResponse) GetType2() uint32 {
	if m != nil && m.Type2 != nil {
		return *m.Type2
	}
	return 0
}

func (m *PBMessageResponse) GetMessageData() []byte {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *PBMessageResponse) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *PBMessageResponse) GetResultInfo() string {
	if m != nil && m.ResultInfo != nil {
		return *m.ResultInfo
	}
	return ""
}

type PBStudentListReq struct {
	Offset               *uint32  `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit                *uint32  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBStudentListReq) Reset()         { *m = PBStudentListReq{} }
func (m *PBStudentListReq) String() string { return proto.CompactTextString(m) }
func (*PBStudentListReq) ProtoMessage()    {}
func (*PBStudentListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_a02d4af750960a73, []int{3}
}
func (m *PBStudentListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBStudentListReq.Unmarshal(m, b)
}
func (m *PBStudentListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBStudentListReq.Marshal(b, m, deterministic)
}
func (dst *PBStudentListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBStudentListReq.Merge(dst, src)
}
func (m *PBStudentListReq) XXX_Size() int {
	return xxx_messageInfo_PBStudentListReq.Size(m)
}
func (m *PBStudentListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PBStudentListReq.DiscardUnknown(m)
}

var xxx_messageInfo_PBStudentListReq proto.InternalMessageInfo

func (m *PBStudentListReq) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *PBStudentListReq) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type PBStudentListRsp struct {
	List                 []uint32 `protobuf:"varint,1,rep,name=list" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBStudentListRsp) Reset()         { *m = PBStudentListRsp{} }
func (m *PBStudentListRsp) String() string { return proto.CompactTextString(m) }
func (*PBStudentListRsp) ProtoMessage()    {}
func (*PBStudentListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_a02d4af750960a73, []int{4}
}
func (m *PBStudentListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBStudentListRsp.Unmarshal(m, b)
}
func (m *PBStudentListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBStudentListRsp.Marshal(b, m, deterministic)
}
func (dst *PBStudentListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBStudentListRsp.Merge(dst, src)
}
func (m *PBStudentListRsp) XXX_Size() int {
	return xxx_messageInfo_PBStudentListRsp.Size(m)
}
func (m *PBStudentListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PBStudentListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PBStudentListRsp proto.InternalMessageInfo

func (m *PBStudentListRsp) GetList() []uint32 {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloWorld)(nil), "goprotobuf.HelloWorld")
	proto.RegisterType((*PBMessageRequest)(nil), "goprotobuf.PBMessageRequest")
	proto.RegisterType((*PBMessageResponse)(nil), "goprotobuf.PBMessageResponse")
	proto.RegisterType((*PBStudentListReq)(nil), "goprotobuf.PBStudentListReq")
	proto.RegisterType((*PBStudentListRsp)(nil), "goprotobuf.PBStudentListRsp")
	proto.RegisterEnum("goprotobuf.PBMessageType", PBMessageType_name, PBMessageType_value)
}

func init() {
	proto.RegisterFile("src/go-lab/test_protobuf/test.proto", fileDescriptor_test_a02d4af750960a73)
}

var fileDescriptor_test_a02d4af750960a73 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x85, 0xff, 0x4d, 0xd3, 0x96, 0xce, 0x6f, 0x4a, 0x5d, 0x44, 0xf6, 0x42, 0x24, 0x44, 0x90,
	0xe0, 0x85, 0x05, 0x9f, 0xa0, 0xa8, 0xa0, 0x42, 0x05, 0x59, 0x05, 0x2f, 0x25, 0x35, 0x93, 0xb2,
	0x98, 0x64, 0xd3, 0xec, 0x44, 0xe8, 0x33, 0x78, 0xef, 0xf3, 0xca, 0x6e, 0x5a, 0x1b, 0xea, 0xdd,
	0x39, 0x5f, 0x66, 0x26, 0x67, 0x0f, 0x44, 0xa6, 0x7e, 0x9f, 0xde, 0xe9, 0x79, 0xb2, 0x98, 0x12,
	0x1a, 0x7a, 0xab, 0x6a, 0x4d, 0x7a, 0xd1, 0x64, 0xce, 0x5d, 0x3a, 0xc7, 0x61, 0xa9, 0xb7, 0x38,
	0x9a, 0x01, 0xdc, 0x63, 0x9e, 0xeb, 0x57, 0x5d, 0xe7, 0x29, 0x1f, 0x83, 0xa7, 0x52, 0xc1, 0x42,
	0x2f, 0xee, 0x4b, 0x4f, 0xa5, 0x7c, 0x02, 0x3d, 0x43, 0xb5, 0xf0, 0x42, 0x2f, 0x1e, 0x49, 0x2b,
	0x2d, 0xd1, 0x15, 0x89, 0x5e, 0xc8, 0xe2, 0xbe, 0xb4, 0x32, 0xfa, 0x66, 0x30, 0x79, 0xba, 0x7e,
	0x44, 0x63, 0x92, 0x25, 0x4a, 0x5c, 0x35, 0x68, 0x88, 0x73, 0xf0, 0x69, 0x5d, 0xa1, 0x60, 0x21,
	0x8b, 0x03, 0xe9, 0x34, 0x0f, 0xe1, 0x7f, 0xd1, 0x4e, 0xdd, 0x26, 0x94, 0x08, 0x2f, 0x64, 0xf1,
	0x81, 0xec, 0x22, 0x7e, 0x02, 0x23, 0x52, 0x05, 0x1a, 0x4a, 0x8a, 0xca, 0xfd, 0xc2, 0x97, 0x3b,
	0xc0, 0x05, 0x0c, 0x3f, 0xb1, 0x36, 0x4a, 0x97, 0xc2, 0x0f, 0x59, 0x3c, 0x92, 0x5b, 0xcb, 0x8f,
	0xa0, 0x4f, 0xfa, 0x03, 0x4b, 0x31, 0x76, 0xbc, 0x35, 0xd1, 0x17, 0x83, 0xc3, 0x4e, 0x30, 0x53,
	0xe9, 0xd2, 0xa0, 0x9b, 0x5d, 0x57, 0x78, 0xe5, 0xee, 0x07, 0xb2, 0x35, 0xfb, 0xd9, 0xfc, 0xbf,
	0xd9, 0x4e, 0x01, 0x6a, 0x34, 0x4d, 0x4e, 0x37, 0x3a, 0x45, 0x31, 0x70, 0xcb, 0x1d, 0xb2, 0xfb,
	0xfe, 0x50, 0x66, 0x5a, 0x0c, 0x5d, 0x90, 0x0e, 0x89, 0x66, 0xb6, 0xa5, 0x67, 0x6a, 0x52, 0x2c,
	0x69, 0xae, 0x0c, 0x49, 0x5c, 0xf1, 0x63, 0x18, 0xe8, 0x2c, 0x33, 0x48, 0x9b, 0x9e, 0x36, 0xce,
	0x66, 0xcc, 0x55, 0xa1, 0xc8, 0x75, 0x14, 0xc8, 0xd6, 0x44, 0xe7, 0xfb, 0x17, 0x4c, 0x65, 0x7b,
	0xce, 0x95, 0xb1, 0xfb, 0x3d, 0xdb, 0xb3, 0xd5, 0x17, 0x67, 0x10, 0xfc, 0x3e, 0xfb, 0xc5, 0x16,
	0xcf, 0x61, 0xbc, 0x44, 0xea, 0x6c, 0x4e, 0xfe, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xdf,
	0xed, 0xb8, 0x28, 0x02, 0x00, 0x00,
}
