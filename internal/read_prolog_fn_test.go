package internal_test

import (
	"context"
	"strings"
	"testing"

	"github.com/driedxml/parser/api"
	"github.com/driedxml/parser/internal"
	"github.com/driedxml/parser/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestReadPrologFn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("ctx is cancelled", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		parser := internal.NewParserDelegate(strings.NewReader(""), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		internal.ReadProlog(parser, ctx, tokChan, errorChan)
		require.Empty(t, tokChan)
		require.Empty(t, errorChan)
	})

	t.Run("prolog read", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		parser := internal.NewParserDelegate(strings.NewReader("<?drxml?>"), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		internal.ReadProlog(parser, ctx, tokChan, errorChan)
		require.NotEmpty(t, tokChan)
		tok := <-tokChan

		decl, ok := tok.(api.DriedXMLDecl)
		require.Truef(t, ok, "expected DriedXMLDecl, got %T", tok)
		require.Equalf(t, decl.Column, 1, "unexpected column value")
		require.Equalf(t, decl.Line, 1, "unexpected line value")
		require.Empty(t, errorChan)
	})
}
