package request_control

import (
	"context"
	"net/http"
	"sync"
)

type Concurrence interface {
	RunConcurrence() error
}

type RequestConcurrence struct {
	LibHttpRequests []http.Request
	IgnoreFail      bool
	Context         context.Context
}

func NewRequestConcurrence() *RequestConcurrence {
	return &RequestConcurrence{}
}

func (c *RequestConcurrence) RunConcurrence() error {
	requestCount := len(c.LibHttpRequests)
	workerCount := requestCount

	jobs := make(chan http.Request, requestCount)
	errCh := make(chan error, 1)

	workerContext, cancel := context.WithCancel(c.Context)
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		request := c.LibHttpRequests[i]
		go c.worker(workerContext, &wg, jobs, errCh, request)
	}

	return nil
}

func (c *RequestConcurrence) worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	jobs <-chan http.Request,
	errCh chan<- error,
	targetRequest http.Request,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case request, ok := <-jobs:
			if !ok {
				return
			}
			res, err := nc.
	}
}

func (c *RequestConcurrence) sendJobs() {
	return
}
