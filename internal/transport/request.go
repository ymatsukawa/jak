package transport

import (
	"context"
)

type Request interface {
	GetURL() string
	GetMethod() string
	GetBody() (string, error)
	GetContentType() string
	GetContentLength() int
	GetHeaders() Headers
	GetContext() *context.Context
	WithContext(ctx context.Context) *Request
	WithHeaders(headers Headers) *Request
	WithBody(body RequestBody) *Request
}

type BaseRequest struct {
	URL     string
	Method  string
	Headers *Headers
	Body    *RequestBody
	Context *context.Context
}

func NewRequest(method, url string) *BaseRequest {
	return &BaseRequest{
		URL:     url,
		Method:  method,
		Headers: nil,
		Body:    nil,
		Context: nil,
	}
}

func (req *BaseRequest) GetURL() string {
	return req.URL
}

func (req *BaseRequest) GetMethod() string {
	return req.Method
}

func (req *BaseRequest) GetBody() (string, error) {
	if *req.Body == nil {
		return "", nil
	}

	body := *req.Body
	if err := body.Validate(); err != nil {
		return "", err
	}

	return body.GetContent(), nil
}

func (req *BaseRequest) GetContentType() string {
	if req.Body != nil && !(*req.Body).IsEmpty() {
		return (*req.Body).GetContentType()
	}

	return ""
}

func (req *BaseRequest) GetContentLength() int {
	if req.Body != nil && !(*req.Body).IsEmpty() {
		return len((*req.Body).GetContent())
	}

	return 0
}

func (req *BaseRequest) GetHeaders() Headers {
	return *req.Headers
}

func (req *BaseRequest) GetContext() *context.Context {
	return req.Context
}

func (req *BaseRequest) WithContext(ctx context.Context) *BaseRequest {
	req.Context = &ctx
	return req
}

func (req *BaseRequest) WithHeaders(headers Headers) *BaseRequest {
	req.Headers = &headers
	return req
}

func (req *BaseRequest) WithBody(body RequestBody) *BaseRequest {
	req.Body = &body
	return req
}
