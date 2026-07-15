package api

import "context"

type Comment struct {
	Position
	Value string
}

type CDStart struct {
	Position
}

type CDEnd struct {
	Position
}

type CData struct {
	Position
	Value string
}

type DriedXMLDecl struct {
	Position
}

type Reference struct {
	Position
	Value string
}

type CharData struct {
	Position
	Value string
}

type WhiteSpace struct {
	Position
	Value string
}

type StartElement struct {
	Position
	ElementName          ElementName
	NSAttributes         []NSAttribute
	PrefixAttributes     []PrefixAttribute
	UnPrefixedAttributes []UnPrefixedAttribute
}

type EndElement struct {
	Position
	ElementName ElementName
}

type EmptyElement struct {
	Position
	ElementName          ElementName
	NSAttributes         []NSAttribute
	PrefixAttributes     []PrefixAttribute
	UnPrefixedAttributes []UnPrefixedAttribute
}

// Token - any type of DriedXMLDecl Comment CDStart CDEnd CData Reference CharData WhiteSpace StartElement EndElement EmptyElement
type Token any

type Parser interface {
	Run(ctx context.Context, tokChan chan<- Token, errorChan chan<- error)
}
