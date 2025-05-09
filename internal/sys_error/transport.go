package sys_error

import "errors"

var (
	ErrRequestIsNil            = errors.New("request is nil")
	ErrRequestCreation         = errors.New("failed to create HTTP request")
	ErrInvalidMethod           = errors.New("invalid HTTP method")
	ErrHeaderInvalidFormat     = errors.New("invalid header format")
	ErrHeaderKeyEmpty          = errors.New("header key cannot be empty")
	ErrBodyEmpty               = errors.New("body cannot be empty")
	ErrContentTypeEmpty        = errors.New("content type cannot be empty")
	ErrContentEmpty            = errors.New("content cannot be empty")
	ErrInvalidURLEncodedFormat = errors.New("invalid URL-encoded format")
	ErrInvalidJSONFormat       = errors.New("invalid JSON body")
	ErrResponseReadFailed      = errors.New("failed to read response body")
)
