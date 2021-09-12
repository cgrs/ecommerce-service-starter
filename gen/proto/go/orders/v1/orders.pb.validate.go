// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: orders/v1/orders.proto

package ordersv1

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
)

// define the regex for a UUID once up-front
var _orders_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Order) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return OrderValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if len(m.GetLines()) < 1 {
		return OrderValidationError{
			field:  "Lines",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetLines() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderValidationError{
					field:  fmt.Sprintf("Lines[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if _, ok := Status_name[int32(m.GetStatus())]; !ok {
		return OrderValidationError{
			field:  "Status",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

func (m *Order) _validateUuid(uuid string) error {
	if matched := _orders_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OrderValidationError is the validation error returned by Order.Validate if
// the designated constraints aren't met.
type OrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderValidationError) ErrorName() string { return "OrderValidationError" }

// Error satisfies the builtin error interface
func (e OrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderValidationError{}

// Validate checks the field values on Line with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Line) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetItemId()); err != nil {
		return LineValidationError{
			field:  "ItemId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if m.GetQuantity() <= 0 {
		return LineValidationError{
			field:  "Quantity",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

func (m *Line) _validateUuid(uuid string) error {
	if matched := _orders_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// LineValidationError is the validation error returned by Line.Validate if the
// designated constraints aren't met.
type LineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LineValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LineValidationError) ErrorName() string { return "LineValidationError" }

// Error satisfies the builtin error interface
func (e LineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LineValidationError{}

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOrder() == nil {
		return CreateOrderRequestValidationError{
			field:  "Order",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrderRequestValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on ListOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetIds() {
		_, _ = idx, item

		if err := m._validateUuid(item); err != nil {
			return ListOrderRequestValidationError{
				field:  fmt.Sprintf("Ids[%v]", idx),
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *ListOrderRequest) _validateUuid(uuid string) error {
	if matched := _orders_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ListOrderRequestValidationError is the validation error returned by
// ListOrderRequest.Validate if the designated constraints aren't met.
type ListOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderRequestValidationError) ErrorName() string { return "ListOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetOrders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Orders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on FindOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FindOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return FindOrderRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *FindOrderRequest) _validateUuid(uuid string) error {
	if matched := _orders_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// FindOrderRequestValidationError is the validation error returned by
// FindOrderRequest.Validate if the designated constraints aren't met.
type FindOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindOrderRequestValidationError) ErrorName() string { return "FindOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e FindOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindOrderRequestValidationError{}

// Validate checks the field values on UpdateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateOrderRequestValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateOrderRequestValidationError is the validation error returned by
// UpdateOrderRequest.Validate if the designated constraints aren't met.
type UpdateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateOrderRequestValidationError) ErrorName() string {
	return "UpdateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateOrderRequestValidationError{}

// Validate checks the field values on DeleteOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteOrderRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return DeleteOrderRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *DeleteOrderRequest) _validateUuid(uuid string) error {
	if matched := _orders_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteOrderRequestValidationError is the validation error returned by
// DeleteOrderRequest.Validate if the designated constraints aren't met.
type DeleteOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrderRequestValidationError) ErrorName() string {
	return "DeleteOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrderRequestValidationError{}
