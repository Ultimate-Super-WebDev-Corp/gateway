// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: services/comment/comment.proto

package comment

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

type Rating int32

const (
	Rating_ZERO_STARS  Rating = 0
	Rating_ONE_STARS   Rating = 1
	Rating_TWO_STARS   Rating = 2
	Rating_THREE_STARS Rating = 3
	Rating_FOUR_STARS  Rating = 4
	Rating_FIVE_STARS  Rating = 5
)

// Enum value maps for Rating.
var (
	Rating_name = map[int32]string{
		0: "ZERO_STARS",
		1: "ONE_STARS",
		2: "TWO_STARS",
		3: "THREE_STARS",
		4: "FOUR_STARS",
		5: "FIVE_STARS",
	}
	Rating_value = map[string]int32{
		"ZERO_STARS":  0,
		"ONE_STARS":   1,
		"TWO_STARS":   2,
		"THREE_STARS": 3,
		"FOUR_STARS":  4,
		"FIVE_STARS":  5,
	}
)

func (x Rating) Enum() *Rating {
	p := new(Rating)
	*p = x
	return p
}

func (x Rating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rating) Descriptor() protoreflect.EnumDescriptor {
	return file_services_comment_comment_proto_enumTypes[0].Descriptor()
}

func (Rating) Type() protoreflect.EnumType {
	return &file_services_comment_comment_proto_enumTypes[0]
}

func (x Rating) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rating.Descriptor instead.
func (Rating) EnumDescriptor() ([]byte, []int) {
	return file_services_comment_comment_proto_rawDescGZIP(), []int{0}
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Rating    Rating `protobuf:"varint,3,opt,name=Rating,proto3,enum=comment.Rating" json:"Rating,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_comment_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_comment_comment_proto_msgTypes[0]
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
	return file_services_comment_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateRequest) GetRating() Rating {
	if x != nil {
		return x.Rating
	}
	return Rating_ZERO_STARS
}

type CreateFromSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Source    string `protobuf:"bytes,3,opt,name=Source,proto3" json:"Source,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Rating    Rating `protobuf:"varint,6,opt,name=Rating,proto3,enum=comment.Rating" json:"Rating,omitempty"`
}

func (x *CreateFromSourceRequest) Reset() {
	*x = CreateFromSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_comment_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFromSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFromSourceRequest) ProtoMessage() {}

func (x *CreateFromSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_comment_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFromSourceRequest.ProtoReflect.Descriptor instead.
func (*CreateFromSourceRequest) Descriptor() ([]byte, []int) {
	return file_services_comment_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFromSourceRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateFromSourceRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateFromSourceRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreateFromSourceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFromSourceRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateFromSourceRequest) GetRating() Rating {
	if x != nil {
		return x.Rating
	}
	return Rating_ZERO_STARS
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Offset    uint64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     uint64 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_comment_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_comment_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_services_comment_comment_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ListRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments   []*CommentMsg `protobuf:"bytes,1,rep,name=Comments,proto3" json:"Comments,omitempty"`
	NextOffset uint64        `protobuf:"varint,2,opt,name=NextOffset,proto3" json:"NextOffset,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_comment_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_comment_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_services_comment_comment_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetComments() []*CommentMsg {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListResponse) GetNextOffset() uint64 {
	if x != nil {
		return x.NextOffset
	}
	return 0
}

type CommentMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Source    string `protobuf:"bytes,3,opt,name=Source,proto3" json:"Source,omitempty"`
	Rating    Rating `protobuf:"varint,4,opt,name=Rating,proto3,enum=comment.Rating" json:"Rating,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *CommentMsg) Reset() {
	*x = CommentMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_comment_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentMsg) ProtoMessage() {}

func (x *CommentMsg) ProtoReflect() protoreflect.Message {
	mi := &file_services_comment_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentMsg.ProtoReflect.Descriptor instead.
func (*CommentMsg) Descriptor() ([]byte, []int) {
	return file_services_comment_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentMsg) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CommentMsg) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CommentMsg) GetRating() Rating {
	if x != nil {
		return x.Rating
	}
	return Rating_ZERO_STARS
}

