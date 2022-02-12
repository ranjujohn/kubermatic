// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodegroupScalingConfig nodegroup scaling config
//
// swagger:model NodegroupScalingConfig
type NodegroupScalingConfig struct {

	// MaxCount - The maximum number of nodes for auto-scaling
	MaxCount int32 `json:"maxCount,omitempty"`

	// MinCount - The minimum number of nodes for auto-scaling
	MinCount int32 `json:"minCount,omitempty"`
}

// Validate validates this nodegroup scaling config
func (m *NodegroupScalingConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nodegroup scaling config based on context it is used
func (m *NodegroupScalingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodegroupScalingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodegroupScalingConfig) UnmarshalBinary(b []byte) error {
	var res NodegroupScalingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}