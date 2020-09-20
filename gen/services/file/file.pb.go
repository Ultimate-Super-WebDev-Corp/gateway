// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: services/file/file.proto

package file

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type FileType int32

const (
	FileType_UNDEFINED FileType = 0
	FileType_JPEG      FileType = 1
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "UNDEFINED",
		1: "JPEG",
	}
	FileType_value = map[string]int32{
		"UNDEFINED": 0,
		"JPEG":      1,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_services_file_file_proto_enumTypes[0].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_services_file_file_proto_enumTypes[0]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_services_file_file_proto_rawDescGZIP(), []int{0}
}

type FileUUIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUIDs []string `protobuf:"bytes,1,rep,name=UUIDs,proto3" json:"UUIDs,omitempty"`
}

func (x *FileUUIDs) Reset() {
	*x = FileUUIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUUIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUUIDs) ProtoMessage() {}

func (x *FileUUIDs) ProtoReflect() protoreflect.Message {
	mi := &file_services_file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUUIDs.ProtoReflect.Descriptor instead.
func (*FileUUIDs) Descriptor() ([]byte, []int) {
	return file_services_file_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileUUIDs) GetUUIDs() []string {
	if x != nil {
		return x.UUIDs
	}
	return nil
}

type FileURLs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URLs []string `protobuf:"bytes,1,rep,name=URLs,proto3" json:"URLs,omitempty"`
}

func (x *FileURLs) Reset() {
	*x = FileURLs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileURLs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileURLs) ProtoMessage() {}

func (x *FileURLs) ProtoReflect() protoreflect.Message {
	mi := &file_services_file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileURLs.ProtoReflect.Descriptor instead.
func (*FileURLs) Descriptor() ([]byte, []int) {
	return file_services_file_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileURLs) GetURLs() []string {
	if x != nil {
		return x.URLs
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OneOfChunk:
	//	*Chunk_Chunk
	//	*Chunk_Meta
	OneOfChunk isChunk_OneOfChunk `protobuf_oneof:"oneOfChunk"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_services_file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_services_file_file_proto_rawDescGZIP(), []int{2}
}

func (m *Chunk) GetOneOfChunk() isChunk_OneOfChunk {
	if m != nil {
		return m.OneOfChunk
	}
	return nil
}

func (x *Chunk) GetChunk() []byte {
	if x, ok := x.GetOneOfChunk().(*Chunk_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (x *Chunk) GetMeta() *FileMetadata {
	if x, ok := x.GetOneOfChunk().(*Chunk_Meta); ok {
		return x.Meta
	}
	return nil
}

type isChunk_OneOfChunk interface {
	isChunk_OneOfChunk()
}

type Chunk_Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=Chunk,proto3,oneof"`
}

type Chunk_Meta struct {
	Meta *FileMetadata `protobuf:"bytes,2,opt,name=Meta,proto3,oneof"`
}

func (*Chunk_Chunk) isChunk_OneOfChunk() {}

func (*Chunk_Meta) isChunk_OneOfChunk() {}

type FileUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *FileUploadResponse) Reset() {
	*x = FileUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResponse) ProtoMessage() {}

func (x *FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResponse.ProtoReflect.Descriptor instead.
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_services_file_file_proto_rawDescGZIP(), []int{3}
}

func (x *FileUploadResponse) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type FileType `protobuf:"varint,1,opt,name=Type,proto3,enum=file.FileType" json:"Type,omitempty"`
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_services_file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_services_file_file_proto_rawDescGZIP(), []int{4}
}

func (x *FileMetadata) GetType() FileType {
	if x != nil {
		return x.Type
	}
	return FileType_UNDEFINED
}

var File_services_file_file_proto protoreflect.FileDescriptor

