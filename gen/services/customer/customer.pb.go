// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: github.com/Ultimate-Super-WebDev-Corp/gateway/services/customer/customer.proto

package customer

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CustomerMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *CustomerMsg) Reset() {
	*x = CustomerMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerMsg) ProtoMessage() {}

func (x *CustomerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerMsg.ProtoReflect.Descriptor instead.
func (*CustomerMsg) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerMsg) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassword string `protobuf:"bytes,1,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string       `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
	Customer *CustomerMsg `protobuf:"bytes,2,opt,name=Customer,proto3" json:"Customer,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateRequest) GetCustomer() *CustomerMsg {
	if x != nil {
		return x.Customer
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto protoreflect.FileDescriptor

var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6c, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x53, 0x75, 0x70, 0x65, 0x72, 0x2d, 0x57, 0x65, 0x62, 0x44,
	0x65, 0x76, 0x2d, 0x43, 0x6f, 0x72, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x71, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x46,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe2,
	0xdf, 0x1f, 0x2c, 0x0a, 0x2a, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2e, 0x5f, 0x25,
	0x2b, 0x5c, 0x2d, 0x5d, 0x2b, 0x40, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2e, 0x5c, 0x2d,
	0x5d, 0x2b, 0x5c, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x7b, 0x32, 0x2c, 0x34, 0x7d, 0x24, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x48, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x4e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x0a, 0x07, 0x5e, 0x2e, 0x7b, 0x34, 0x2c, 0x7d, 0x24, 0x52,
	0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x75, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x0a, 0x07, 0x5e, 0x2e, 0x7b, 0x34, 0x2c, 0x7d, 0x24, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x30, 0xe2, 0xdf, 0x1f, 0x2c, 0x0a, 0x2a, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x2e, 0x5f, 0x25, 0x2b, 0x5c, 0x2d, 0x5d, 0x2b, 0x40, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x2e, 0x5c, 0x2d, 0x5d, 0x2b, 0x5c, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x7b,
	0x32, 0x2c, 0x34, 0x7d, 0x24, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xe2, 0xdf, 0x1f, 0x09, 0x0a, 0x07, 0x5e, 0x2e, 0x7b, 0x34, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0xfd, 0x02, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x55, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x2d, 0x57, 0x65, 0x62, 0x44, 0x65, 0x76, 0x2d, 0x43, 0x6f, 0x72, 0x70, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescOnce sync.Once
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescData = file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDesc
)

func file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescGZIP() []byte {
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescOnce.Do(func() {
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescData)
	})
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDescData
}

var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_goTypes = []interface{}{
	(*CustomerMsg)(nil),           // 0: customer.CustomerMsg
	(*ChangePasswordRequest)(nil), // 1: customer.ChangePasswordRequest
	(*CreateRequest)(nil),         // 2: customer.CreateRequest
	(*LoginRequest)(nil),          // 3: customer.LoginRequest
	(*UpdateRequest)(nil),         // 4: customer.UpdateRequest
	(*empty.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_depIdxs = []int32{
	0, // 0: customer.CreateRequest.Customer:type_name -> customer.CustomerMsg
	2, // 1: customer.Customer.Create:input_type -> customer.CreateRequest
	5, // 2: customer.Customer.Get:input_type -> google.protobuf.Empty
	3, // 3: customer.Customer.Login:input_type -> customer.LoginRequest
	5, // 4: customer.Customer.Logout:input_type -> google.protobuf.Empty
	1, // 5: customer.Customer.ChangePassword:input_type -> customer.ChangePasswordRequest
	4, // 6: customer.Customer.Update:input_type -> customer.UpdateRequest
	0, // 7: customer.Customer.Create:output_type -> customer.CustomerMsg
	0, // 8: customer.Customer.Get:output_type -> customer.CustomerMsg
	0, // 9: customer.Customer.Login:output_type -> customer.CustomerMsg
	5, // 10: customer.Customer.Logout:output_type -> google.protobuf.Empty
	5, // 11: customer.Customer.ChangePassword:output_type -> google.protobuf.Empty
	0, // 12: customer.Customer.Update:output_type -> customer.CustomerMsg
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_init()
}
func file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_init() {
	if File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_goTypes,
		DependencyIndexes: file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_depIdxs,
		MessageInfos:      file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_msgTypes,
	}.Build()
	File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto = out.File
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_rawDesc = nil
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_goTypes = nil
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_customer_customer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CustomerMsg, error)
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CustomerMsg, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CustomerMsg, error)
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*CustomerMsg, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CustomerMsg, error) {
	out := new(CustomerMsg)
	err := c.cc.Invoke(ctx, "/customer.Customer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CustomerMsg, error) {
	out := new(CustomerMsg)
	err := c.cc.Invoke(ctx, "/customer.Customer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CustomerMsg, error) {
	out := new(CustomerMsg)
	err := c.cc.Invoke(ctx, "/customer.Customer/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/customer.Customer/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/customer.Customer/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*CustomerMsg, error) {
	out := new(CustomerMsg)
	err := c.cc.Invoke(ctx, "/customer.Customer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	Create(context.Context, *CreateRequest) (*CustomerMsg, error)
	Get(context.Context, *empty.Empty) (*CustomerMsg, error)
	Login(context.Context, *LoginRequest) (*CustomerMsg, error)
	Logout(context.Context, *empty.Empty) (*empty.Empty, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*empty.Empty, error)
	Update(context.Context, *UpdateRequest) (*CustomerMsg, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) Create(context.Context, *CreateRequest) (*CustomerMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCustomerServer) Get(context.Context, *empty.Empty) (*CustomerMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCustomerServer) Login(context.Context, *LoginRequest) (*CustomerMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedCustomerServer) Logout(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedCustomerServer) ChangePassword(context.Context, *ChangePasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (*UnimplementedCustomerServer) Update(context.Context, *UpdateRequest) (*CustomerMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Customer_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Customer_Get_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Customer_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Customer_Logout_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Customer_ChangePassword_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Customer_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Ultimate-Super-WebDev-Corp/gateway/services/customer/customer.proto",
}
