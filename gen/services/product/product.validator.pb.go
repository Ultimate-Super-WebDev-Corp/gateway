// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Ultimate-Super-WebDev-Corp/gateway/services/product/product.proto

package product

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SearchByUUIDsRequest) Validate() error {
	if len(this.UUIDs) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("UUIDs", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.UUIDs))
	}
	return nil
}
func (this *ProductMsg) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Brand == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Brand", fmt.Errorf(`value '%v' must not be an empty string`, this.Brand))
	}
	if this.Description == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Description", fmt.Errorf(`value '%v' must not be an empty string`, this.Description))
	}
	return nil
}
func (this *ProductWithID) Validate() error {
	if this.Product != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Product); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Product", err)
		}
	}
	return nil
}
func (this *GetByIDRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	return nil
}
func (this *CatalogRequest) Validate() error {
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	return nil
}
func (this *CatalogMetaRequest) Validate() error {
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	return nil
}
func (this *Sort) Validate() error {
	return nil
}
func (this *Filter) Validate() error {
	if oneOfNester, ok := this.GetValue().(*Filter_ListFilter); ok {
		if oneOfNester.ListFilter != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ListFilter); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ListFilter", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*Filter_RangeFilter); ok {
		if oneOfNester.RangeFilter != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RangeFilter); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RangeFilter", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*Filter_SwitchFilter); ok {
		if oneOfNester.SwitchFilter != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SwitchFilter); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SwitchFilter", err)
			}
		}
	}
	return nil
}
func (this *ListFilter) Validate() error {
	return nil
}
func (this *SwitchFilter) Validate() error {
	return nil
}
func (this *RangeFilter) Validate() error {
	if this.AvailableValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AvailableValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AvailableValue", err)
		}
	}
	if this.SelectedValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectedValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectedValue", err)
		}
	}
	return nil
}
func (this *RangeValue) Validate() error {
	return nil
}
func (this *CatalogResponse) Validate() error {
	for _, item := range this.Products {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Products", err)
			}
		}
	}
	return nil
}
func (this *CatalogMetaResponse) Validate() error {
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	for _, item := range this.Categories {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Categories", err)
			}
		}
	}
	for _, item := range this.Sorts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sorts", err)
			}
		}
	}
	return nil
}
func (this *CatalogProduct) Validate() error {
	return nil
}
func (this *Category) Validate() error {
	for _, item := range this.Categories {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Categories", err)
			}
		}
	}
	return nil
}
