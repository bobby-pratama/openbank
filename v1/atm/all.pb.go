// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/atm/all.proto

package atm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ATM holds all details about an ATM.
type ATM struct {
	// ID is the unique identifier of an ATM.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	// BankID is the identifier of the bank associated with the ATM.
	BankID string `protobuf:"bytes,2,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// Name is the name of ATM.
	Name string `protobuf:"bytes,3,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// Address is the ATM's address.
	Address *types.Address `protobuf:"bytes,4,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the ATM longitude and latitude.
	Location *types.Location `protobuf:"bytes,5,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the ATM's description.
	Description string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the ATM's metadata.
	Metadata             string   `protobuf:"bytes,7,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ATM) Reset()         { *m = ATM{} }
func (m *ATM) String() string { return proto.CompactTextString(m) }
func (*ATM) ProtoMessage()    {}
func (*ATM) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fee59d76a66276, []int{0}
}

func (m *ATM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ATM.Unmarshal(m, b)
}
func (m *ATM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ATM.Marshal(b, m, deterministic)
}
func (m *ATM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ATM.Merge(m, src)
}
func (m *ATM) XXX_Size() int {
	return xxx_messageInfo_ATM.Size(m)
}
func (m *ATM) XXX_DiscardUnknown() {
	xxx_messageInfo_ATM.DiscardUnknown(m)
}

var xxx_messageInfo_ATM proto.InternalMessageInfo

func (m *ATM) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ATM) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *ATM) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ATM) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ATM) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *ATM) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ATM) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// CreateATMRequest is a request to create an ATM.
type CreateATMRequest struct {
	// BankID is the bank identifier that owned the ATM
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// BankID is the identifier of the bank associated with the ATM.
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// Address is the ATM's address.
	Address *types.Address `protobuf:"bytes,3,opt,name=Address,json=3,proto3" json:"Address,omitempty"`
	// Location is the ATM's longitude and latitude
	Location *types.Location `protobuf:"bytes,4,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the ATM's description.
	Description string `protobuf:"bytes,5,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the ATM's metadata.
	Metadata             string   `protobuf:"bytes,6,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateATMRequest) Reset()         { *m = CreateATMRequest{} }
func (m *CreateATMRequest) String() string { return proto.CompactTextString(m) }
func (*CreateATMRequest) ProtoMessage()    {}
func (*CreateATMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fee59d76a66276, []int{1}
}

func (m *CreateATMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateATMRequest.Unmarshal(m, b)
}
func (m *CreateATMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateATMRequest.Marshal(b, m, deterministic)
}
func (m *CreateATMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateATMRequest.Merge(m, src)
}
func (m *CreateATMRequest) XXX_Size() int {
	return xxx_messageInfo_CreateATMRequest.Size(m)
}
func (m *CreateATMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateATMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateATMRequest proto.InternalMessageInfo

func (m *CreateATMRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateATMRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateATMRequest) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CreateATMRequest) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *CreateATMRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateATMRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// CreateATMResponse is the response envelope for creating an ATM.
type CreateATMResponse struct {
	// ATM_ID is the unique identifier of the ATM.
	ATM_ID               string   `protobuf:"bytes,1,opt,name=ATM_ID,json=atm_id,proto3" json:"ATM_ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateATMResponse) Reset()         { *m = CreateATMResponse{} }
func (m *CreateATMResponse) String() string { return proto.CompactTextString(m) }
func (*CreateATMResponse) ProtoMessage()    {}
func (*CreateATMResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fee59d76a66276, []int{2}
}

func (m *CreateATMResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateATMResponse.Unmarshal(m, b)
}
func (m *CreateATMResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateATMResponse.Marshal(b, m, deterministic)
}
func (m *CreateATMResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateATMResponse.Merge(m, src)
}
func (m *CreateATMResponse) XXX_Size() int {
	return xxx_messageInfo_CreateATMResponse.Size(m)
}
func (m *CreateATMResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateATMResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateATMResponse proto.InternalMessageInfo

func (m *CreateATMResponse) GetATM_ID() string {
	if m != nil {
		return m.ATM_ID
	}
	return ""
}

// GetATMRequest is the request envelope for retrieving a specific ATM's information.
type GetATMRequest struct {
	// ATM_ID is a unique identifier of the ATM.
	ATM_ID               string   `protobuf:"bytes,1,opt,name=ATM_ID,json=atm_id,proto3" json:"ATM_ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetATMRequest) Reset()         { *m = GetATMRequest{} }
func (m *GetATMRequest) String() string { return proto.CompactTextString(m) }
func (*GetATMRequest) ProtoMessage()    {}
func (*GetATMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fee59d76a66276, []int{3}
}

func (m *GetATMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetATMRequest.Unmarshal(m, b)
}
func (m *GetATMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetATMRequest.Marshal(b, m, deterministic)
}
func (m *GetATMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetATMRequest.Merge(m, src)
}
func (m *GetATMRequest) XXX_Size() int {
	return xxx_messageInfo_GetATMRequest.Size(m)
}
func (m *GetATMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetATMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetATMRequest proto.InternalMessageInfo

func (m *GetATMRequest) GetATM_ID() string {
	if m != nil {
		return m.ATM_ID
	}
	return ""
}

// GetATMsResponse is the response envelope for retrieving ATM information.
type GetATMsResponse struct {
	// Result is the list of the ATMs.
	Result               []*ATM   `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetATMsResponse) Reset()         { *m = GetATMsResponse{} }
func (m *GetATMsResponse) String() string { return proto.CompactTextString(m) }
func (*GetATMsResponse) ProtoMessage()    {}
func (*GetATMsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80fee59d76a66276, []int{4}
}

func (m *GetATMsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetATMsResponse.Unmarshal(m, b)
}
func (m *GetATMsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetATMsResponse.Marshal(b, m, deterministic)
}
func (m *GetATMsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetATMsResponse.Merge(m, src)
}
func (m *GetATMsResponse) XXX_Size() int {
	return xxx_messageInfo_GetATMsResponse.Size(m)
}
func (m *GetATMsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetATMsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetATMsResponse proto.InternalMessageInfo

func (m *GetATMsResponse) GetResult() []*ATM {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ATM)(nil), "atm.ATM")
	proto.RegisterType((*CreateATMRequest)(nil), "atm.CreateATMRequest")
	proto.RegisterType((*CreateATMResponse)(nil), "atm.CreateATMResponse")
	proto.RegisterType((*GetATMRequest)(nil), "atm.GetATMRequest")
	proto.RegisterType((*GetATMsResponse)(nil), "atm.GetATMsResponse")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/atm/all.proto", fileDescriptor_80fee59d76a66276)
}

var fileDescriptor_80fee59d76a66276 = []byte{
	// 1248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x6f, 0x1b, 0x45,
	0x1b, 0xc7, 0x77, 0xd7, 0x89, 0xed, 0x4e, 0xd4, 0xb7, 0xe9, 0xb4, 0x6f, 0x65, 0xb9, 0x55, 0x35,
	0xb8, 0xa5, 0x44, 0x55, 0xba, 0x76, 0xdc, 0xa2, 0xa2, 0x20, 0x54, 0x6d, 0x9a, 0x36, 0x24, 0xd4,
	0x10, 0x19, 0x0b, 0xa1, 0x5e, 0xa2, 0xf1, 0xee, 0x63, 0x7b, 0xe8, 0xee, 0xcc, 0x32, 0x33, 0xeb,
	0x34, 0x20, 0x24, 0x54, 0x09, 0x54, 0x89, 0x4b, 0x09, 0x77, 0x24, 0x24, 0xfe, 0x03, 0x2e, 0x9c,
	0x90, 0x38, 0x71, 0x45, 0xe2, 0xc2, 0x89, 0x0b, 0x5c, 0x90, 0x10, 0x12, 0x37, 0x8e, 0x68, 0x67,
	0xd7, 0x76, 0x1c, 0xa7, 0x4a, 0x2b, 0x38, 0xb5, 0x3b, 0xcf, 0xaf, 0xef, 0xf3, 0x99, 0x67, 0x9e,
	0x18, 0x5d, 0xeb, 0x33, 0x3d, 0x48, 0xba, 0xae, 0x2f, 0xa2, 0xba, 0x88, 0x81, 0x77, 0x29, 0x7f,
	0x30, 0xf9, 0xcf, 0x70, 0xa5, 0x4e, 0x75, 0x54, 0xa7, 0x61, 0xe8, 0xc6, 0x52, 0x68, 0x81, 0x0b,
	0x54, 0x47, 0xd5, 0x0b, 0x7d, 0x21, 0xfa, 0x21, 0xd4, 0x69, 0xcc, 0xea, 0x94, 0x73, 0xa1, 0xa9,
	0x66, 0x82, 0xab, 0xcc, 0xa5, 0xba, 0x6c, 0xfe, 0xf1, 0xaf, 0xf5, 0x81, 0x5f, 0x53, 0xbb, 0xb4,
	0xdf, 0x07, 0x59, 0x17, 0xb1, 0xf1, 0x38, 0xc2, 0xfb, 0x7c, 0x9e, 0xcb, 0x7c, 0x75, 0x93, 0x5e,
	0x1d, 0xa2, 0x58, 0xef, 0xe5, 0xc6, 0xfa, 0x71, 0xe2, 0xf4, 0x5e, 0x0c, 0x6a, 0x22, 0xaf, 0xf6,
	0x8d, 0x83, 0x0a, 0x5e, 0xa7, 0x85, 0xab, 0xc8, 0xd9, 0x5c, 0xaf, 0xd8, 0xc4, 0x5e, 0x3a, 0xb1,
	0x86, 0xca, 0x56, 0xc5, 0x5a, 0xb2, 0x1a, 0xd6, 0xb6, 0xd5, 0x76, 0x58, 0x80, 0x2f, 0xa1, 0xe2,
	0x1a, 0xe5, 0x0f, 0x36, 0xd7, 0x2b, 0xce, 0x8c, 0xbd, 0x94, 0x26, 0xde, 0x61, 0x01, 0xbe, 0x88,
	0xe6, 0xde, 0xa4, 0x11, 0x54, 0x0a, 0x33, 0x2e, 0x73, 0x9c, 0x46, 0x80, 0x6f, 0xa0, 0x92, 0x17,
	0x04, 0x12, 0x94, 0xaa, 0xcc, 0x11, 0x7b, 0x69, 0xa1, 0xf9, 0x3f, 0xd7, 0x68, 0x71, 0xf3, 0xd3,
	0xe9, 0xac, 0x34, 0x3b, 0xc4, 0x37, 0x51, 0xf9, 0x9e, 0xf0, 0x4d, 0xff, 0x95, 0x79, 0x13, 0x76,
	0x2a, 0x0f, 0x1b, 0x1d, 0x4f, 0xc5, 0x95, 0xc3, 0xfc, 0x14, 0x2f, 0xa3, 0x85, 0x75, 0x50, 0xbe,
	0x64, 0x86, 0x63, 0xa5, 0x38, 0xa3, 0x6a, 0x21, 0x98, 0x98, 0xf1, 0x15, 0x54, 0x6e, 0x81, 0xa6,
	0x01, 0xd5, 0xb4, 0x52, 0x9a, 0x71, 0x2d, 0x47, 0xb9, 0x6d, 0xb5, 0x58, 0xb6, 0x16, 0xad, 0x8a,
	0x55, 0xfb, 0xd2, 0x41, 0x8b, 0xb7, 0x25, 0x50, 0x0d, 0x5e, 0xa7, 0xd5, 0x86, 0xf7, 0x13, 0x50,
	0xfa, 0x00, 0x26, 0xfb, 0x78, 0x4c, 0xce, 0x53, 0x30, 0xb9, 0x13, 0x4c, 0x85, 0x63, 0x31, 0xd9,
	0xd7, 0xa7, 0x00, 0xcd, 0xfd, 0x0b, 0x40, 0xf3, 0xcf, 0x0e, 0xa8, 0xf8, 0x0c, 0x80, 0x6e, 0xa1,
	0xd3, 0x07, 0xf8, 0xa8, 0x58, 0x70, 0x05, 0xb8, 0x86, 0x8a, 0x5e, 0xa7, 0xb5, 0x73, 0x24, 0xa0,
	0x22, 0xd5, 0xd1, 0x0e, 0x0b, 0xc6, 0x09, 0x5e, 0x45, 0x27, 0x37, 0x40, 0x1f, 0xa0, 0xfb, 0x3c,
	0xc1, 0x1b, 0xe8, 0x54, 0x16, 0xac, 0xc6, 0xb5, 0x97, 0x51, 0xb1, 0x0d, 0x2a, 0x09, 0x75, 0xc5,
	0x26, 0x85, 0xa5, 0x85, 0x66, 0xd9, 0xa5, 0x3a, 0x72, 0xbd, 0x4e, 0x6b, 0x3a, 0x91, 0x34, 0x3e,
	0xa3, 0x44, 0xcd, 0xaf, 0x4a, 0x08, 0x79, 0x9d, 0xd6, 0xdb, 0x20, 0x87, 0xcc, 0x07, 0xfc, 0x9d,
	0x83, 0x8a, 0x59, 0x62, 0x8c, 0x4d, 0xfc, 0x94, 0xc4, 0xea, 0x38, 0x67, 0xed, 0x73, 0x67, 0xdf,
	0xfb, 0xd3, 0xce, 0x5e, 0xd6, 0xa9, 0x36, 0x68, 0xc9, 0x60, 0x08, 0x84, 0x72, 0xe2, 0x75, 0x5a,
	0xd5, 0x7b, 0xa3, 0x03, 0x45, 0x18, 0xef, 0x09, 0x19, 0x99, 0x0b, 0x21, 0x12, 0xfa, 0x54, 0x06,
	0x8c, 0xf7, 0x09, 0x25, 0x2a, 0x06, 0x9f, 0xf5, 0x98, 0x9f, 0xfa, 0x2f, 0x13, 0x05, 0x21, 0xf8,
	0x1a, 0x02, 0xd2, 0xdd, 0x23, 0x7a, 0x00, 0x44, 0x25, 0x71, 0x1c, 0x32, 0x08, 0xc8, 0xe6, 0xba,
	0xbb, 0xe5, 0xa1, 0x42, 0xb3, 0xd1, 0xc0, 0xab, 0xe8, 0x62, 0x2e, 0x83, 0xc0, 0x43, 0xf0, 0x93,
	0xd4, 0x5d, 0x25, 0xbe, 0x0f, 0x4a, 0xf5, 0x92, 0x30, 0xdc, 0x73, 0x71, 0x05, 0x9d, 0xab, 0x9e,
	0xbd, 0x54, 0x0f, 0xa0, 0xc7, 0x38, 0xcb, 0x97, 0x8b, 0x8e, 0xbc, 0x4e, 0x6b, 0x6b, 0x05, 0x15,
	0x6e, 0x34, 0x6e, 0xe0, 0xab, 0x68, 0xa9, 0x0d, 0x3a, 0x91, 0x1c, 0x02, 0xb2, 0x3b, 0x00, 0x6e,
	0x6a, 0x49, 0x50, 0x22, 0x91, 0x3e, 0x10, 0xa6, 0x08, 0x17, 0x9a, 0xf4, 0x44, 0xc2, 0x03, 0xb7,
	0x8b, 0xd1, 0x22, 0x2a, 0xbe, 0xe5, 0x25, 0x7a, 0xd0, 0xc4, 0x45, 0x34, 0x27, 0x81, 0x06, 0x8f,
	0x7e, 0xfa, 0xed, 0x0b, 0xe7, 0x0c, 0x3e, 0x9d, 0xaf, 0x40, 0x55, 0xff, 0x30, 0xbb, 0xaf, 0x8f,
	0x1e, 0x3b, 0xd6, 0x13, 0xc7, 0x10, 0xc6, 0x3f, 0x38, 0xa8, 0x94, 0xdf, 0x0c, 0x3e, 0xe7, 0x66,
	0x8b, 0xcc, 0x1d, 0x2d, 0x32, 0xf7, 0x4e, 0xba, 0xc8, 0xaa, 0x67, 0x0f, 0x90, 0x1d, 0xdf, 0x5f,
	0xed, 0x33, 0x67, 0xdf, 0xfb, 0x2b, 0x27, 0x7a, 0x7e, 0x42, 0x34, 0x0c, 0x09, 0x1d, 0x52, 0x16,
	0xd2, 0x6e, 0x08, 0x29, 0x2c, 0x55, 0xbd, 0x79, 0x2c, 0xdd, 0x30, 0x34, 0x9d, 0x4d, 0xc7, 0xb9,
	0x5b, 0xad, 0x0c, 0xe4, 0xdd, 0x63, 0x41, 0x5e, 0x46, 0xb5, 0x2a, 0x99, 0x01, 0x79, 0x48, 0xf0,
	0x7f, 0x09, 0x15, 0xe1, 0xf2, 0x08, 0x6a, 0xc6, 0xb2, 0x61, 0xe1, 0x47, 0x0e, 0x3a, 0x31, 0x7e,
	0x61, 0xf8, 0xff, 0x86, 0xd9, 0xe1, 0x8d, 0x54, 0x3d, 0x77, 0xf8, 0x38, 0x87, 0xf9, 0xbd, 0xbd,
	0xef, 0x7d, 0x9d, 0xc3, 0x3c, 0x99, 0x99, 0x47, 0xc3, 0xf9, 0x62, 0xf6, 0xa9, 0x08, 0x25, 0x1c,
	0x76, 0xd3, 0x23, 0x42, 0x79, 0x40, 0xa4, 0xe9, 0x41, 0x11, 0xa6, 0x15, 0x61, 0x81, 0xbb, 0xd5,
	0x4e, 0x61, 0xad, 0xe0, 0x37, 0xd0, 0x0b, 0x1d, 0x49, 0xb9, 0xa2, 0xbe, 0x81, 0xec, 0x9b, 0xd8,
	0xc3, 0xbc, 0xae, 0xa0, 0xcb, 0xd5, 0xda, 0x0c, 0xaf, 0x19, 0x55, 0xdd, 0x33, 0xe8, 0xf4, 0xb8,
	0xfd, 0x12, 0x9a, 0xdf, 0x95, 0x4c, 0x83, 0xe9, 0xff, 0xe4, 0xaa, 0x7d, 0xb5, 0x76, 0x08, 0x41,
	0x3a, 0x4e, 0xd5, 0xc2, 0x63, 0xc7, 0x5a, 0xfb, 0xa4, 0xb8, 0xef, 0xfd, 0x3a, 0x8f, 0x0a, 0x4d,
	0xb7, 0x81, 0xb7, 0x50, 0x29, 0x55, 0xec, 0x6d, 0x6f, 0xe2, 0x57, 0xb6, 0xa5, 0x18, 0xb2, 0x00,
	0x54, 0xae, 0x2a, 0x6f, 0x83, 0x06, 0x44, 0xc4, 0x20, 0xb3, 0xbf, 0xa8, 0x44, 0x64, 0xf7, 0x91,
	0xc6, 0x8c, 0xee, 0xc4, 0x6d, 0xce, 0xaf, 0xb8, 0x0d, 0xb7, 0x71, 0xd5, 0x76, 0x9a, 0x8b, 0x34,
	0x7d, 0x68, 0xd9, 0xc2, 0xac, 0xbf, 0xa7, 0x04, 0x5f, 0x9d, 0x39, 0x69, 0x3f, 0xff, 0x3d, 0xb7,
	0xef, 0xa2, 0xc2, 0xcb, 0x8d, 0x06, 0xbe, 0x85, 0x5e, 0x9b, 0x0e, 0xa1, 0x9c, 0x24, 0x1c, 0x1e,
	0xc6, 0xd9, 0x63, 0x07, 0x29, 0x85, 0x24, 0xc2, 0xf7, 0x13, 0x09, 0xc1, 0x48, 0xa9, 0x02, 0x39,
	0x04, 0x49, 0x14, 0x0b, 0xc0, 0x6d, 0xef, 0xa4, 0xa5, 0x1b, 0xf8, 0x5d, 0xf4, 0xce, 0x51, 0xa5,
	0xb3, 0x19, 0xee, 0x8a, 0x60, 0x2f, 0x2d, 0x1f, 0xd1, 0x30, 0x7b, 0x0c, 0x69, 0x6a, 0x21, 0x49,
	0x20, 0x20, 0xd3, 0x14, 0x51, 0xed, 0x0f, 0x4c, 0xc8, 0xb8, 0x72, 0x1e, 0xeb, 0xb6, 0xef, 0xa5,
	0x05, 0x56, 0xf0, 0x1d, 0x74, 0xfb, 0xe9, 0x05, 0xc6, 0x89, 0x7c, 0xc1, 0x35, 0x65, 0x5c, 0x19,
	0x6b, 0xa2, 0x40, 0xbe, 0x64, 0xd0, 0x07, 0xc0, 0x35, 0xa3, 0xa1, 0x72, 0xdb, 0xdb, 0x69, 0xb6,
	0xeb, 0x78, 0x13, 0x6d, 0xcc, 0x66, 0x4b, 0xfd, 0x27, 0xa9, 0x06, 0x74, 0x08, 0x24, 0x06, 0x19,
	0x31, 0xa5, 0xd2, 0xb9, 0xd2, 0x82, 0x50, 0x33, 0x51, 0x53, 0x54, 0xdd, 0xfb, 0x7f, 0xd8, 0xe8,
	0x77, 0x7b, 0x3c, 0x33, 0xbf, 0xd8, 0xe5, 0x02, 0xfe, 0xd4, 0xf6, 0x72, 0x4f, 0x41, 0xf4, 0x64,
	0x34, 0xd5, 0x38, 0x4a, 0xa5, 0x34, 0x24, 0x28, 0x2d, 0x99, 0x69, 0x36, 0xcd, 0x9d, 0xe8, 0x41,
	0xaa, 0xd2, 0x37, 0xc3, 0x9b, 0x4a, 0x51, 0x2e, 0xe9, 0xa4, 0x2b, 0x62, 0x62, 0x48, 0x65, 0xc4,
	0x52, 0x98, 0xd4, 0x3d, 0x11, 0x86, 0x62, 0x37, 0x13, 0x93, 0x96, 0x16, 0x92, 0x7d, 0x90, 0x79,
	0xdc, 0x16, 0x01, 0x90, 0x5e, 0x28, 0x76, 0xdd, 0xa5, 0xb9, 0x66, 0x36, 0xb1, 0x89, 0x1e, 0xac,
	0x9e, 0x30, 0xbf, 0xbc, 0xc4, 0x03, 0xe0, 0x6b, 0xab, 0xe8, 0x42, 0x3e, 0xda, 0xf8, 0xcc, 0x86,
	0xa4, 0x5c, 0x2b, 0x62, 0xbe, 0xf2, 0x06, 0x51, 0x35, 0x7b, 0xf7, 0x18, 0xe7, 0x46, 0x33, 0xb1,
	0x99, 0xed, 0x75, 0x7b, 0xdb, 0xba, 0x9f, 0xfe, 0x9c, 0xfc, 0xd8, 0xb6, 0x1e, 0xdb, 0xd6, 0x13,
	0xdb, 0xfa, 0xd6, 0xb6, 0x7e, 0xb6, 0xad, 0xbf, 0x6d, 0xeb, 0x47, 0xc7, 0xea, 0x16, 0xcd, 0x1e,
	0xbd, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x39, 0xed, 0x39, 0xa0, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ATMServiceClient is the client API for ATMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ATMServiceClient interface {
	// GetATM retrieves the details regarding an specific ATM.
	GetATM(ctx context.Context, in *GetATMRequest, opts ...grpc.CallOption) (*ATM, error)
	// GetATMs retrieves information regarding all the available ATMs.
	GetATMs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetATMsResponse, error)
	// CreateATM creates a new ATM and returns its id.
	CreateATM(ctx context.Context, in *CreateATMRequest, opts ...grpc.CallOption) (*CreateATMResponse, error)
}

type aTMServiceClient struct {
	cc *grpc.ClientConn
}

func NewATMServiceClient(cc *grpc.ClientConn) ATMServiceClient {
	return &aTMServiceClient{cc}
}

func (c *aTMServiceClient) GetATM(ctx context.Context, in *GetATMRequest, opts ...grpc.CallOption) (*ATM, error) {
	out := new(ATM)
	err := c.cc.Invoke(ctx, "/atm.ATMService/GetATM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aTMServiceClient) GetATMs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetATMsResponse, error) {
	out := new(GetATMsResponse)
	err := c.cc.Invoke(ctx, "/atm.ATMService/GetATMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aTMServiceClient) CreateATM(ctx context.Context, in *CreateATMRequest, opts ...grpc.CallOption) (*CreateATMResponse, error) {
	out := new(CreateATMResponse)
	err := c.cc.Invoke(ctx, "/atm.ATMService/CreateATM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ATMServiceServer is the server API for ATMService service.
type ATMServiceServer interface {
	// GetATM retrieves the details regarding an specific ATM.
	GetATM(context.Context, *GetATMRequest) (*ATM, error)
	// GetATMs retrieves information regarding all the available ATMs.
	GetATMs(context.Context, *empty.Empty) (*GetATMsResponse, error)
	// CreateATM creates a new ATM and returns its id.
	CreateATM(context.Context, *CreateATMRequest) (*CreateATMResponse, error)
}

// UnimplementedATMServiceServer can be embedded to have forward compatible implementations.
type UnimplementedATMServiceServer struct {
}

func (*UnimplementedATMServiceServer) GetATM(ctx context.Context, req *GetATMRequest) (*ATM, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetATM not implemented")
}
func (*UnimplementedATMServiceServer) GetATMs(ctx context.Context, req *empty.Empty) (*GetATMsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetATMs not implemented")
}
func (*UnimplementedATMServiceServer) CreateATM(ctx context.Context, req *CreateATMRequest) (*CreateATMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateATM not implemented")
}

func RegisterATMServiceServer(s *grpc.Server, srv ATMServiceServer) {
	s.RegisterService(&_ATMService_serviceDesc, srv)
}

func _ATMService_GetATM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetATMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ATMServiceServer).GetATM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atm.ATMService/GetATM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ATMServiceServer).GetATM(ctx, req.(*GetATMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ATMService_GetATMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ATMServiceServer).GetATMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atm.ATMService/GetATMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ATMServiceServer).GetATMs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ATMService_CreateATM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateATMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ATMServiceServer).CreateATM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atm.ATMService/CreateATM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ATMServiceServer).CreateATM(ctx, req.(*CreateATMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ATMService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atm.ATMService",
	HandlerType: (*ATMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetATM",
			Handler:    _ATMService_GetATM_Handler,
		},
		{
			MethodName: "GetATMs",
			Handler:    _ATMService_GetATMs_Handler,
		},
		{
			MethodName: "CreateATM",
			Handler:    _ATMService_CreateATM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/atm/all.proto",
}
