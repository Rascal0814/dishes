// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/side_dishes.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateSideDishesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSideDishesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSideDishesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSideDishesRequestMultiError, or nil if none found.
func (m *CreateSideDishesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSideDishesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Remark

	if len(errors) > 0 {
		return CreateSideDishesRequestMultiError(errors)
	}

	return nil
}

// CreateSideDishesRequestMultiError is an error wrapping multiple validation
// errors returned by CreateSideDishesRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateSideDishesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSideDishesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSideDishesRequestMultiError) AllErrors() []error { return m }

// CreateSideDishesRequestValidationError is the validation error returned by
// CreateSideDishesRequest.Validate if the designated constraints aren't met.
type CreateSideDishesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSideDishesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSideDishesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSideDishesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSideDishesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSideDishesRequestValidationError) ErrorName() string {
	return "CreateSideDishesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSideDishesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSideDishesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSideDishesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSideDishesRequestValidationError{}

// Validate checks the field values on CreateSideDishesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSideDishesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSideDishesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSideDishesResponseMultiError, or nil if none found.
func (m *CreateSideDishesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSideDishesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSideDishes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSideDishesResponseValidationError{
					field:  "SideDishes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSideDishesResponseValidationError{
					field:  "SideDishes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSideDishes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSideDishesResponseValidationError{
				field:  "SideDishes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateSideDishesResponseMultiError(errors)
	}

	return nil
}

// CreateSideDishesResponseMultiError is an error wrapping multiple validation
// errors returned by CreateSideDishesResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateSideDishesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSideDishesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSideDishesResponseMultiError) AllErrors() []error { return m }

// CreateSideDishesResponseValidationError is the validation error returned by
// CreateSideDishesResponse.Validate if the designated constraints aren't met.
type CreateSideDishesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSideDishesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSideDishesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSideDishesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSideDishesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSideDishesResponseValidationError) ErrorName() string {
	return "CreateSideDishesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSideDishesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSideDishesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSideDishesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSideDishesResponseValidationError{}

// Validate checks the field values on UpdateSideDishesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSideDishesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSideDishesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSideDishesRequestMultiError, or nil if none found.
func (m *UpdateSideDishesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSideDishesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Remark

	if len(errors) > 0 {
		return UpdateSideDishesRequestMultiError(errors)
	}

	return nil
}

// UpdateSideDishesRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateSideDishesRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateSideDishesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSideDishesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSideDishesRequestMultiError) AllErrors() []error { return m }

// UpdateSideDishesRequestValidationError is the validation error returned by
// UpdateSideDishesRequest.Validate if the designated constraints aren't met.
type UpdateSideDishesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSideDishesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSideDishesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSideDishesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSideDishesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSideDishesRequestValidationError) ErrorName() string {
	return "UpdateSideDishesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSideDishesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSideDishesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSideDishesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSideDishesRequestValidationError{}

// Validate checks the field values on UpdateSideDishesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSideDishesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSideDishesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSideDishesResponseMultiError, or nil if none found.
func (m *UpdateSideDishesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSideDishesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSideDishes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSideDishesResponseValidationError{
					field:  "SideDishes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSideDishesResponseValidationError{
					field:  "SideDishes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSideDishes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSideDishesResponseValidationError{
				field:  "SideDishes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSideDishesResponseMultiError(errors)
	}

	return nil
}

// UpdateSideDishesResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateSideDishesResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateSideDishesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSideDishesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSideDishesResponseMultiError) AllErrors() []error { return m }

// UpdateSideDishesResponseValidationError is the validation error returned by
// UpdateSideDishesResponse.Validate if the designated constraints aren't met.
type UpdateSideDishesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSideDishesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSideDishesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSideDishesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSideDishesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSideDishesResponseValidationError) ErrorName() string {
	return "UpdateSideDishesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSideDishesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSideDishesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSideDishesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSideDishesResponseValidationError{}

// Validate checks the field values on GetSideDishesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSideDishesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSideDishesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSideDishesRequestMultiError, or nil if none found.
func (m *GetSideDishesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSideDishesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetSideDishesRequestMultiError(errors)
	}

	return nil
}

