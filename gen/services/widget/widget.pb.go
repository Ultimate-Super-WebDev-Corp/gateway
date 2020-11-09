// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: github.com/Ultimate-Super-WebDev-Corp/gateway/services/widget/widget.proto

package widget

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type HtmlBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *HtmlBody) Reset() {
	*x = HtmlBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HtmlBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HtmlBody) ProtoMessage() {}

func (x *HtmlBody) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HtmlBody.ProtoReflect.Descriptor instead.
func (*HtmlBody) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescGZIP(), []int{0}
}

func (x *HtmlBody) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ProductPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID uint64 `protobuf:"varint,5,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *ProductPriceRequest) Reset() {
	*x = ProductPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPriceRequest) ProtoMessage() {}

func (x *ProductPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPriceRequest.ProtoReflect.Descriptor instead.
func (*ProductPriceRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescGZIP(), []int{1}
}

func (x *ProductPriceRequest) GetProductID() uint64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

var File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto protoreflect.FileDescriptor

var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6c, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x53, 0x75, 0x70, 0x65, 0x72, 0x2d, 0x57, 0x65, 0x62, 0x44,
	0x65, 0x76, 0x2d, 0x43, 0x6f, 0x72, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2f,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x33, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x32, 0x81, 0x01, 0x0a, 0x06, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x48,
	0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e,
	0x48, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x2d, 0x53, 0x75, 0x70, 0x65, 0x72, 0x2d, 0x57, 0x65, 0x62, 0x44, 0x65, 0x76, 0x2d, 0x43,
	0x6f, 0x72, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescOnce sync.Once
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescData = file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDesc
)

func file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescGZIP() []byte {
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescOnce.Do(func() {
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescData)
	})
	return file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDescData
}

var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_goTypes = []interface{}{
	(*HtmlBody)(nil),            // 0: widget.HtmlBody
	(*ProductPriceRequest)(nil), // 1: widget.ProductPriceRequest
	(*empty.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_depIdxs = []int32{
	2, // 0: widget.Widget.MainPage:input_type -> google.protobuf.Empty
	1, // 1: widget.Widget.ProductPrice:input_type -> widget.ProductPriceRequest
	0, // 2: widget.Widget.MainPage:output_type -> widget.HtmlBody
	0, // 3: widget.Widget.ProductPrice:output_type -> widget.HtmlBody
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_init() }
func file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_init() {
	if File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HtmlBody); i {
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
		file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPriceRequest); i {
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
			RawDescriptor: file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_goTypes,
		DependencyIndexes: file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_depIdxs,
		MessageInfos:      file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_msgTypes,
	}.Build()
	File_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto = out.File
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_rawDesc = nil
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_goTypes = nil
	file_github_com_Ultimate_Super_WebDev_Corp_gateway_services_widget_widget_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WidgetClient is the client API for Widget service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WidgetClient interface {
	MainPage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HtmlBody, error)
	ProductPrice(ctx context.Context, in *ProductPriceRequest, opts ...grpc.CallOption) (*HtmlBody, error)
}

type widgetClient struct {
	cc grpc.ClientConnInterface
}

func NewWidgetClient(cc grpc.ClientConnInterface) WidgetClient {
	return &widgetClient{cc}
}

func (c *widgetClient) MainPage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HtmlBody, error) {
	out := new(HtmlBody)
	err := c.cc.Invoke(ctx, "/widget.Widget/MainPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *widgetClient) ProductPrice(ctx context.Context, in *ProductPriceRequest, opts ...grpc.CallOption) (*HtmlBody, error) {
	out := new(HtmlBody)
	err := c.cc.Invoke(ctx, "/widget.Widget/ProductPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WidgetServer is the server API for Widget service.
type WidgetServer interface {
	MainPage(context.Context, *empty.Empty) (*HtmlBody, error)
	ProductPrice(context.Context, *ProductPriceRequest) (*HtmlBody, error)
}

// UnimplementedWidgetServer can be embedded to have forward compatible implementations.
type UnimplementedWidgetServer struct {
}

func (*UnimplementedWidgetServer) MainPage(context.Context, *empty.Empty) (*HtmlBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MainPage not implemented")
}
func (*UnimplementedWidgetServer) ProductPrice(context.Context, *ProductPriceRequest) (*HtmlBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductPrice not implemented")
}

func RegisterWidgetServer(s *grpc.Server, srv WidgetServer) {
	s.RegisterService(&_Widget_serviceDesc, srv)
}

func _Widget_MainPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WidgetServer).MainPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/widget.Widget/MainPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WidgetServer).MainPage(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Widget_ProductPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WidgetServer).ProductPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/widget.Widget/ProductPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WidgetServer).ProductPrice(ctx, req.(*ProductPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Widget_serviceDesc = grpc.ServiceDesc{
	ServiceName: "widget.Widget",
	HandlerType: (*WidgetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MainPage",
			Handler:    _Widget_MainPage_Handler,
		},
		{
			MethodName: "ProductPrice",
			Handler:    _Widget_ProductPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Ultimate-Super-WebDev-Corp/gateway/services/widget/widget.proto",
}