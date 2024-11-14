// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: server.proto

package proto

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

// Validate checks the field values on PageInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageInfoMultiError, or nil
// if none found.
func (m *PageInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PageInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pn

	// no validation rules for Psize

	if len(errors) > 0 {
		return PageInfoMultiError(errors)
	}

	return nil
}

// PageInfoMultiError is an error wrapping multiple validation errors returned
// by PageInfo.ValidateAll() if the designated constraints aren't met.
type PageInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageInfoMultiError) AllErrors() []error { return m }

// PageInfoValidationError is the validation error returned by
// PageInfo.Validate if the designated constraints aren't met.
type PageInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageInfoValidationError) ErrorName() string { return "PageInfoValidationError" }

// Error satisfies the builtin error interface
func (e PageInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageInfoValidationError{}

// Validate checks the field values on UserListRespons with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserListRespons) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserListRespons with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserListResponsMultiError, or nil if none found.
func (m *UserListRespons) ValidateAll() error {
	return m.validate(true)
}

func (m *UserListRespons) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserListResponsValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserListResponsValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListResponsValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserListResponsMultiError(errors)
	}

	return nil
}

// UserListResponsMultiError is an error wrapping multiple validation errors
// returned by UserListRespons.ValidateAll() if the designated constraints
// aren't met.
type UserListResponsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListResponsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListResponsMultiError) AllErrors() []error { return m }

// UserListResponsValidationError is the validation error returned by
// UserListRespons.Validate if the designated constraints aren't met.
type UserListResponsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListResponsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListResponsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListResponsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListResponsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListResponsValidationError) ErrorName() string { return "UserListResponsValidationError" }

// Error satisfies the builtin error interface
func (e UserListResponsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListRespons.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListResponsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListResponsValidationError{}

// Validate checks the field values on MobileInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MobileInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MobileInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MobileInfoMultiError, or
// nil if none found.
func (m *MobileInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *MobileInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Mobile

	if len(errors) > 0 {
		return MobileInfoMultiError(errors)
	}

	return nil
}

// MobileInfoMultiError is an error wrapping multiple validation errors
// returned by MobileInfo.ValidateAll() if the designated constraints aren't met.
type MobileInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MobileInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MobileInfoMultiError) AllErrors() []error { return m }

// MobileInfoValidationError is the validation error returned by
// MobileInfo.Validate if the designated constraints aren't met.
type MobileInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MobileInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MobileInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MobileInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MobileInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MobileInfoValidationError) ErrorName() string { return "MobileInfoValidationError" }

// Error satisfies the builtin error interface
func (e MobileInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMobileInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MobileInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MobileInfoValidationError{}

// Validate checks the field values on UserRespons with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserRespons) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRespons with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserResponsMultiError, or
// nil if none found.
func (m *UserRespons) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRespons) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Moblie

	// no validation rules for Name

	// no validation rules for Birthday

	// no validation rules for Gender

	// no validation rules for Role

	if len(errors) > 0 {
		return UserResponsMultiError(errors)
	}

	return nil
}

// UserResponsMultiError is an error wrapping multiple validation errors
// returned by UserRespons.ValidateAll() if the designated constraints aren't met.
type UserResponsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserResponsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserResponsMultiError) AllErrors() []error { return m }

// UserResponsValidationError is the validation error returned by
// UserRespons.Validate if the designated constraints aren't met.
type UserResponsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserResponsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserResponsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserResponsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserResponsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserResponsValidationError) ErrorName() string { return "UserResponsValidationError" }

// Error satisfies the builtin error interface
func (e UserResponsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRespons.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserResponsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserResponsValidationError{}

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserInfoMultiError, or nil
// if none found.
func (m *UserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Password

	// no validation rules for Moblie

	// no validation rules for Name

	// no validation rules for Birthday

	// no validation rules for Gender

	// no validation rules for Id

	// no validation rules for Salt

	if len(errors) > 0 {
		return UserInfoMultiError(errors)
	}

	return nil
}

// UserInfoMultiError is an error wrapping multiple validation errors returned
// by UserInfo.ValidateAll() if the designated constraints aren't met.
type UserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserInfoMultiError) AllErrors() []error { return m }

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}

// Validate checks the field values on SuccessResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SuccessResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SuccessResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SuccessResponseMultiError, or nil if none found.
func (m *SuccessResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SuccessResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return SuccessResponseMultiError(errors)
	}

	return nil
}

// SuccessResponseMultiError is an error wrapping multiple validation errors
// returned by SuccessResponse.ValidateAll() if the designated constraints
// aren't met.
type SuccessResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SuccessResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SuccessResponseMultiError) AllErrors() []error { return m }

// SuccessResponseValidationError is the validation error returned by
// SuccessResponse.Validate if the designated constraints aren't met.
type SuccessResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SuccessResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SuccessResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SuccessResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SuccessResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SuccessResponseValidationError) ErrorName() string { return "SuccessResponseValidationError" }

// Error satisfies the builtin error interface
func (e SuccessResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSuccessResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SuccessResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SuccessResponseValidationError{}

// Validate checks the field values on PasswordInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PasswordInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PasswordInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PasswordInfoMultiError, or
// nil if none found.
func (m *PasswordInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *PasswordInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Password

	// no validation rules for Id

	if len(errors) > 0 {
		return PasswordInfoMultiError(errors)
	}

	return nil
}

// PasswordInfoMultiError is an error wrapping multiple validation errors
// returned by PasswordInfo.ValidateAll() if the designated constraints aren't met.
type PasswordInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PasswordInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PasswordInfoMultiError) AllErrors() []error { return m }

// PasswordInfoValidationError is the validation error returned by
// PasswordInfo.Validate if the designated constraints aren't met.
type PasswordInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PasswordInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PasswordInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PasswordInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PasswordInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PasswordInfoValidationError) ErrorName() string { return "PasswordInfoValidationError" }

// Error satisfies the builtin error interface
func (e PasswordInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPasswordInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PasswordInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PasswordInfoValidationError{}

// Validate checks the field values on IdInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdInfoMultiError, or nil if none found.
func (m *IdInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *IdInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdInfoMultiError(errors)
	}

	return nil
}

// IdInfoMultiError is an error wrapping multiple validation errors returned by
// IdInfo.ValidateAll() if the designated constraints aren't met.
type IdInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdInfoMultiError) AllErrors() []error { return m }

// IdInfoValidationError is the validation error returned by IdInfo.Validate if
// the designated constraints aren't met.
type IdInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdInfoValidationError) ErrorName() string { return "IdInfoValidationError" }

// Error satisfies the builtin error interface
func (e IdInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdInfoValidationError{}

// Validate checks the field values on IdResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdResponseMultiError, or
// nil if none found.
func (m *IdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Role

	// no validation rules for Success

	if len(errors) > 0 {
		return IdResponseMultiError(errors)
	}

	return nil
}

// IdResponseMultiError is an error wrapping multiple validation errors
// returned by IdResponse.ValidateAll() if the designated constraints aren't met.
type IdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdResponseMultiError) AllErrors() []error { return m }

// IdResponseValidationError is the validation error returned by
// IdResponse.Validate if the designated constraints aren't met.
type IdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdResponseValidationError) ErrorName() string { return "IdResponseValidationError" }

// Error satisfies the builtin error interface
func (e IdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdResponseValidationError{}