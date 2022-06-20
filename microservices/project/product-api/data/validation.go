package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError is a wrapper around validator.FieldError
type ValidationError struct {
	validator.FieldError
}

// Error returns the error message
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a list of ValidationError
type ValidationErrors []ValidationError

// Errors converts the error slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation is a wrapper around validator.Validate
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// Validate validates the given struct
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i)
	if errs == nil {
		return nil
	}

	ves := errs.(validator.ValidationErrors)
	if len(ves) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range ves {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// validateSKU validates the SKU
func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in the format abc-abc-abc
	re := regexp.MustCompile(`[a-z]+-[0-9]+-[a-z]+`)
	sku := re.FindAllString(fl.Field().String(), -1)

	return len(sku) == 1
}
