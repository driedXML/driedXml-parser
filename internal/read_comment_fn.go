package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/driedxml/parser/api"
)

const (
	CommentState0 = iota // in comment
	CommentState1        //	'-' read
	CommentState2        //	'--' read
	CommentState3        //	'-->' complete
)

func ReadComment(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	bytes, _ := p.rb.Peek(CommentStartLength)
	if CommentStart != string(bytes) {
		return
	}
	_, _ = p.rb.Discard(CommentStartLength)

	var commentContext strings.Builder
	commentContext.Write(bytes)
	var position = api.Position{
		Line:   p.line,
		Column: p.pos,
	}
	p.pos += CommentStartLength

	state := CommentState0

	var tokenSize = CommentStartLength
	for {
		for ; tokenSize < MaxTokenSize; tokenSize++ {
			if r, err := p.rb.Peek(1); err != nil {
				select {
				case <-ctx.Done():
					return
				case errorChan <- fmt.Errorf("%w at position %v", ErrEOFParsing, api.Position{
					Line:   p.line,
					Column: p.pos,
				}):
					return
				}
			} else {
				commentContext.WriteByte(r[0])
				_, _ = p.rb.Discard(1)
				if r[0] == '\n' {
					p.pos = 1
					p.line++
				} else {
					p.pos++
				}
				state = nextCommentState(state, r[0])
				if state == CommentState3 {
					break
				}
			}
		}

		select {
		case <-ctx.Done():
			return
		case tokChan <- api.Comment{
			Position: position,
			Value:    commentContext.String(),
		}:
			if state == CommentState3 {
				return
			}
			position = api.Position{
				Line:   p.line,
				Column: p.pos,
			}
			commentContext.Reset()
			tokenSize = 0
			continue
		}
	}
}

func nextCommentState(state int, next byte) int {
	switch state {
	case CommentState0:
		if next == '-' {
			return CommentState1
		}
		return CommentState0
	case CommentState1:
		if next == '-' {
			return CommentState2
		}
		return CommentState0
	case CommentState2:
		switch next {
		case '>':
			return CommentState3
		case '-':
			return CommentState2
		default:
			return CommentState0
		}
	default:
		return CommentState0
	}
}
