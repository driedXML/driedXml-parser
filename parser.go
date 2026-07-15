package parser

import (
	"io"

	"github.com/driedxml/parser/api"
	"github.com/driedxml/parser/internal"
)

func NewParser(r io.Reader) api.Parser {
	return internal.NewParser(r)
}
