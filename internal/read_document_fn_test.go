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

func TestReadDocumentFn(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("ctx is cancelled", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		parser := internal.NewParserDelegate(strings.NewReader(""), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		internal.ReadDocument(parser, ctx, tokChan, errorChan)

		require.Empty(t, tokChan)
		require.Empty(t, errorChan)
	})

	t.Run("ctx is not cancelled", func(t *testing.T) {
		mockParserDelegate := mocks.NewMockParserDelegate(ctrl)
		mockParserDelegate.EXPECT().ReadProlog(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

		parser := internal.NewParserDelegate(strings.NewReader(""), mockParserDelegate)

		tokChan := make(chan api.Token, 1)
		errorChan := make(chan error, 1)

		ctx := context.Background()

		internal.ReadDocument(parser, ctx, tokChan, errorChan)
		require.Empty(t, tokChan)
		require.Empty(t, errorChan)
	})
}
