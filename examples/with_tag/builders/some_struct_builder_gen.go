// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kilianpaquier/go-builder-generator/examples/with_tag"
)

// SomeStructBuilder is an alias of SomeStruct to build SomeStruct with builder-pattern.
type SomeStructBuilder with_tag.SomeStruct

// NewSomeStructBuilder creates a new SomeStructBuilder.
func NewSomeStructBuilder() *SomeStructBuilder {
	return &SomeStructBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *SomeStructBuilder) Copy() *SomeStructBuilder {
	c := *b
	return &c
}

// Build returns built SomeStruct.
func (b *SomeStructBuilder) Build() (*with_tag.SomeStruct, error) {
	b = b.SetTheChan()

	result := (*with_tag.SomeStruct)(b)
	if err := validator.New().Struct(result); err != nil {
		return nil, fmt.Errorf("failed to validate 'SomeStruct' struct: %w", err)
	}
	return result, nil
}

// SetSomeSlice sets SomeStruct's SomeSlice.
func (b *SomeStructBuilder) SetSomeSlice(someSlice ...string) *SomeStructBuilder {
	b.SomeSlice = append(b.SomeSlice, someSlice...)
	return b
}

// SetValidation sets SomeStruct's Validation.
func (b *SomeStructBuilder) SetValidation(validation string) *SomeStructBuilder {
	b.Validation = validation
	return b
}
