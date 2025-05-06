package errors

import (
	"errors"
	"fmt"
)

// Common errors shared across packages
var (
	ErrCLIInput           = errors.New("invalid CLI input")
	ErrInvalidURL         = errors.New("invalid URL")
	ErrInvalidConfig      = errors.New("configuration is not loadable")
	ErrConfigValidation   = errors.New("configuration validation failed")
	ErrRequestExecution   = errors.New("request execution failed")
	ErrRequestCreation    = errors.New("request creation failed")
	ErrRequestPreparation = errors.New("request preparation failed")
	ErrBatchRequestFailed = errors.New("batch request failed")
	ErrInvalidMethod      = errors.New("invalid HTTP method")
	ErrInvalidHeader      = errors.New("invalid header format")
	ErrInvalidBody        = errors.New("invalid request body")
	ErrResponseRead       = errors.New("failed to read response")
)

// WrapError wraps an error with additional context
func WrapError(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf(format+": %w", append(args, err)...)
}
