package internal

import (
	"errors"
	"fmt"
)

var ErrBase = errors.New("parsing error")

var ErrImplementation = fmt.Errorf("%w - implementation problem", ErrBase)

var ErrEOFParsing = fmt.Errorf("%w - unexpected end of file", ErrBase)
