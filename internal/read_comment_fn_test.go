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

func TestReadCommentFn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("ctx is cancelled", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		parser := internal.NewParserDelegate(strings.NewReader(""), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		internal.ReadComment(parser, ctx, tokChan, errorChan)
		require.Empty(t, tokChan)
		require.Empty(t, errorChan)
	})

	t.Run("comment read", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		value := "<!--comment-asdas--c-->"
		parser := internal.NewParserDelegate(strings.NewReader(value), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		internal.ReadComment(parser, ctx, tokChan, errorChan)
		require.NotEmpty(t, tokChan)
		tok := <-tokChan

		ws, ok := tok.(api.Comment)
		require.Truef(t, ok, "expected Comment, got %T", tok)
		require.Equalf(t, value, ws.Value, "unexpected comment value")
		require.Equalf(t, ws.Column, 1, "unexpected column value")
		require.Equalf(t, ws.Line, 1, "unexpected line value")
		require.Empty(t, errorChan)
	})

	t.Run("long comment read", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)

		var sb1 strings.Builder
		sb1.WriteString("<!--")
		for i := 0; i < internal.MaxTokenSize-1-internal.CommentStartLength; i++ {
			sb1.WriteByte(' ')
		}
		sb1.WriteByte('\n')
		sb1String := sb1.String()

		var sb2 strings.Builder
		for i := 0; i < internal.MaxTokenSize-1-internal.CommentEndLength; i++ {
			sb2.WriteByte('c')
		}
		sb2.WriteByte('\n')
		sb2.WriteString("-->")
		sb2String := sb2.String()

		parser := internal.NewParserDelegate(strings.NewReader(
			strings.Join([]string{sb1String, sb2String}, ""),
		), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		go func() {
			internal.ReadComment(parser, ctx, tokChan, errorChan)
		}()

		tok := <-tokChan

		ws, ok := tok.(api.Comment)
		require.Truef(t, ok, "expected Comment, got %T", tok)
		require.Equalf(t, sb1String, ws.Value, "unexpected comment value")
		require.Equalf(t, ws.Column, 1, "unexpected column value")
		require.Equalf(t, ws.Line, 1, "unexpected line value")
		require.Empty(t, errorChan)

		tok = <-tokChan
		ws, ok = tok.(api.Comment)
		require.Truef(t, ok, "expected Comment, got %T", tok)
		require.Equalf(t, sb2String, ws.Value, "unexpected comment value")
		require.Equalf(t, ws.Column, 1, "unexpected column value")
		require.Equalf(t, ws.Line, 2, "unexpected line value")
		require.Empty(t, errorChan)
	})

	t.Run("comment unexpected eof", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		value := "<!--comment-asdas--c--"
		parser := internal.NewParserDelegate(strings.NewReader(value), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		internal.ReadComment(parser, ctx, tokChan, errorChan)
		require.Empty(t, tokChan)
		require.NotEmpty(t, errorChan)

		err := <-errorChan

		require.ErrorIsf(t, err, internal.ErrBase, "expected ErrBase %T", err)
		require.ErrorIsf(t, err, internal.ErrEOFParsing, "expected ErrEOFParsing %T", err)
	})
}
