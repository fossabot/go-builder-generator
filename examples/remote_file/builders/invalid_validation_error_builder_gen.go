// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// InvalidValidationErrorBuilder is an alias of InvalidValidationError to build InvalidValidationError with builder-pattern.
type InvalidValidationErrorBuilder validator.InvalidValidationError

// NewInvalidValidationErrorBuilder creates a new InvalidValidationErrorBuilder.
func NewInvalidValidationErrorBuilder() *InvalidValidationErrorBuilder {
	return &InvalidValidationErrorBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *InvalidValidationErrorBuilder) Copy() *InvalidValidationErrorBuilder {
	c := *b
	return &c
}

// Build returns built InvalidValidationError.
func (b *InvalidValidationErrorBuilder) Build() *validator.InvalidValidationError {
	c := (validator.InvalidValidationError)(*b)
	return &c
}

// SetType sets InvalidValidationError's Type.
func (b *InvalidValidationErrorBuilder) SetType(t reflect.Type) *InvalidValidationErrorBuilder {
	b.Type = t
	return b
}