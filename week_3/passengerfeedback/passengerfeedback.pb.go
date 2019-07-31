// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passengerfeedback.proto

package passengerfeedback

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{0}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type AddPassengerFeedbackRequest struct {
	PassengerFeedback    *PassengerFeedback `protobuf:"bytes,1,opt,name=passengerFeedback,proto3" json:"passengerFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddPassengerFeedbackRequest) Reset()         { *m = AddPassengerFeedbackRequest{} }
func (m *AddPassengerFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*AddPassengerFeedbackRequest) ProtoMessage()    {}
func (*AddPassengerFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{1}
}

func (m *AddPassengerFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Unmarshal(m, b)
}
func (m *AddPassengerFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *AddPassengerFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPassengerFeedbackRequest.Merge(m, src)
}
func (m *AddPassengerFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Size(m)
}
func (m *AddPassengerFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPassengerFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPassengerFeedbackRequest proto.InternalMessageInfo

func (m *AddPassengerFeedbackRequest) GetPassengerFeedback() *PassengerFeedback {
	if m != nil {
		return m.PassengerFeedback
	}
	return nil
}

type GetPassengerFeedbackByPassengerIdResponse struct {
	PassengerFeedbacks   []*PassengerFeedback `protobuf:"bytes,1,rep,name=passengerFeedbacks,proto3" json:"passengerFeedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetPassengerFeedbackByPassengerIdResponse) Reset() {
	*m = GetPassengerFeedbackByPassengerIdResponse{}
}
func (m *GetPassengerFeedbackByPassengerIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetPassengerFeedbackByPassengerIdResponse) ProtoMessage()    {}
func (*GetPassengerFeedbackByPassengerIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{2}
}

func (m *GetPassengerFeedbackByPassengerIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPassengerFeedbackByPassengerIdResponse.Unmarshal(m, b)
}
func (m *GetPassengerFeedbackByPassengerIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPassengerFeedbackByPassengerIdResponse.Marshal(b, m, deterministic)
}
func (m *GetPassengerFeedbackByPassengerIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPassengerFeedbackByPassengerIdResponse.Merge(m, src)
}
func (m *GetPassengerFeedbackByPassengerIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetPassengerFeedbackByPassengerIdResponse.Size(m)
}
func (m *GetPassengerFeedbackByPassengerIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPassengerFeedbackByPassengerIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPassengerFeedbackByPassengerIdResponse proto.InternalMessageInfo

func (m *GetPassengerFeedbackByPassengerIdResponse) GetPassengerFeedbacks() []*PassengerFeedback {
	if m != nil {
		return m.PassengerFeedbacks
	}
	return nil
}

type GetPassengerFeedbackByBookingCodeResponse struct {
	PassengerFeedback    *PassengerFeedback `protobuf:"bytes,1,opt,name=passengerFeedback,proto3" json:"passengerFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetPassengerFeedbackByBookingCodeResponse) Reset() {
	*m = GetPassengerFeedbackByBookingCodeResponse{}
}
func (m *GetPassengerFeedbackByBookingCodeResponse) String() string { return proto.CompactTextString(m) }
func (*GetPassengerFeedbackByBookingCodeResponse) ProtoMessage()    {}
func (*GetPassengerFeedbackByBookingCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{3}
}

func (m *GetPassengerFeedbackByBookingCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPassengerFeedbackByBookingCodeResponse.Unmarshal(m, b)
}
func (m *GetPassengerFeedbackByBookingCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPassengerFeedbackByBookingCodeResponse.Marshal(b, m, deterministic)
}
func (m *GetPassengerFeedbackByBookingCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPassengerFeedbackByBookingCodeResponse.Merge(m, src)
}
func (m *GetPassengerFeedbackByBookingCodeResponse) XXX_Size() int {
	return xxx_messageInfo_GetPassengerFeedbackByBookingCodeResponse.Size(m)
}
func (m *GetPassengerFeedbackByBookingCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPassengerFeedbackByBookingCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPassengerFeedbackByBookingCodeResponse proto.InternalMessageInfo

func (m *GetPassengerFeedbackByBookingCodeResponse) GetPassengerFeedback() *PassengerFeedback {
	if m != nil {
		return m.PassengerFeedback
	}
	return nil
}

type AddPassengerFeedbackResponse struct {
	PassengerFeedback    *PassengerFeedback `protobuf:"bytes,1,opt,name=passengerFeedback,proto3" json:"passengerFeedback,omitempty"`
	Success              bool               `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddPassengerFeedbackResponse) Reset()         { *m = AddPassengerFeedbackResponse{} }
func (m *AddPassengerFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*AddPassengerFeedbackResponse) ProtoMessage()    {}
func (*AddPassengerFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{4}
}

func (m *AddPassengerFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Unmarshal(m, b)
}
func (m *AddPassengerFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *AddPassengerFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPassengerFeedbackResponse.Merge(m, src)
}
func (m *AddPassengerFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Size(m)
}
func (m *AddPassengerFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPassengerFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPassengerFeedbackResponse proto.InternalMessageInfo

func (m *AddPassengerFeedbackResponse) GetPassengerFeedback() *PassengerFeedback {
	if m != nil {
		return m.PassengerFeedback
	}
	return nil
}

func (m *AddPassengerFeedbackResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetPassengerFeedbackByPassengerRequest struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPassengerFeedbackByPassengerRequest) Reset() {
	*m = GetPassengerFeedbackByPassengerRequest{}
}
func (m *GetPassengerFeedbackByPassengerRequest) String() string { return proto.CompactTextString(m) }
func (*GetPassengerFeedbackByPassengerRequest) ProtoMessage()    {}
func (*GetPassengerFeedbackByPassengerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{5}
}

func (m *GetPassengerFeedbackByPassengerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPassengerFeedbackByPassengerRequest.Unmarshal(m, b)
}
func (m *GetPassengerFeedbackByPassengerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPassengerFeedbackByPassengerRequest.Marshal(b, m, deterministic)
}
func (m *GetPassengerFeedbackByPassengerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPassengerFeedbackByPassengerRequest.Merge(m, src)
}
func (m *GetPassengerFeedbackByPassengerRequest) XXX_Size() int {
	return xxx_messageInfo_GetPassengerFeedbackByPassengerRequest.Size(m)
}
func (m *GetPassengerFeedbackByPassengerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPassengerFeedbackByPassengerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPassengerFeedbackByPassengerRequest proto.InternalMessageInfo

func (m *GetPassengerFeedbackByPassengerRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type GetPassengerFeedbackByBookingCodeRequest struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPassengerFeedbackByBookingCodeRequest) Reset() {
	*m = GetPassengerFeedbackByBookingCodeRequest{}
}
func (m *GetPassengerFeedbackByBookingCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetPassengerFeedbackByBookingCodeRequest) ProtoMessage()    {}
func (*GetPassengerFeedbackByBookingCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{6}
}

func (m *GetPassengerFeedbackByBookingCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPassengerFeedbackByBookingCodeRequest.Unmarshal(m, b)
}
func (m *GetPassengerFeedbackByBookingCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPassengerFeedbackByBookingCodeRequest.Marshal(b, m, deterministic)
}
func (m *GetPassengerFeedbackByBookingCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPassengerFeedbackByBookingCodeRequest.Merge(m, src)
}
func (m *GetPassengerFeedbackByBookingCodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetPassengerFeedbackByBookingCodeRequest.Size(m)
}
func (m *GetPassengerFeedbackByBookingCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPassengerFeedbackByBookingCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPassengerFeedbackByBookingCodeRequest proto.InternalMessageInfo

func (m *GetPassengerFeedbackByBookingCodeRequest) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

type DeletePassengerFeedbackByPassengerIdRequest struct {
	PassengerId          int32    `protobuf:"varint,1,opt,name=passengerId,proto3" json:"passengerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePassengerFeedbackByPassengerIdRequest) Reset() {
	*m = DeletePassengerFeedbackByPassengerIdRequest{}
}
func (m *DeletePassengerFeedbackByPassengerIdRequest) String() string {
	return proto.CompactTextString(m)
}
func (*DeletePassengerFeedbackByPassengerIdRequest) ProtoMessage() {}
func (*DeletePassengerFeedbackByPassengerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{7}
}

func (m *DeletePassengerFeedbackByPassengerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePassengerFeedbackByPassengerIdRequest.Unmarshal(m, b)
}
func (m *DeletePassengerFeedbackByPassengerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePassengerFeedbackByPassengerIdRequest.Marshal(b, m, deterministic)
}
func (m *DeletePassengerFeedbackByPassengerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePassengerFeedbackByPassengerIdRequest.Merge(m, src)
}
func (m *DeletePassengerFeedbackByPassengerIdRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePassengerFeedbackByPassengerIdRequest.Size(m)
}
func (m *DeletePassengerFeedbackByPassengerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePassengerFeedbackByPassengerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePassengerFeedbackByPassengerIdRequest proto.InternalMessageInfo

func (m *DeletePassengerFeedbackByPassengerIdRequest) GetPassengerId() int32 {
	if m != nil {
		return m.PassengerId
	}
	return 0
}

type DeletePassengerFeedbackByPassengerIdResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePassengerFeedbackByPassengerIdResponse) Reset() {
	*m = DeletePassengerFeedbackByPassengerIdResponse{}
}
func (m *DeletePassengerFeedbackByPassengerIdResponse) String() string {
	return proto.CompactTextString(m)
}
func (*DeletePassengerFeedbackByPassengerIdResponse) ProtoMessage() {}
func (*DeletePassengerFeedbackByPassengerIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe323b0704a1f6c, []int{8}
}

func (m *DeletePassengerFeedbackByPassengerIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePassengerFeedbackByPassengerIdResponse.Unmarshal(m, b)
}
func (m *DeletePassengerFeedbackByPassengerIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePassengerFeedbackByPassengerIdResponse.Marshal(b, m, deterministic)
}
func (m *DeletePassengerFeedbackByPassengerIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePassengerFeedbackByPassengerIdResponse.Merge(m, src)
}
func (m *DeletePassengerFeedbackByPassengerIdResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePassengerFeedbackByPassengerIdResponse.Size(m)
}
func (m *DeletePassengerFeedbackByPassengerIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePassengerFeedbackByPassengerIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePassengerFeedbackByPassengerIdResponse proto.InternalMessageInfo

func (m *DeletePassengerFeedbackByPassengerIdResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*PassengerFeedback)(nil), "passengerfeedback.PassengerFeedback")
	proto.RegisterType((*AddPassengerFeedbackRequest)(nil), "passengerfeedback.AddPassengerFeedbackRequest")
	proto.RegisterType((*GetPassengerFeedbackByPassengerIdResponse)(nil), "passengerfeedback.GetPassengerFeedbackByPassengerIdResponse")
	proto.RegisterType((*GetPassengerFeedbackByBookingCodeResponse)(nil), "passengerfeedback.GetPassengerFeedbackByBookingCodeResponse")
	proto.RegisterType((*AddPassengerFeedbackResponse)(nil), "passengerfeedback.AddPassengerFeedbackResponse")
	proto.RegisterType((*GetPassengerFeedbackByPassengerRequest)(nil), "passengerfeedback.GetPassengerFeedbackByPassengerRequest")
	proto.RegisterType((*GetPassengerFeedbackByBookingCodeRequest)(nil), "passengerfeedback.GetPassengerFeedbackByBookingCodeRequest")
	proto.RegisterType((*DeletePassengerFeedbackByPassengerIdRequest)(nil), "passengerfeedback.DeletePassengerFeedbackByPassengerIdRequest")
	proto.RegisterType((*DeletePassengerFeedbackByPassengerIdResponse)(nil), "passengerfeedback.DeletePassengerFeedbackByPassengerIdResponse")
}

func init() { proto.RegisterFile("passengerfeedback.proto", fileDescriptor_fbe323b0704a1f6c) }

var fileDescriptor_fbe323b0704a1f6c = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x51, 0x4b, 0x32, 0x41,
	0x14, 0x75, 0x3e, 0xf1, 0xcb, 0xae, 0x4f, 0x5e, 0x82, 0x96, 0xad, 0x87, 0x6d, 0x90, 0x30, 0x0a,
	0x03, 0x7b, 0x8a, 0xa2, 0xc8, 0xa4, 0x32, 0x82, 0x62, 0xeb, 0x0f, 0xe8, 0xce, 0x4d, 0xc4, 0x70,
	0xd5, 0x59, 0x03, 0x9f, 0xa2, 0x87, 0xde, 0xfa, 0x0d, 0xd1, 0x6b, 0xff, 0x32, 0xda, 0xdc, 0x75,
	0x6b, 0x46, 0x1d, 0x03, 0x1f, 0xe7, 0xce, 0xdc, 0x73, 0xce, 0x9e, 0x7b, 0x2e, 0x0b, 0xab, 0xdd,
	0xba, 0x94, 0xd4, 0x69, 0x52, 0xff, 0x9e, 0x48, 0x34, 0xea, 0x5e, 0xbb, 0xd4, 0xed, 0xfb, 0x81,
	0x8f, 0x79, 0xe5, 0x82, 0x4b, 0xc8, 0xdf, 0x44, 0xc5, 0xb3, 0x51, 0x11, 0x1d, 0xc8, 0x35, 0x7c,
	0xbf, 0xdd, 0xea, 0x34, 0x4f, 0x7d, 0x41, 0x16, 0x73, 0x58, 0x71, 0xd9, 0x4d, 0x96, 0xbe, 0x5e,
	0xc4, 0x58, 0xb5, 0xaa, 0xf5, 0xcf, 0x61, 0xc5, 0x8c, 0x9b, 0x2c, 0xa1, 0x0d, 0xd9, 0x88, 0xc4,
	0x4a, 0x87, 0x00, 0xf1, 0x99, 0xf7, 0x60, 0xed, 0x44, 0x08, 0x85, 0xd7, 0xa5, 0xde, 0x80, 0x64,
	0x80, 0x2e, 0x8c, 0x85, 0x46, 0x77, 0xa1, 0x88, 0x5c, 0xb9, 0x50, 0x52, 0xbf, 0x4d, 0xc5, 0x51,
	0xdb, 0xf9, 0x33, 0x83, 0xad, 0x73, 0x0a, 0x94, 0xb7, 0x95, 0x61, 0x5c, 0xaa, 0x09, 0x97, 0x64,
	0xd7, 0xef, 0x48, 0xc2, 0x3b, 0x40, 0x05, 0x42, 0x5a, 0xcc, 0x49, 0x1b, 0x4b, 0xd0, 0xf4, 0xf3,
	0xa7, 0x49, 0x12, 0x2a, 0x63, 0x67, 0x63, 0x09, 0x8b, 0x30, 0xe1, 0x95, 0xc1, 0xba, 0xde, 0xf8,
	0xc5, 0x91, 0xa2, 0x05, 0x4b, 0x72, 0xe0, 0x79, 0x24, 0x65, 0x18, 0x93, 0xac, 0x1b, 0x1d, 0xf9,
	0x25, 0x6c, 0xce, 0x18, 0x49, 0x94, 0x88, 0x5f, 0x71, 0x63, 0x4a, 0xdc, 0xf8, 0x15, 0x14, 0x0d,
	0xbc, 0x8d, 0xd1, 0xa6, 0xc7, 0x9b, 0x5f, 0xc3, 0x76, 0x95, 0x1e, 0x28, 0xa0, 0x59, 0x79, 0xd1,
	0xc8, 0x13, 0xaa, 0x3c, 0xc1, 0x2f, 0x60, 0xc7, 0x0c, 0x70, 0x34, 0x88, 0x84, 0x69, 0xec, 0x87,
	0x69, 0xe5, 0x97, 0x0c, 0x58, 0x0a, 0xc8, 0x2d, 0xf5, 0x1f, 0x5b, 0x1e, 0xe1, 0x10, 0x56, 0x74,
	0xf3, 0xc5, 0x92, 0x66, 0x78, 0x53, 0x36, 0xd0, 0xde, 0x35, 0x7e, 0xff, 0xad, 0x97, 0xa7, 0xf0,
	0x8d, 0xc1, 0xc6, 0xcc, 0x05, 0xc3, 0x7d, 0x0d, 0xb0, 0x59, 0x06, 0xec, 0xc3, 0xf9, 0x5b, 0xc7,
	0x86, 0xf2, 0x14, 0xbe, 0x4f, 0x14, 0x98, 0x88, 0x08, 0x1e, 0x18, 0xb3, 0xa8, 0xc1, 0x9a, 0x43,
	0xa2, 0x66, 0xe3, 0x79, 0x0a, 0x3f, 0x18, 0x14, 0x4c, 0x62, 0x82, 0x47, 0x1a, 0xa2, 0x39, 0x02,
	0x6b, 0x1f, 0xff, 0xb9, 0x3f, 0xd2, 0xda, 0xf8, 0x1f, 0xfe, 0x52, 0xf6, 0x3e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0x11, 0xda, 0x9d, 0x6d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PassengerFeedbackServiceClient is the client API for PassengerFeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerFeedbackServiceClient interface {
	AddPassengerFeedback(ctx context.Context, in *AddPassengerFeedbackRequest, opts ...grpc.CallOption) (*AddPassengerFeedbackResponse, error)
	GetPassengerFeedbackByPassengerId(ctx context.Context, in *GetPassengerFeedbackByPassengerRequest, opts ...grpc.CallOption) (*GetPassengerFeedbackByPassengerIdResponse, error)
	GetPassengerFeedbackByBookingCode(ctx context.Context, in *GetPassengerFeedbackByBookingCodeRequest, opts ...grpc.CallOption) (*GetPassengerFeedbackByBookingCodeResponse, error)
	DeletePassengerFeedbackByPassengerId(ctx context.Context, in *DeletePassengerFeedbackByPassengerIdRequest, opts ...grpc.CallOption) (*DeletePassengerFeedbackByPassengerIdResponse, error)
}

type passengerFeedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewPassengerFeedbackServiceClient(cc *grpc.ClientConn) PassengerFeedbackServiceClient {
	return &passengerFeedbackServiceClient{cc}
}

func (c *passengerFeedbackServiceClient) AddPassengerFeedback(ctx context.Context, in *AddPassengerFeedbackRequest, opts ...grpc.CallOption) (*AddPassengerFeedbackResponse, error) {
	out := new(AddPassengerFeedbackResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/AddPassengerFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetPassengerFeedbackByPassengerId(ctx context.Context, in *GetPassengerFeedbackByPassengerRequest, opts ...grpc.CallOption) (*GetPassengerFeedbackByPassengerIdResponse, error) {
	out := new(GetPassengerFeedbackByPassengerIdResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/GetPassengerFeedbackByPassengerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetPassengerFeedbackByBookingCode(ctx context.Context, in *GetPassengerFeedbackByBookingCodeRequest, opts ...grpc.CallOption) (*GetPassengerFeedbackByBookingCodeResponse, error) {
	out := new(GetPassengerFeedbackByBookingCodeResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/GetPassengerFeedbackByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) DeletePassengerFeedbackByPassengerId(ctx context.Context, in *DeletePassengerFeedbackByPassengerIdRequest, opts ...grpc.CallOption) (*DeletePassengerFeedbackByPassengerIdResponse, error) {
	out := new(DeletePassengerFeedbackByPassengerIdResponse)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackService/DeletePassengerFeedbackByPassengerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerFeedbackServiceServer is the server API for PassengerFeedbackService service.
type PassengerFeedbackServiceServer interface {
	AddPassengerFeedback(context.Context, *AddPassengerFeedbackRequest) (*AddPassengerFeedbackResponse, error)
	GetPassengerFeedbackByPassengerId(context.Context, *GetPassengerFeedbackByPassengerRequest) (*GetPassengerFeedbackByPassengerIdResponse, error)
	GetPassengerFeedbackByBookingCode(context.Context, *GetPassengerFeedbackByBookingCodeRequest) (*GetPassengerFeedbackByBookingCodeResponse, error)
	DeletePassengerFeedbackByPassengerId(context.Context, *DeletePassengerFeedbackByPassengerIdRequest) (*DeletePassengerFeedbackByPassengerIdResponse, error)
}

// UnimplementedPassengerFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerFeedbackServiceServer struct {
}

func (*UnimplementedPassengerFeedbackServiceServer) AddPassengerFeedback(ctx context.Context, req *AddPassengerFeedbackRequest) (*AddPassengerFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPassengerFeedback not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetPassengerFeedbackByPassengerId(ctx context.Context, req *GetPassengerFeedbackByPassengerRequest) (*GetPassengerFeedbackByPassengerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassengerFeedbackByPassengerId not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetPassengerFeedbackByBookingCode(ctx context.Context, req *GetPassengerFeedbackByBookingCodeRequest) (*GetPassengerFeedbackByBookingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassengerFeedbackByBookingCode not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) DeletePassengerFeedbackByPassengerId(ctx context.Context, req *DeletePassengerFeedbackByPassengerIdRequest) (*DeletePassengerFeedbackByPassengerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassengerFeedbackByPassengerId not implemented")
}

func RegisterPassengerFeedbackServiceServer(s *grpc.Server, srv PassengerFeedbackServiceServer) {
	s.RegisterService(&_PassengerFeedbackService_serviceDesc, srv)
}

func _PassengerFeedbackService_AddPassengerFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPassengerFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).AddPassengerFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/AddPassengerFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).AddPassengerFeedback(ctx, req.(*AddPassengerFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetPassengerFeedbackByPassengerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPassengerFeedbackByPassengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetPassengerFeedbackByPassengerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/GetPassengerFeedbackByPassengerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetPassengerFeedbackByPassengerId(ctx, req.(*GetPassengerFeedbackByPassengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetPassengerFeedbackByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPassengerFeedbackByBookingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetPassengerFeedbackByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/GetPassengerFeedbackByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetPassengerFeedbackByBookingCode(ctx, req.(*GetPassengerFeedbackByBookingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_DeletePassengerFeedbackByPassengerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePassengerFeedbackByPassengerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).DeletePassengerFeedbackByPassengerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackService/DeletePassengerFeedbackByPassengerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).DeletePassengerFeedbackByPassengerId(ctx, req.(*DeletePassengerFeedbackByPassengerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PassengerFeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passengerfeedback.PassengerFeedbackService",
	HandlerType: (*PassengerFeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPassengerFeedback",
			Handler:    _PassengerFeedbackService_AddPassengerFeedback_Handler,
		},
		{
			MethodName: "GetPassengerFeedbackByPassengerId",
			Handler:    _PassengerFeedbackService_GetPassengerFeedbackByPassengerId_Handler,
		},
		{
			MethodName: "GetPassengerFeedbackByBookingCode",
			Handler:    _PassengerFeedbackService_GetPassengerFeedbackByBookingCode_Handler,
		},
		{
			MethodName: "DeletePassengerFeedbackByPassengerId",
			Handler:    _PassengerFeedbackService_DeletePassengerFeedbackByPassengerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passengerfeedback.proto",
}