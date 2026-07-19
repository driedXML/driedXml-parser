package internal

import (
	"context"
	"strings"

	"github.com/driedxml/parser/api"
)

func ReadWhiteSpace(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	var stringBld strings.Builder
	var counter int
	var err error
	var r []byte

	var position = api.Position{
		Line:   p.line,
		Column: p.pos,
	}

	for {
		stringBld.Reset()
		for counter = 0; counter < MaxTokenSize; counter++ {
			if r, err = p.rb.Peek(1); err != nil || !isSpace(r[0]) {
				break
			}

			stringBld.WriteByte(r[0])
			_, _ = p.rb.Discard(1)
			if r[0] == '\n' {
				p.pos = 1
				p.line++
			} else {
				p.pos++
			}
		}

		if counter != 0 {
			select {
			case <-ctx.Done():
				return
			case tokChan <- api.WhiteSpace{
				Position: position,
				Value:    stringBld.String(),
			}:
				if counter != MaxTokenSize {
					return
				}
				position = api.Position{
					Line:   p.line,
					Column: p.pos,
				}
				continue
			}
		} else {
			return
		}
	}
}

func isSpace(r byte) bool {
	return r == 0x09 || r == 0x20 || r == 0x0a || r == 0x0d
}
