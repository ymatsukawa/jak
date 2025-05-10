package sys_error

import (
	"errors"
)

var (
	ErrBuildOptionHeaderNil     = errors.New("option header is nil")
	ErrBuildOptionHeaderInvalid = errors.New("option header is invalid formt")
	ErrBuildOptionBodyNil       = errors.New("option body is nil")
	ErrBuildOptionContextNil    = errors.New("option context is nil")
)
