// Code generated by go generate; DO NOT EDIT.
package generate

import (
	"net/url"

	"github.com/containers/podman/v3/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *SystemdOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *SystemdOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithUseName set field UseName to given value
func (o *SystemdOptions) WithUseName(value bool) *SystemdOptions {
	o.UseName = &value
	return o
}

// GetUseName returns value of field UseName
func (o *SystemdOptions) GetUseName() bool {
	if o.UseName == nil {
		var z bool
		return z
	}
	return *o.UseName
}

// WithNew set field New to given value
func (o *SystemdOptions) WithNew(value bool) *SystemdOptions {
	o.New = &value
	return o
}

// GetNew returns value of field New
func (o *SystemdOptions) GetNew() bool {
	if o.New == nil {
		var z bool
		return z
	}
	return *o.New
}

// WithNoHeader set field NoHeader to given value
func (o *SystemdOptions) WithNoHeader(value bool) *SystemdOptions {
	o.NoHeader = &value
	return o
}

// GetNoHeader returns value of field NoHeader
func (o *SystemdOptions) GetNoHeader() bool {
	if o.NoHeader == nil {
		var z bool
		return z
	}
	return *o.NoHeader
}

// WithRestartPolicy set field RestartPolicy to given value
func (o *SystemdOptions) WithRestartPolicy(value string) *SystemdOptions {
	o.RestartPolicy = &value
	return o
}

// GetRestartPolicy returns value of field RestartPolicy
func (o *SystemdOptions) GetRestartPolicy() string {
	if o.RestartPolicy == nil {
		var z string
		return z
	}
	return *o.RestartPolicy
}

// WithRestartSec set field RestartSec to given value
func (o *SystemdOptions) WithRestartSec(value uint) *SystemdOptions {
	o.RestartSec = &value
	return o
}

// GetRestartSec returns value of field RestartSec
func (o *SystemdOptions) GetRestartSec() uint {
	if o.RestartSec == nil {
		var z uint
		return z
	}
	return *o.RestartSec
}

// WithStopTimeout set field StopTimeout to given value
func (o *SystemdOptions) WithStopTimeout(value uint) *SystemdOptions {
	o.StopTimeout = &value
	return o
}

// GetStopTimeout returns value of field StopTimeout
func (o *SystemdOptions) GetStopTimeout() uint {
	if o.StopTimeout == nil {
		var z uint
		return z
	}
	return *o.StopTimeout
}

// WithContainerPrefix set field ContainerPrefix to given value
func (o *SystemdOptions) WithContainerPrefix(value string) *SystemdOptions {
	o.ContainerPrefix = &value
	return o
}

// GetContainerPrefix returns value of field ContainerPrefix
func (o *SystemdOptions) GetContainerPrefix() string {
	if o.ContainerPrefix == nil {
		var z string
		return z
	}
	return *o.ContainerPrefix
}

// WithPodPrefix set field PodPrefix to given value
func (o *SystemdOptions) WithPodPrefix(value string) *SystemdOptions {
	o.PodPrefix = &value
	return o
}

// GetPodPrefix returns value of field PodPrefix
func (o *SystemdOptions) GetPodPrefix() string {
	if o.PodPrefix == nil {
		var z string
		return z
	}
	return *o.PodPrefix
}

// WithSeparator set field Separator to given value
func (o *SystemdOptions) WithSeparator(value string) *SystemdOptions {
	o.Separator = &value
	return o
}

// GetSeparator returns value of field Separator
func (o *SystemdOptions) GetSeparator() string {
	if o.Separator == nil {
		var z string
		return z
	}
	return *o.Separator
}