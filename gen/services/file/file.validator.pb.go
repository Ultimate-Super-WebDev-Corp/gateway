// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/file/file.proto

package file

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FileUUID) Validate() error {
	if this.UUID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UUID", fmt.Errorf(`value '%v' must not be an empty string`, this.UUID))
	}
	return nil
}
func (this *FileURLs) Validate() error {
	return nil
}
func (this *Chunk) Validate() error {
	if oneOfNester, ok := this.GetOneOfChunk().(*Chunk_Meta); ok {
		if oneOfNester.Meta != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Meta); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
			}
		}
	}
	return nil
}
func (this *FileUploadResponse) Validate() error {
	return nil
}
func (this *FileMetadata) Validate() error {
	return nil
}
func (this *UpdateFileMetadata) Validate() error {
	if this.UUID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UUID", fmt.Errorf(`value '%v' must not be an empty string`, this.UUID))
	}
	if this.Meta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Meta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
		}
	}
	return nil
}
