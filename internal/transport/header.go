package transport

import (
	"net/http"
	"strings"

	se "github.com/ymatsukawa/jak/internal/sys_error"
)

const (
	ContentTypePlainText      = "text/plain"
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
	ContentTypeJSON           = "application/json"
)

type Headers interface {
	Set(key, val string) error
	Get(key string) string
	Add(key, val string) error
	Apply(req *http.Request)
	IsEmpty() bool
	GetAll() map[string]string
}

type RequestHeaders struct {
	Headers map[string]string
}

func NewHeaders() *RequestHeaders {
	return &RequestHeaders{
		Headers: make(map[string]string),
	}
}

func NewFromStringHeaders(headers []string) (*RequestHeaders, error) {
	resHeader := NewHeaders()

	for _, header := range headers {
		kv := strings.SplitN(header, ":", 2)
		if len(kv) != 2 {
			return nil, se.ErrHeaderInvalidFormat
		}
		err := resHeader.Add(kv[0], kv[1])
		if err != nil {
			return nil, err
		}
	}

	return resHeader, nil
}

func (h *RequestHeaders) Add(key, val string) error {
	k := strings.TrimSpace(key)
	v := strings.TrimSpace(val)
	if k == "" {
		return se.ErrHeaderKeyEmpty
	}
	h.Headers[k] = strings.TrimSpace(v)

	return nil
}

func (h *RequestHeaders) Set(key, val string) error {
	k := strings.TrimSpace(key)
	v := strings.TrimSpace(val)
	if k == "" {
		return se.ErrHeaderKeyEmpty
	}

	h.Headers[k] = strings.TrimSpace(v)
	return nil
}

func (h *RequestHeaders) Get(key string) string {
	return h.Headers[key]
}

func (h *RequestHeaders) Apply(req *http.Request) {
	for k, v := range h.Headers {
		req.Header.Set(k, v)
	}
}

func (h *RequestHeaders) IsEmpty() bool {
	return len(h.Headers) == 0
}

func (h *RequestHeaders) GetAll() map[string]string {
	if len(h.Headers) == 0 {
		return make(map[string]string)
	}

	res := make(map[string]string, len(h.Headers))
	for k, v := range h.Headers {
		res[k] = v
	}

	return res
}
