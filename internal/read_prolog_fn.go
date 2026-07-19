package internal

import (
	"context"

	"github.com/driedxml/parser/api"
)

func ReadProlog(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	bytes, _ := p.rb.Peek(DriedXMLDeclLength)
	if DriedXMLDecl == string(bytes) {
		_, _ = p.rb.Discard(DriedXMLDeclLength)

		token := api.DriedXMLDecl{
			Position: api.Position{
				Line:   p.line,
				Column: p.pos,
			},
		}
		p.pos += DriedXMLDeclLength
		select {
		case <-ctx.Done():
			return
		case tokChan <- token:
		}
	}
}