// GetSideDishesRequestMultiError is an error wrapping multiple validation
// errors returned by GetSideDishesRequest.ValidateAll() if the designated
// constraints aren't met.
type GetSideDishesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSideDishesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSideDishesRequestMultiError) AllErrors() []error { return m }

// GetSideDishesRequestValidationError is the validation error returned by
// GetSideDishesRequest.Validate if the designated constraints aren't met.
type GetSideDishesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSideDishesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSideDishesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSideDishesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSideDishesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSideDishesRequestValidationError) ErrorName() string {
	return "GetSideDishesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSideDishesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSideDishesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSideDishesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSideDishesRequestValidationError{}

// Validate checks the field values on GetSideDishesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSideDishesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSideDishesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSideDishesResponseMultiError, or nil if none found.
func (m *GetSideDishesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSideDishesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSideDishes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetSideDishesResponseValidationError{
					field:  "SideDishes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetSideDishesResponseValidationError{
					field:  "SideDishes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSideDishes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSideDishesResponseValidationError{
				field:  "SideDishes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetSideDishesResponseMultiError(errors)
	}

	return nil
}

// GetSideDishesResponseMultiError is an error wrapping multiple validation
// errors returned by GetSideDishesResponse.ValidateAll() if the designated
// constraints aren't met.
type GetSideDishesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSideDishesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSideDishesResponseMultiError) AllErrors() []error { return m }

// GetSideDishesResponseValidationError is the validation error returned by
// GetSideDishesResponse.Validate if the designated constraints aren't met.
type GetSideDishesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSideDishesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSideDishesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSideDishesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSideDishesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSideDishesResponseValidationError) ErrorName() string {
	return "GetSideDishesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSideDishesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSideDishesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSideDishesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSideDishesResponseValidationError{}

// Validate checks the field values on DelSideDishesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DelSideDishesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DelSideDishesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DelSideDishesRequestMultiError, or nil if none found.
func (m *DelSideDishesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DelSideDishesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DelSideDishesRequestMultiError(errors)
	}

	return nil
}

// DelSideDishesRequestMultiError is an error wrapping multiple validation
// errors returned by DelSideDishesRequest.ValidateAll() if the designated
// constraints aren't met.
type DelSideDishesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DelSideDishesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DelSideDishesRequestMultiError) AllErrors() []error { return m }

// DelSideDishesRequestValidationError is the validation error returned by
// DelSideDishesRequest.Validate if the designated constraints aren't met.
type DelSideDishesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelSideDishesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelSideDishesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelSideDishesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelSideDishesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelSideDishesRequestValidationError) ErrorName() string {
	return "DelSideDishesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DelSideDishesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelSideDishesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelSideDishesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelSideDishesRequestValidationError{}

// Validate checks the field values on DelSideDishesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DelSideDishesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DelSideDishesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DelSideDishesResponseMultiError, or nil if none found.
func (m *DelSideDishesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DelSideDishesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DelSideDishesResponseMultiError(errors)
	}

	return nil
}

// DelSideDishesResponseMultiError is an error wrapping multiple validation
// errors returned by DelSideDishesResponse.ValidateAll() if the designated
// constraints aren't met.
type DelSideDishesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DelSideDishesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DelSideDishesResponseMultiError) AllErrors() []error { return m }

// DelSideDishesResponseValidationError is the validation error returned by
// DelSideDishesResponse.Validate if the designated constraints aren't met.
type DelSideDishesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelSideDishesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelSideDishesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelSideDishesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelSideDishesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelSideDishesResponseValidationError) ErrorName() string {
	return "DelSideDishesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DelSideDishesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelSideDishesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelSideDishesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelSideDishesResponseValidationError{}