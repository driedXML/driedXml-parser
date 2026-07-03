package parser

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

type EntityRef struct {
	Value string
}

type CharRef struct {
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
