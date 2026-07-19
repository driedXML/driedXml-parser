package internal

const MaxLexemeSize = 9 // DriedXMLDecl or CDStart

const DriedXMLDecl = "<?drxml?>"
const DriedXMLDeclLength = len(DriedXMLDecl)

const CommentStart = "<!--"
const CommentStartLength = len(CommentStart)

const CommentEnd = "-->"
const CommentEndLength = len(CommentEnd)

const StartElement = "<"

const MaxTokenSize = 1000
