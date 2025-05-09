package transport

import (
	"bytes"
	"io"
	libhttp "net/http"
	"strconv"
	"time"

	se "github.com/ymatsukawa/jak/internal/sys_error"
)

const (
	defaultTimeout = 30 * time.Second
)

type Client interface {
	Do(req *Request) (*Response, error)
	SetTimeout(second int)
}

type NetClient struct {
	core *libhttp.Client
}

func NewNetClient() *NetClient {
	return &NetClient{
		core: &libhttp.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (client *NetClient) SetTimeout(second int) *NetClient {
	client.core.Timeout = time.Duration(second) * time.Second

	return client
}

func (client *NetClient) Do(userReq *Request) (*Response, error) {
	if err := client.validateRequest(userReq); err != nil {
		return nil, err
	}

	libHttpReq, err := client.createHttpRequest(userReq)
	if err != nil {
		return nil, err
	}

	return client.doRequest(libHttpReq)
}

func (client *NetClient) validateRequest(userReq *Request) error {
	if (*userReq) == nil {
		return se.ErrRequestIsNil
	}

	return nil
}

func (client *NetClient) createHttpRequest(userReq *Request) (*libhttp.Request, error) {
	method, err := NormalizeMethod((*userReq).GetMethod())
	if err != nil {
		return nil, err
	}

	body, err := (*userReq).GetBody()
	if err != nil {
		return nil, err
	}

	byteBody := bytes.NewReader([]byte(body))
	libHttpReq, err := libhttp.NewRequest(method, (*userReq).GetURL(), byteBody)
	if err != nil {
		return nil, se.ErrRequestCreation
	}

	if ctx := (*userReq).GetContext(); ctx != nil {
		libHttpReq = libHttpReq.WithContext(*ctx)
	}

	if err := client.setHeaders(libHttpReq, userReq); err != nil {
		return nil, err
	}

	return libHttpReq, nil
}

func (client *NetClient) setHeaders(libHttpReq *libhttp.Request, userReq *Request) error {
	if headers := (*userReq).GetHeaders(); headers != nil && !headers.IsEmpty() {
		headers.Apply(libHttpReq)
	}

	contentType := (*userReq).GetContentType()

	contentLength := (*userReq).GetContentLength()
	libHttpReq.Header.Set("Content-Type", contentType)
	libHttpReq.Header.Set("Content-Length", strconv.Itoa(contentLength))

	return nil
}

func (client *NetClient) doRequest(libHttpReq *libhttp.Request) (*Response, error) {
	res, err := client.core.Do(libHttpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, se.ErrResponseReadFailed
	}

	return &Response{
		Header:     res.Header,
		Body:       io.NopCloser(bytes.NewReader(bodyBytes)),
		StatusCode: res.StatusCode,
	}, nil
}
