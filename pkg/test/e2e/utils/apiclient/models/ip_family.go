// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPFamily +kubebuilder:validation:Enum="";IPv4;IPv4+IPv6
//
// swagger:model IPFamily
type IPFamily string

// Validate validates this IP family
func (m IPFamily) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP family based on context it is used
func (m IPFamily) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}