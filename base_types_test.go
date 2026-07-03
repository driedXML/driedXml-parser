package parser_test

import (
	"testing"

	"github.com/driedXML/parser"
	"github.com/stretchr/testify/require"
)

func TestAttributeValue(t *testing.T) {
	prefixed := parser.PrefixAttribute{
		Prefix:    "Prefix",
		LocalPart: "LocalPart",
		Value:     "Value",
	}

	unPrefixed := parser.UnPrefixedAttribute{
		LocalPart: "LocalPart",
		Value:     "Value",
	}

	vl := make(parser.ValueAttribute, 10)

	require.True(t, true)

}
