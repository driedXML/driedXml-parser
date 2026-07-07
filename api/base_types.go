package api

import (
	"fmt"
)

type Position struct {
	Line   int
	Column int
}

// A SyntaxError represents a syntax error in the XML input stream.
type SyntaxError struct {
	Err error
	Pos Position
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("XML syntax error at position at position %v:%v - %v", e.Pos.Line, e.Pos.Column, e.Err)
}

func (e *SyntaxError) Unwrap() error { return e.Err }

// A ElementName represents an driedXML ElementName
type ElementName struct {
	Prefix, LocalPart string
}

// A UnPrefixedName represents an driedXML UnPrefixedName
type UnPrefixedName struct {
	LocalPart string
}

type NSAttribute struct {
	Prefix string
	Space  string
}

type ValueAttribute interface {
	PrefixAttribute | UnPrefixedAttribute
}

type PrefixAttribute struct {
	Prefix    string
	LocalPart string
	Value     string
}

type UnPrefixedAttribute struct {
	LocalPart string
	Value     string
}

type EntityRef struct {
	Value string
}

type CharRef struct {
	Value string
}
