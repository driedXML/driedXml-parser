package internal

import (
	"io"

	"github.com/driedxml/parser/api"
)

type Parser struct {
	r io.ByteReader
}

func (p *Parser) NextToken() (api.Token, error) {
	return nil, io.EOF
}
