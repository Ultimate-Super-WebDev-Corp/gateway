// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/search/search.proto

package search

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SearchByIdsRequest) Validate() error {
	if len(this.UUIDs) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("UUIDs", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.UUIDs))
	}
	return nil
}
func (this *Product) Validate() error {
	return nil
}
