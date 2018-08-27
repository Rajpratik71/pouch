// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CPUUsage CPUUsage stores All CPU stats aggregated since container inception.
// swagger:model CPUUsage

type CPUUsage struct {

	// Total CPU time consumed per core (Linux).
	PercpuUsage []uint64 `json:"percpu_usage"`

	// Total CPU time consumed.
	TotalUsage uint64 `json:"total_usage,omitempty"`

	// Time spent by tasks of the cgroup in kernel mode (Linux).
	// Units, nanoseconds (Linux)
	//
	UsageInKernelmode uint64 `json:"usage_in_kernelmode,omitempty"`

	// Time spent by tasks of the cgroup in user mode (Linux).
	// Units, nanoseconds (Linux)
	//
	UsageInUsermode uint64 `json:"usage_in_usermode,omitempty"`
}

/* polymorph CPUUsage percpu_usage false */

/* polymorph CPUUsage total_usage false */

/* polymorph CPUUsage usage_in_kernelmode false */

/* polymorph CPUUsage usage_in_usermode false */

// Validate validates this CPU usage
func (m *CPUUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePercpuUsage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUUsage) validatePercpuUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.PercpuUsage) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPUUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUUsage) UnmarshalBinary(b []byte) error {
	var res CPUUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}