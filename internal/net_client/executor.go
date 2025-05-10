package net_client

import (
	"bytes"
	"io"
	"net/http"

	se "github.com/ymatsukawa/jak/internal/sys_error"
	tpt "github.com/ymatsukawa/jak/internal/transport"
)

type Executor interface {
	Execute()
}

type NetExecutor struct {
	LibHttpRequest http.Request
}

func NewNetExecutor(libHttpRequest http.Request) *NetExecutor {
	return &NetExecutor{
		LibHttpRequest: libHttpRequest,
	}
}

func (e *NetExecutor) Execute() (*tpt.Response, error) {
	client := &http.Client{}
	res, err := client.Do(&e.LibHttpRequest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, se.ErrResponseReadFailed
	}

	return &tpt.Response{
		Header:     res.Header,
		Body:       io.NopCloser(bytes.NewReader(bodyBytes)),
		StatusCode: res.StatusCode,
	}, nil
}
