package transport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMethod(t *testing.T) {
	tests := []struct {
		name   string
		method string
		expect bool
	}{
		{"valid GET", MethodGet, true},
		{"valid POST", MethodPost, true},
		{"valid PUT", MethodPut, true},
		{"valid PATCH", MethodPatch, true},
		{"valid DELETE", MethodDelete, true},
		{"valid HEAD", MethodHead, true},
		{"valid OPTIONS", MethodOptions, true},
		{"invalid method", "INVALID", false},
		{"invalid empty method", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidMethod(tt.method)
			assert.Equal(t, tt.expect, got)
		})
	}
}

func TestIsBodyRequired(t *testing.T) {
	tests := []struct {
		name   string
		method string
		expect bool
	}{
		{"POST requires body", MethodPost, true},
		{"PUT requires body", MethodPut, true},
		{"PATCH requires body", MethodPatch, true},
		{"GET no body required", MethodGet, false},
		{"DELETE no body required", MethodDelete, false},
		{"HEAD no body required", MethodHead, false},
		{"OPTIONS no body required", MethodOptions, false},
		{"invalid method no body", "INVALID", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBodyRequired(tt.method)
			assert.Equal(t, tt.expect, got)
		})
	}
}

func TestNormalizeMethod(t *testing.T) {
	tests := []struct {
		name      string
		method    string
		expect    string
		expectErr bool
	}{
		{"lowercase get", "get", "GET", false},
		{"mixed case post", "PoSt", "POST", false},
		{"uppercase put", "PUT", "PUT", false},
		{"invalid method", "INVALID", "", true},
		{"empty method", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeMethod(tt.method)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, got)
		})
	}
}
