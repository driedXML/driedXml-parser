package internal

import (
	"errors"
	"fmt"
)

var ErrBase = errors.New("parsing error")

var ErrImplementation = fmt.Errorf("%w - implementation problem", ErrBase)

var ErrParsing = fmt.Errorf("%w - parsing error", ErrBase)
