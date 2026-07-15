package internal

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/driedxml/parser/api"
	"github.com/driedxml/parser/internal/utils"
)

func NewParser(r io.Reader) *Parser {
	return NewParserDelegate(r, createDelegate())
}

func NewParserDelegate(r io.Reader, parserDelegate ParserDelegate) *Parser {
	parser := &Parser{
		parserDelegate: parserDelegate,
		pos:            1,
		line:           1,
	}
	if rb, ok := r.(*bufio.Reader); ok {
		parser.rb = rb
	} else {
		parser.rb = bufio.NewReader(r)
	}
	return parser
}

type ProcessingFn func(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error)

type Parser struct {
	rb             *bufio.Reader
	nsMap          map[string]utils.Stack // stack string
	defPrefix      utils.Stack            // stack []string
	parserDelegate ParserDelegate
	pos            int
	line           int
}

func (p *Parser) Run(ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	p.parserDelegate.ReadDocument(p, ctx, tokChan, errorChan)
}

func (p *Parser) PushNS(attributes []api.NSAttribute) {
	var prefixes []string
	for _, attr := range attributes {
		prefixes = append(prefixes, attr.Prefix)
		if nsStack, ok := p.nsMap[attr.Prefix]; ok {
			nsStack.Push(attr.Space)
		} else {
			stack := utils.NewStack()
			stack.Push(attr.Space)
			p.nsMap[attr.Prefix] = stack
		}
	}
	p.defPrefix.Push(prefixes)
}

func (p *Parser) PopNS() error {
	if prefixes, err := p.defPrefix.Peek(); err != nil {
		return err
	} else {
		for _, prefix := range prefixes.([]string) {
			if stack, ok := p.nsMap[prefix]; ok {
				if _, err := stack.Pop(); err != nil {
					return fmt.Errorf("%w - stack %s", ErrImplementation, err.Error())
				}
			} else {
				return fmt.Errorf("%w prefix %s is unknown", ErrImplementation, prefix)
			}
		}
	}

	if _, err := p.defPrefix.Pop(); err != nil {
		return fmt.Errorf("%w - defPrefix %s", ErrImplementation, err.Error())
	}
	return nil
}

func (p *Parser) Namespace(prefix string) (string, error) {
	if value, ok := p.nsMap[prefix]; ok {
		if ns, err := value.Peek(); err == nil {
			return ns.(string), nil
		} else {
			return "", err
		}
	}
	return "", nil
}