func (x *CommentMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommentMsg) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CommentMsg) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_services_comment_comment_proto protoreflect.FileDescriptor

var file_services_comment_comment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xce, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10,
	0x00, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x61, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5f,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0xb9, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1a,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x67, 0x0a, 0x06, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x0a, 0x5a, 0x45, 0x52, 0x4f, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x4e, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x57, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x55, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x53, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x53, 0x10, 0x05, 0x32, 0xc2, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_comment_comment_proto_rawDescOnce sync.Once
	file_services_comment_comment_proto_rawDescData = file_services_comment_comment_proto_rawDesc
)

func file_services_comment_comment_proto_rawDescGZIP() []byte {
	file_services_comment_comment_proto_rawDescOnce.Do(func() {
		file_services_comment_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_comment_comment_proto_rawDescData)
	})
	return file_services_comment_comment_proto_rawDescData
}

var file_services_comment_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_comment_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_comment_comment_proto_goTypes = []interface{}{
	(Rating)(0),                     // 0: comment.Rating
	(*CreateRequest)(nil),           // 1: comment.CreateRequest
	(*CreateFromSourceRequest)(nil), // 2: comment.CreateFromSourceRequest
	(*ListRequest)(nil),             // 3: comment.ListRequest
	(*ListResponse)(nil),            // 4: comment.ListResponse
	(*CommentMsg)(nil),              // 5: comment.CommentMsg
	(*empty.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_services_comment_comment_proto_depIdxs = []int32{
	0, // 0: comment.CreateRequest.Rating:type_name -> comment.Rating
	0, // 1: comment.CreateFromSourceRequest.Rating:type_name -> comment.Rating
	5, // 2: comment.ListResponse.Comments:type_name -> comment.CommentMsg
	0, // 3: comment.CommentMsg.Rating:type_name -> comment.Rating
	1, // 4: comment.Comment.Create:input_type -> comment.CreateRequest
	1, // 5: comment.Comment.CreateFromSource:input_type -> comment.CreateRequest
	3, // 6: comment.Comment.List:input_type -> comment.ListRequest
	6, // 7: comment.Comment.Create:output_type -> google.protobuf.Empty
	6, // 8: comment.Comment.CreateFromSource:output_type -> google.protobuf.Empty
	4, // 9: comment.Comment.List:output_type -> comment.ListResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_comment_comment_proto_init() }
func file_services_comment_comment_proto_init() {
	if File_services_comment_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_comment_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_comment_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFromSourceRequest); i {
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
		file_services_comment_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_services_comment_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_services_comment_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentMsg); i {
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
			RawDescriptor: file_services_comment_comment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_comment_comment_proto_goTypes,
		DependencyIndexes: file_services_comment_comment_proto_depIdxs,
		EnumInfos:         file_services_comment_comment_proto_enumTypes,
		MessageInfos:      file_services_comment_comment_proto_msgTypes,
	}.Build()
	File_services_comment_comment_proto = out.File
	file_services_comment_comment_proto_rawDesc = nil
	file_services_comment_comment_proto_goTypes = nil
	file_services_comment_comment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateFromSource(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/comment.Comment/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) CreateFromSource(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/comment.Comment/CreateFromSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/comment.Comment/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
type CommentServer interface {
	Create(context.Context, *CreateRequest) (*empty.Empty, error)
	CreateFromSource(context.Context, *CreateRequest) (*empty.Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedCommentServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (*UnimplementedCommentServer) Create(context.Context, *CreateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCommentServer) CreateFromSource(context.Context, *CreateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFromSource not implemented")
}
func (*UnimplementedCommentServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterCommentServer(s *grpc.Server, srv CommentServer) {
	s.RegisterService(&_Comment_serviceDesc, srv)
}

func _Comment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.Comment/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_CreateFromSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CreateFromSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.Comment/CreateFromSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CreateFromSource(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.Comment/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Comment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Comment_Create_Handler,
		},
		{
			MethodName: "CreateFromSource",
			Handler:    _Comment_CreateFromSource_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Comment_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/comment/comment.proto",
}
