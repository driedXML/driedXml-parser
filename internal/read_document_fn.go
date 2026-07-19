package internal

import (
	"context"
	"fmt"
	"io"

	"github.com/driedxml/parser/api"
)

func ReadDocument(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	p.parserDelegate.ReadProlog(p, ctx, tokChan, errorChan)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		bytes, err := p.rb.Peek(CommentStartLength)
		if len(bytes) == 0 && err == io.EOF {
			return
		}
		if string(bytes) == CommentStart {
			p.parserDelegate.ReadComment(p, ctx, tokChan, errorChan)
			continue
		}
		errorChan <- fmt.Errorf("unexpected EOF")
	}

	//p.parserDelegate.ReadMisk(p, ctx, tokChan, errorChan)

}
