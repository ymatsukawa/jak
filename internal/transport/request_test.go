package transport

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	var testURL = "http://example.com"

	t.Run("NewRequest creates request with basic options", func(t *testing.T) {
		req := NewRequest("GET", testURL)

		assert.Equal(t, testURL, req.GetURL())
		assert.Equal(t, "GET", req.GetMethod())
	})

	t.Run("GetBody returns the raw body", func(t *testing.T) {
		input := "test body"
		rawBody := NewRawBody(input)
		req := NewRequest("GET", testURL).WithBody(rawBody)

		body, err := req.GetBody()

		assert.Nil(t, err)
		assert.Equal(t, body, input)
	})

	t.Run("GetBody returns the form body", func(t *testing.T) {
		input := "key=value"
		formBody := NewFormBody(input)
		req := NewRequest("GET", testURL).WithBody(formBody)

		body, err := req.GetBody()

		assert.Nil(t, err)
		assert.Equal(t, body, input)
	})

	t.Run("GetBody returns the json body", func(t *testing.T) {
		input := `{"key": "value"}`
		jsonBody := NewJSONBody(input)
		req := NewRequest("GET", testURL).WithBody(jsonBody)

		body, err := req.GetBody()

		assert.Nil(t, err)
		assert.Equal(t, body, input)
	})

	t.Run("GetContentType returns empty for no body", func(t *testing.T) {
		req := NewRequest("GET", testURL)
		assert.Equal(t, "", req.GetContentType())
	})

	t.Run("GetContentType returns content type for raw body", func(t *testing.T) {
		rawBody := NewRawBody("test body")
		req := NewRequest("GET", testURL).WithBody(rawBody)
		assert.Equal(t, "text/plain", req.GetContentType())
	})

	t.Run("GetContentLength returns 0 for no body", func(t *testing.T) {
		req := NewRequest("GET", testURL)

		assert.Equal(t, 0, req.GetContentLength())
	})

	t.Run("GetContentLength returns length for raw body", func(t *testing.T) {
		input := "test body"
		rawBody := NewRawBody(input)
		req := NewRequest("GET", testURL).WithBody(rawBody)

		assert.Equal(t, len(input), req.GetContentLength())
	})

	t.Run("GetHeaders returns request headers", func(t *testing.T) {
		headers := NewHeaders()
		headers.Add("Content-Type", "application/json")
		headers.Add("Authorization", "Bearer test-token")

		req := NewRequest(testURL, "GET").WithHeaders(headers)
		assert.Equal(t, headers, req.GetHeaders())
	})

	t.Run("WithContext sets context", func(t *testing.T) {
		req := NewRequest(testURL, "GET")

		ctx := context.Background()
		req = req.WithContext(ctx)

		assert.Equal(t, &ctx, req.GetContext())
	})
}
