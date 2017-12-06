// Code generated by go-swagger; DO NOT EDIT.

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VCNstruct v c nstruct
// swagger:model VCNstruct
type VCNstruct struct {

	// e tag
	ETag string `json:"ETag,omitempty"`

	// cidr block
	CidrBlock string `json:"cidrBlock,omitempty"`

	// compartment Id
	CompartmentID string `json:"compartmentId,omitempty"`

	// default dhcp options Id
	DefaultDhcpOptionsID string `json:"defaultDhcpOptionsId,omitempty"`

	// default route table Id
	DefaultRouteTableID string `json:"defaultRouteTableId,omitempty"`

	// default security list Id
	DefaultSecurityListID string `json:"defaultSecurityListId,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// dns label
	DNSLabel string `json:"dnsLabel,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// lifecycle state
	LifecycleState string `json:"lifecycleState,omitempty"`

	// time created
	TimeCreated string `json:"timeCreated,omitempty"`

	// vcn domain name
	VcnDomainName string `json:"vcnDomainName,omitempty"`
}

// Validate validates this v c nstruct
func (m *VCNstruct) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VCNstruct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VCNstruct) UnmarshalBinary(b []byte) error {
	var res VCNstruct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}