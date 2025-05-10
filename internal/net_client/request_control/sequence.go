package request_control

import (
	"net/http"

	nc "github.com/ymatsukawa/jak/internal/net_client"
)

type Sequence interface {
	RunSequence() error
}

type SequenceRequest struct {
	LibHttpRequest []http.Request
	IgnoreFail     bool
}

func NewRequestSequence(requests []http.Request, ignoreFail bool) *SequenceRequest {
	return &SequenceRequest{
		LibHttpRequest: requests,
		IgnoreFail:     ignoreFail,
	}
}

func (s *SequenceRequest) RunSequence() error {
	for _, req := range s.LibHttpRequest {
		executor := nc.NewNetExecutor(req)
		_, err := executor.Execute()
		if s.IgnoreFail {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}
