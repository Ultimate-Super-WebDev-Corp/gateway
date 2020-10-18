// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: github.com/Ultimate-Super-WebDev-Corp/gateway/services/product/product.proto

package product

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

type SearchByUUIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUIDs []string `protobuf:"bytes,1,rep,name=UUIDs,proto3" json:"UUIDs,omitempty"`
}

func (x *SearchByUUIDsRequest) Reset() {
	*x = SearchByUUIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchByUUIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchByUUIDsRequest) ProtoMessage() {}

func (x *SearchByUUIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchByUUIDsRequest.ProtoReflect.Descriptor instead.
func (*SearchByUUIDsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *SearchByUUIDsRequest) GetUUIDs() []string {
	if x != nil {
		return x.UUIDs
	}
	return nil
}

type ProductMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Brand       string   `protobuf:"bytes,2,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Images      []string `protobuf:"bytes,4,rep,name=Images,proto3" json:"Images,omitempty"`
	Country     string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *ProductMsg) Reset() {
	*x = ProductMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductMsg) ProtoMessage() {}

func (x *ProductMsg) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductMsg.ProtoReflect.Descriptor instead.
func (*ProductMsg) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductMsg) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *ProductMsg) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductMsg) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *ProductMsg) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type ProductWithID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Product *ProductMsg `protobuf:"bytes,2,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *ProductWithID) Reset() {
	*x = ProductWithID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductWithID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductWithID) ProtoMessage() {}

func (x *ProductWithID) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductWithID.ProtoReflect.Descriptor instead.
func (*ProductWithID) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductWithID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductWithID) GetProduct() *ProductMsg {
	if x != nil {
		return x.Product
	}
	return nil
}

type GetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetByIDRequest) Reset() {
	*x = GetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRequest) ProtoMessage() {}

func (x *GetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto protoreflect.FileDescriptor

var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6c, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x53, 0x75, 0x70, 0x65, 0x72, 0x2d, 0x57, 0x65, 0x62, 0x44,
	0x65, 0x76, 0x2d, 0x43, 0x6f, 0x72, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x14,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x55, 0x55, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x60, 0x01, 0x52, 0x05, 0x55, 0x55, 0x49,
	0x44, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x4e, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x02, 0x49,
	0x64, 0x32, 0xca, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x48, 0x0a,
	0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x73, 0x12, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x79, 0x55, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x49, 0x44, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x22, 0x00, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6c, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x53, 0x75, 0x70, 0x65, 0x72, 0x2d, 0x57, 0x65, 0x62, 0x44,
	0x65, 0x76, 0x2d, 0x43, 0x6f, 0x72, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescOnce sync.Once
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescData = file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDesc
)

func file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescGZIP() []byte {
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescOnce.Do(func() {
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescData)
	})
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDescData
}

var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_goTypes = []interface{}{
	(*SearchByUUIDsRequest)(nil), // 0: product.SearchByUUIDsRequest
	(*ProductMsg)(nil),           // 1: product.ProductMsg
	(*ProductWithID)(nil),        // 2: product.ProductWithID
	(*GetByIDRequest)(nil),       // 3: product.GetByIDRequest
	(*empty.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_depIdxs = []int32{
	1, // 0: product.ProductWithID.Product:type_name -> product.ProductMsg
	0, // 1: product.Product.SearchByUUIDs:input_type -> product.SearchByUUIDsRequest
	1, // 2: product.Product.Create:input_type -> product.ProductMsg
	3, // 3: product.Product.GetByID:input_type -> product.GetByIDRequest
	2, // 4: product.Product.SearchByUUIDs:output_type -> product.ProductWithID
	4, // 5: product.Product.Create:output_type -> google.protobuf.Empty
	2, // 6: product.Product.GetByID:output_type -> product.ProductWithID
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_init() }
func file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_init() {
	if File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchByUUIDsRequest); i {
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
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductMsg); i {
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
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductWithID); i {
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
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRequest); i {
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
			RawDescriptor: file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_goTypes,
		DependencyIndexes: file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_depIdxs,
		MessageInfos:      file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_msgTypes,
	}.Build()
	File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto = out.File
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_rawDesc = nil
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_goTypes = nil
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_product_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	SearchByUUIDs(ctx context.Context, in *SearchByUUIDsRequest, opts ...grpc.CallOption) (*ProductWithID, error)
	Create(ctx context.Context, in *ProductMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*ProductWithID, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) SearchByUUIDs(ctx context.Context, in *SearchByUUIDsRequest, opts ...grpc.CallOption) (*ProductWithID, error) {
	out := new(ProductWithID)
	err := c.cc.Invoke(ctx, "/product.Product/SearchByUUIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Create(ctx context.Context, in *ProductMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.Product/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*ProductWithID, error) {
	out := new(ProductWithID)
	err := c.cc.Invoke(ctx, "/product.Product/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	SearchByUUIDs(context.Context, *SearchByUUIDsRequest) (*ProductWithID, error)
	Create(context.Context, *ProductMsg) (*empty.Empty, error)
	GetByID(context.Context, *GetByIDRequest) (*ProductWithID, error)
}

// UnimplementedProductServer can be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (*UnimplementedProductServer) SearchByUUIDs(context.Context, *SearchByUUIDsRequest) (*ProductWithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByUUIDs not implemented")
}
func (*UnimplementedProductServer) Create(context.Context, *ProductMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProductServer) GetByID(context.Context, *GetByIDRequest) (*ProductWithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_SearchByUUIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchByUUIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).SearchByUUIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/SearchByUUIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).SearchByUUIDs(ctx, req.(*SearchByUUIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Create(ctx, req.(*ProductMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchByUUIDs",
			Handler:    _Product_SearchByUUIDs_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Product_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Product_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Ultimate-Super-WebDev-Corp/gateway/services/product/product.proto",
}
