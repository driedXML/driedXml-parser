package internal

import (
	"context"

	"github.com/driedxml/parser/api"
)

func ReadDocument(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	p.parserDelegate.ReadProlog(p, ctx, tokChan, errorChan)
}
