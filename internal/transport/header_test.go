package transport

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeaders(t *testing.T) {
	headers := NewHeaders()

	assert.NotNil(t, headers)
	assert.Empty(t, headers.Headers)
}

func TestNewFromStringHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  []string
		expected map[string]string
		isErr    bool
	}{
		{
			name:     "valid headers",
			headers:  []string{"Key1:Value1", "Key2:Value2"},
			expected: map[string]string{"Key1": "Value1", "Key2": "Value2"},
			isErr:    false,
		},
		{
			name:     "invalid format",
			headers:  []string{"InvalidHeader"},
			expected: nil,
			isErr:    true,
		},
		{
			name:     "empty key",
			headers:  []string{":Value"},
			expected: nil,
			isErr:    true,
		},
		{
			name:     "empty key value",
			headers:  []string{":"},
			expected: nil,
			isErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers, err := NewFromStringHeaders(tt.headers)
			if tt.isErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, headers.Headers)
		})
	}
}

func TestHeaders_AddSetGet(t *testing.T) {
	tests := []struct {
		name        string
		operation   string
		key         string
		value       string
		isErr       bool
		expectValue string
	}{
		{
			name:        "add valid header",
			operation:   "add",
			key:         "k1",
			value:       "v1",
			isErr:       false,
			expectValue: "v1",
		},
		{
			name:        "add empty key",
			operation:   "add",
			key:         "",
			value:       "v",
			isErr:       true,
			expectValue: "",
		},
		{
			name:        "set new value",
			operation:   "set",
			key:         "k1",
			value:       "newV1",
			isErr:       false,
			expectValue: "newV1",
		},
		{
			name:        "set empty key",
			operation:   "set",
			key:         "",
			value:       "vInvalid",
			isErr:       true,
			expectValue: "",
		},
		{
			name:        "get non-existent key",
			operation:   "get",
			key:         "NonExistentKey",
			value:       "",
			isErr:       false,
			expectValue: "",
		},
	}

	headers := NewHeaders()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch tt.operation {
			case "add":
				err = headers.Add(tt.key, tt.value)
			case "set":
				err = headers.Set(tt.key, tt.value)
			}

			if tt.isErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectValue, headers.Get(tt.key))
		})
	}
}

func TestHeaders_Apply(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
	}{
		{
			name:    "empty headers",
			headers: map[string]string{},
		},
		{
			name: "single header",
			headers: map[string]string{
				"k1": "v1",
			},
		},
		{
			name: "multiple headers",
			headers: map[string]string{
				"k1": "v1",
				"k2": "v2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers := NewHeaders()
			request := &http.Request{Header: make(http.Header)}

			for key, value := range test.headers {
				headers.Add(key, value)
			}

			headers.Apply(request)

			for key, expectedValue := range test.headers {
				val := request.Header.Get(key)
				assert.Equal(t, expectedValue, val)
			}
		})
	}
}

func TestHeaders_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*RequestHeaders)
		expected bool
	}{
		{
			name:     "empty headers",
			setup:    func(h *RequestHeaders) {},
			expected: true,
		},
		{
			name: "non-empty headers",
			setup: func(h *RequestHeaders) {
				h.Add("k", "v")
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := NewHeaders()
			tt.setup(headers)
			assert.Equal(t, tt.expected, headers.IsEmpty())
		})
	}
}

func TestHeaders_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected map[string]string
	}{
		{
			name:     "empty headers",
			headers:  map[string]string{},
			expected: map[string]string{},
		},
		{
			name: "multiple headers",
			headers: map[string]string{
				"k1": "v1",
				"k2": "v2",
			},
			expected: map[string]string{
				"k1": "v1",
				"k2": "v2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := NewHeaders()
			for k, v := range tt.headers {
				headers.Add(k, v)
			}

			result := headers.GetAll()
			assert.Equal(t, tt.expected, result)
		})
	}
}
