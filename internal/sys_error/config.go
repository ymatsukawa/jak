package sys_error

import (
	"errors"
)

var (
	ErrConfigIsNil = errors.New("config is nil")
)
