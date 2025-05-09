package transport

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("TestResponse", func(t *testing.T) {
		res := &Response{
			Header:     make(http.Header),
			StatusCode: 200,
		}
		res.Header.Set("Content-Type", "application/json")

		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	})
}
