package transport

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawBody(t *testing.T) {
	t.Run("valid raw body", func(t *testing.T) {
		body := NewRawBody("test content")

		assert.Equal(t, ContentTypePlainText, body.GetContentType())
		assert.Equal(t, "test content", body.Content)
		assert.NoError(t, body.Validate())
		assert.False(t, body.IsEmpty())
	})

	t.Run("empty raw body", func(t *testing.T) {
		body := NewRawBody("")

		assert.Error(t, body.Validate())
		assert.True(t, body.IsEmpty())
	})
}

func TestFormBodyX(t *testing.T) {
	t.Run("valid Empty form body", func(t *testing.T) {
		body := NewFormBody("")
		assert.Equal(t, ContentTypeFormURLEncoded, body.GetContentType())
		assert.Equal(t, "", body.GetContent())
		assert.Error(t, body.Validate())
		assert.True(t, body.IsEmpty())
	})

	t.Run("valid form body", func(t *testing.T) {
		var validCases = []string{
			"name=John&age=30",
			"name=John+Doe&age=30",
			"email=john%40example.com",
			"name=&age=30",
			"=value&age=30",
			"color=red&color=blue",
			"message=%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%81%AF",
			"address=123+Main+St%2C+Apt+456",
			"query=a%26b%3Dc",
			"user[name]=John&user[age]=30",
			"data[user][name]=John",
			"key1&key2&key3",
			"symbols=%21%40%23%24%25%5E%26%2A%28%29",
			"single=value",
			"key=",
			"name&age", // same as "name=&age="
		}
		for _, tc := range validCases {
			body := NewFormBody(tc)
			assert.Equal(t, ContentTypeFormURLEncoded, body.GetContentType())
			assert.Equal(t, tc, body.GetContent())
			assert.NoError(t, body.Validate())
			assert.False(t, body.IsEmpty())
		}
	})

	t.Run("invalid form body", func(t *testing.T) {
		var testCases = []string{
			"name=John%",
			"name=John%2",
			"name=John%ZZ",
			"name=John%GH",
			"value=50%%",
			"name=John%2",
		}

		for _, tc := range testCases {
			fmt.Println(tc)
			body := NewFormBody(tc)
			assert.Error(t, body.Validate(), tc)
		}
	})

	t.Run("strictly, invalid but accepted", func(t *testing.T) {
		var testCases = []string{
			"email=john@example.com",
			"name=John Doe",
			"nameJohn&age=30",
			"message=hello\\world",
			"name=John&&age=30",
			"message=こんにちは世界",
		}

		for _, tc := range testCases {
			body := NewFormBody(tc)
			assert.Equal(t, ContentTypeFormURLEncoded, body.GetContentType())
			assert.Equal(t, tc, body.GetContent())
			assert.NoError(t, body.Validate())
			assert.False(t, body.IsEmpty())
		}
	})
}

func TestJsonBody(t *testing.T) {
	t.Run("valid json body", func(t *testing.T) {
		body := NewJSONBody(`{"key": "value"}`)

		assert.Equal(t, ContentTypeJSON, body.GetContentType())
		assert.NoError(t, body.Validate())
	})

	t.Run("invalid json body", func(t *testing.T) {
		testCases := []string{
			"",
			"{invalid json}",
			`{"unclosed": "object"`,
		}

		for _, tc := range testCases {
			body := NewJSONBody(tc)
			assert.Error(t, body.Validate())
		}
	})
}
