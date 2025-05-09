package transport

import (
	"io"
	"net/http"
)

type Response struct {
	Header     http.Header
	Body       io.ReadCloser
	StatusCode int
}