var file_services_file_file_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69,
	0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x55, 0x49, 0x44, 0x73, 0x12, 0x1c, 0x0a, 0x05, 0x55, 0x55, 0x49, 0x44, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x60, 0x01, 0x52, 0x05, 0x55, 0x55,
	0x49, 0x44, 0x73, 0x22, 0x1e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x52, 0x4c, 0x73, 0x22, 0x57, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x05,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x28, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x0c,
	0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x28, 0x0a, 0x12,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x23, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x45, 0x47, 0x10, 0x01, 0x32,
	0x6d, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x0b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x18,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x30, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x0f, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x55, 0x49, 0x44, 0x73, 0x1a, 0x0e, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_file_file_proto_rawDescOnce sync.Once
	file_services_file_file_proto_rawDescData = file_services_file_file_proto_rawDesc
)

func file_services_file_file_proto_rawDescGZIP() []byte {
	file_services_file_file_proto_rawDescOnce.Do(func() {
		file_services_file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_file_file_proto_rawDescData)
	})
	return file_services_file_file_proto_rawDescData
}

var file_services_file_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_file_file_proto_goTypes = []interface{}{
	(FileType)(0),              // 0: file.FileType
	(*FileUUIDs)(nil),          // 1: file.FileUUIDs
	(*FileURLs)(nil),           // 2: file.FileURLs
	(*Chunk)(nil),              // 3: file.Chunk
	(*FileUploadResponse)(nil), // 4: file.FileUploadResponse
	(*FileMetadata)(nil),       // 5: file.FileMetadata
}
var file_services_file_file_proto_depIdxs = []int32{
	5, // 0: file.Chunk.Meta:type_name -> file.FileMetadata
	0, // 1: file.FileMetadata.Type:type_name -> file.FileType
	3, // 2: file.File.Upload:input_type -> file.Chunk
	1, // 3: file.File.GetFileURLs:input_type -> file.FileUUIDs
	4, // 4: file.File.Upload:output_type -> file.FileUploadResponse
	2, // 5: file.File.GetFileURLs:output_type -> file.FileURLs
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_file_file_proto_init() }
func file_services_file_file_proto_init() {
	if File_services_file_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUUIDs); i {
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
		file_services_file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileURLs); i {
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
		file_services_file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_services_file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadResponse); i {
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
		file_services_file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadata); i {
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
	file_services_file_file_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Chunk_Chunk)(nil),
		(*Chunk_Meta)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_file_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_file_file_proto_goTypes,
		DependencyIndexes: file_services_file_file_proto_depIdxs,
		EnumInfos:         file_services_file_file_proto_enumTypes,
		MessageInfos:      file_services_file_file_proto_msgTypes,
	}.Build()
	File_services_file_file_proto = out.File
	file_services_file_file_proto_rawDesc = nil
	file_services_file_file_proto_goTypes = nil
	file_services_file_file_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (File_UploadClient, error)
	GetFileURLs(ctx context.Context, in *FileUUIDs, opts ...grpc.CallOption) (*FileURLs, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) Upload(ctx context.Context, opts ...grpc.CallOption) (File_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_File_serviceDesc.Streams[0], "/file.File/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadClient{stream}
	return x, nil
}

type File_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type fileUploadClient struct {
	grpc.ClientStream
}

func (x *fileUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadClient) CloseAndRecv() (*FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileClient) GetFileURLs(ctx context.Context, in *FileUUIDs, opts ...grpc.CallOption) (*FileURLs, error) {
	out := new(FileURLs)
	err := c.cc.Invoke(ctx, "/file.File/GetFileURLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
type FileServer interface {
	Upload(File_UploadServer) error
	GetFileURLs(context.Context, *FileUUIDs) (*FileURLs, error)
}

// UnimplementedFileServer can be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (*UnimplementedFileServer) Upload(File_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedFileServer) GetFileURLs(context.Context, *FileUUIDs) (*FileURLs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileURLs not implemented")
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServer).Upload(&fileUploadServer{stream})
}

type File_UploadServer interface {
	SendAndClose(*FileUploadResponse) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type fileUploadServer struct {
	grpc.ServerStream
}

func (x *fileUploadServer) SendAndClose(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _File_GetFileURLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUUIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetFileURLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.File/GetFileURLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetFileURLs(ctx, req.(*FileUUIDs))
	}
	return interceptor(ctx, in, info, handler)
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileURLs",
			Handler:    _File_GetFileURLs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _File_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "services/file/file.proto",
}
