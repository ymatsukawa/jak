package transport

import (
	"fmt"
	"slices"
	"strings"
)

const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

const (
	StatusOK           = 200 // OK indicates successful request
	StatusCreated      = 201 // Created indicates successful resource creation
	StatusNoContent    = 204 // No Content indicates success with no response body
	StatusBadRequest   = 400 // Bad Request indicates invalid request
	StatusUnauthorized = 401 // Unauthorized indicates authentication required
	StatusForbidden    = 403 // Forbidden indicates insufficient permissions
	StatusNotFound     = 404 // Not Found indicates resource not found
	StatusServerError  = 500 // Server Error indicates server-side error
)

var validMethods = []string{
	MethodGet,
	MethodPost,
	MethodPut,
	MethodPatch,
	MethodDelete,
	MethodHead,
	MethodOptions,
}

var bodyRequiredMethods = []string{
	MethodPost,
	MethodPut,
	MethodPatch,
}

func IsValidMethod(method string) bool {
	for _, validMethod := range validMethods {
		if method == validMethod {
			return true
		}
	}

	return false
}

func IsBodyRequired(method string) bool {
	for _, bodyMethod := range bodyRequiredMethods {
		if method == bodyMethod {
			return true
		}
	}

	return false
}

func NormalizeMethod(method string) (string, error) {
	method = strings.ToUpper(method)
	if !slices.Contains(validMethods, method) {
		return "", fmt.Errorf("invalid method: %s", method)
	}

	return method, nil
}
