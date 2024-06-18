package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL marks that an unkown token/character was encountered
	ILLEGAL = "ILLEGAL"
	// EOF marks the end of file to stop parsing
	EOF = "EOF"

	// Identifiers
	IDENT = "IDENT"
	// Integer literals
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
