// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on ClusterLoadAssignment with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterLoadAssignment) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetClusterName()) < 1 {
		return ClusterLoadAssignmentValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 bytes",
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClusterLoadAssignmentValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for NamedEndpoints

	{
		tmp := m.GetPolicy()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterLoadAssignmentValidationError{
					field:  "Policy",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClusterLoadAssignmentValidationError is the validation error returned by
// ClusterLoadAssignment.Validate if the designated constraints aren't met.
type ClusterLoadAssignmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterLoadAssignmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterLoadAssignmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterLoadAssignmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterLoadAssignmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterLoadAssignmentValidationError) ErrorName() string {
	return "ClusterLoadAssignmentValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterLoadAssignmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterLoadAssignment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterLoadAssignmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterLoadAssignmentValidationError{}

// Validate checks the field values on ClusterLoadAssignment_Policy with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterLoadAssignment_Policy) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetDropOverloads() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClusterLoadAssignment_PolicyValidationError{
						field:  fmt.Sprintf("DropOverloads[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	if wrapper := m.GetOverprovisioningFactor(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return ClusterLoadAssignment_PolicyValidationError{
				field:  "OverprovisioningFactor",
				reason: "value must be greater than 0",
			}
		}

	}

	return nil
}

// ClusterLoadAssignment_PolicyValidationError is the validation error returned
// by ClusterLoadAssignment_Policy.Validate if the designated constraints
// aren't met.
type ClusterLoadAssignment_PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterLoadAssignment_PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterLoadAssignment_PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterLoadAssignment_PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterLoadAssignment_PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterLoadAssignment_PolicyValidationError) ErrorName() string {
	return "ClusterLoadAssignment_PolicyValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterLoadAssignment_PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterLoadAssignment_Policy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterLoadAssignment_PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterLoadAssignment_PolicyValidationError{}

// Validate checks the field values on
// ClusterLoadAssignment_Policy_DropOverload with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ClusterLoadAssignment_Policy_DropOverload) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCategory()) < 1 {
		return ClusterLoadAssignment_Policy_DropOverloadValidationError{
			field:  "Category",
			reason: "value length must be at least 1 bytes",
		}
	}

	{
		tmp := m.GetDropPercentage()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterLoadAssignment_Policy_DropOverloadValidationError{
					field:  "DropPercentage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClusterLoadAssignment_Policy_DropOverloadValidationError is the validation
// error returned by ClusterLoadAssignment_Policy_DropOverload.Validate if the
// designated constraints aren't met.
type ClusterLoadAssignment_Policy_DropOverloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterLoadAssignment_Policy_DropOverloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterLoadAssignment_Policy_DropOverloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterLoadAssignment_Policy_DropOverloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterLoadAssignment_Policy_DropOverloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterLoadAssignment_Policy_DropOverloadValidationError) ErrorName() string {
	return "ClusterLoadAssignment_Policy_DropOverloadValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterLoadAssignment_Policy_DropOverloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterLoadAssignment_Policy_DropOverload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterLoadAssignment_Policy_DropOverloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterLoadAssignment_Policy_DropOverloadValidationError{}
