package api

type Comment struct {
	Value string
}

type CDStart struct {
}

type CDEnd struct {
}

type CData struct {
	Value string
}

type DriedXMLDecl struct {
}

type Reference struct {
	Value string
}

type CharData struct {
	Value string
}

type WhiteSpace struct {
	Value string
}

type StartElement struct {
	ElementName          ElementName
	NSAttributes         []NSAttribute
	PrefixAttributes     []PrefixAttribute
	UnPrefixedAttributes []UnPrefixedAttribute
}

type EndElement struct {
	ElementName ElementName
}

type EmptyElement struct {
	ElementName          ElementName
	NSAttributes         []NSAttribute
	PrefixAttributes     []PrefixAttribute
	UnPrefixedAttributes []UnPrefixedAttribute
}

// Token - any type of DriedXMLDecl Comment CDStart CDEnd CData Reference CharData WhiteSpace StartElement EndElement EmptyElement
type Token any

type Parser interface {
	// NextToken returns the next driedXML token in the input stream.
	// At the end of the input stream, Token returns nil, io.EOF.
	NextToken() (Token, error)
}
