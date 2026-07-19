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

func TestReadWhiteSpaceFn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("ctx is cancelled", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		parser := internal.NewParserDelegate(strings.NewReader(""), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		internal.ReadWhiteSpace(parser, ctx, tokChan, errorChan)
		require.Empty(t, tokChan)
		require.Empty(t, errorChan)
	})

	t.Run("white space read", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		parser := internal.NewParserDelegate(strings.NewReader("     \n\t\r   \t\n"), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		internal.ReadWhiteSpace(parser, ctx, tokChan, errorChan)
		require.NotEmpty(t, tokChan)
		tok := <-tokChan

		ws, ok := tok.(api.WhiteSpace)
		require.Truef(t, ok, "expected WhiteSpace, got %T", tok)
		require.Equalf(t, "     \n\t\r   \t\n", ws.Value, "unexpected white space value")
		require.Equalf(t, ws.Column, 1, "unexpected column value")
		require.Equalf(t, ws.Line, 1, "unexpected line value")
		require.Empty(t, errorChan)
	})

	t.Run("long white space read", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)

		var sb1 strings.Builder
		for i := 0; i < 999; i++ {
			sb1.WriteByte(' ')
		}
		sb1.WriteByte('\n')
		sb1String := sb1.String()

		var sb2 strings.Builder
		for i := 0; i < 999; i++ {
			sb2.WriteByte('\t')
		}
		sb2.WriteByte('\n')
		sb2String := sb2.String()

		parser := internal.NewParserDelegate(strings.NewReader(
			strings.Join([]string{sb1String, sb2String}, ""),
		), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		go func() {
			internal.ReadWhiteSpace(parser, ctx, tokChan, errorChan)
		}()

		tok := <-tokChan

		ws, ok := tok.(api.WhiteSpace)
		require.Truef(t, ok, "expected WhiteSpace, got %T", tok)
		require.Equalf(t, sb1String, ws.Value, "unexpected white space value")
		require.Equalf(t, ws.Column, 1, "unexpected column value")
		require.Equalf(t, ws.Line, 1, "unexpected line value")
		require.Empty(t, errorChan)

		tok = <-tokChan
		ws, ok = tok.(api.WhiteSpace)
		require.Truef(t, ok, "expected WhiteSpace, got %T", tok)
		require.Equalf(t, sb2String, ws.Value, "unexpected white space value")
		require.Equalf(t, ws.Column, 1, "unexpected column value")
		require.Equalf(t, ws.Line, 2, "unexpected line value")
		require.Empty(t, errorChan)
	})

}
