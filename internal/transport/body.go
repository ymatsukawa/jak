package transport

import (
	"encoding/json"
	"net/url"

	se "github.com/ymatsukawa/jak/internal/sys_error"
)

type RequestBody interface {
	GetContent() string
	GetContentType() string
	Validate() error
	IsEmpty() bool
}

type RawBody struct {
	Content string
	Type    string
}

type FormBody struct {
	Content string
	Type    string
}

type JsonBody struct {
	Content string
	Type    string
}

func NewRawBody(content string) *RawBody {
	return &RawBody{
		Content: content,
		Type:    ContentTypePlainText,
	}
}

func (b *RawBody) GetContent() string {
	return b.Content
}

func (b *RawBody) GetContentType() string {
	return b.Type
}

func (b *RawBody) Validate() error {
	err := validateBasicBody(b.Content, b.Type)
	if err != nil {
		return err
	}

	return nil
}

func (b *RawBody) IsEmpty() bool {
	return b.Content == ""
}

func NewFormBody(content string) *FormBody {
	return &FormBody{
		Content: content,
		Type:    ContentTypeFormURLEncoded,
	}
}

func (b *FormBody) GetContent() string {
	return b.Content
}

func (b *FormBody) GetContentType() string {
	return b.Type
}

func (b *FormBody) Validate() error {
	err := validateBasicBody(b.Content, b.Type)
	if err != nil {
		return err
	}
	if !isValidURLEncoded(b.Content) {
		return se.ErrInvalidURLEncodedFormat
	}

	return nil
}

func isValidURLEncoded(s string) bool {
	if s == "" {
		return false
	}

	_, err := url.ParseQuery(s)
	return err == nil
}

func (b *FormBody) IsEmpty() bool {
	return b.Content == ""
}

func NewJSONBody(content string) *JsonBody {
	return &JsonBody{
		Content: content,
		Type:    ContentTypeJSON,
	}
}

func (b *JsonBody) GetContent() string {
	return b.Content
}

func (b *JsonBody) GetContentType() string {
	return b.Type
}

func (b *JsonBody) Validate() error {
	err := validateBasicBody(b.Content, b.Type)
	if err != nil {
		return err
	}
	if !json.Valid([]byte(b.Content)) {
		return se.ErrInvalidJSONFormat
	}

	return nil
}

func (b *JsonBody) IsEmpty() bool {
	return b.Content == ""
}

func validateBasicBody(content, contentType string) error {
	if content == "" {
		return se.ErrContentEmpty
	}
	if contentType == "" {
		return se.ErrContentTypeEmpty
	}

	return nil
}
