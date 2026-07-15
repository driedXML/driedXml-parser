package internal

import (
	"context"

	_ "github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"

	"github.com/driedxml/parser/api"
)

//go:generate mockgen -destination "./mocks/$GOFILE" -package mocks . ParserDelegate
type ParserDelegate interface {
	ReadDocument(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error)
	ReadProlog(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error)
}

type delegate struct {
	readDocumentFn ProcessingFn
	readPrologFn   ProcessingFn
}

func (d *delegate) ReadDocument(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	d.readDocumentFn(p, ctx, tokChan, errorChan)
}

func (d *delegate) ReadProlog(p *Parser, ctx context.Context, tokChan chan<- api.Token, errorChan chan<- error) {
	d.readPrologFn(p, ctx, tokChan, errorChan)
}

func createDelegate() ParserDelegate {
	return &delegate{
		readDocumentFn: ReadDocument,
		readPrologFn:   ReadProlog,
	}
}
