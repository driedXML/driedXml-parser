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

	bytes, _ := p.rb.Peek(MaxTokenSize)
	if DriedXMLDeclToken == string(bytes) {
		_, _ = p.rb.Discard(len(DriedXMLDeclToken))

		token := api.DriedXMLDecl{
			Position: api.Position{
				Line:   p.line,
				Column: p.pos,
			},
		}
		select {
		case <-ctx.Done():
			return
		case tokChan <- token:
		}
	}
}
