package net_client

import (
	"bytes"
	"net/http"

	"github.com/ymatsukawa/jak/internal/rule"
	tpt "github.com/ymatsukawa/jak/internal/transport"
)

const (
	defaultMethod  = "GET"
	defaultBaseURL = "http://example.com/look/internal/net_client/builder.go"
)

type Builder interface {
	SetMethod(method string) *Builder
	SetURL(url, path string) *Builder
	SetHeaders(headers []string) *Builder
	SetBody(contentType, body string) *Builder

	BuildSimple() (*http.Request, error)
	BuildFromConfig(config *rule.Config, request *rule.Request) (*http.Request, error)
}

type RequestBuilder struct {
	Method  string
	URL     string
	Headers tpt.Headers
	Body    tpt.RequestBody
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		Method:  defaultMethod,
		URL:     defaultBaseURL,
		Headers: nil,
		Body:    nil,
	}
}

func (b *RequestBuilder) SetMethod(method string) *RequestBuilder {
	b.Method = method
	return b
}

func (b *RequestBuilder) SetURL(url string, path *string) *RequestBuilder {
	b.URL = url
	if path != nil {
		b.URL += *path
	}
	return b
}

func (b *RequestBuilder) SetHeaders(headers []string) *RequestBuilder {
	hs, err := tpt.NewFromStringHeaders(headers)
	if err != nil {
		return nil
	}

	for k, v := range hs.GetAll() {
		if err := b.Headers.Set(k, v); err != nil {
			return nil
		}
	}

	return b
}

func (b *RequestBuilder) SetBody(body string) *RequestBuilder {
	contentType := tpt.SpecifyBodyType(body)
	switch contentType {
	case tpt.ContentTypeJSON:
		b.Body = tpt.NewJSONBody(body)
	case tpt.ContentTypePlainText:
		b.Body = tpt.NewRawBody(body)
	case tpt.ContentTypeFormURLEncoded:
		b.Body = tpt.NewFormBody(body)
	}
	return b
}

func (b *RequestBuilder) BuildSimple() (*http.Request, error) {
	var body *bytes.Reader = nil
	if b.Body != nil {
		sBody := b.Body.GetContent()
		body = bytes.NewReader([]byte(sBody))
	}

	req, err := http.NewRequest(b.Method, b.URL, body)
	if err != nil {
		return nil, err
	}

	if b.Headers != nil {
		b.Headers.Apply(req)
	}

	return req, nil
}
