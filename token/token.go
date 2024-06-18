// Package token defines types and constants for lexical token analysis in a parser.
package token

// TokenType represents the type of lexical tokens.
type TokenType string

// Token represents a lexical token with a specific type and literal value.
type Token struct {
	Type    TokenType // Type is the category of the token.
	Literal string    // Literal is the textual representation of the token.
}

const (
	// ILLEGAL marks an unknown token or character.
	ILLEGAL TokenType = "ILLEGAL"
	// EOF signifies the end of file, indicating no more tokens are available for parsing.
	EOF TokenType = "EOF"

	// IDENT represents identifiers, e.g., variable names.
	IDENT TokenType = "IDENT"
	// INT represents integer literals.
	INT TokenType = "INT"

	// Operators

	ASSIGN   TokenType = "=" // ASSIGN represents the assignment operator.
	PLUS     TokenType = "+" // PLUS represents the addition operator.
	MINUS    TokenType = "-" // MINUS represents the subtraction operator.
	BANG     TokenType = "!" // BANG represents the logical negation operator.
	ASTERISK TokenType = "*" // ASTERISK represents the multiplication operator.
	SLASH    TokenType = "/" // SLASH represents the division operator.

	LT     TokenType = "<"  // LT represents the less-than comparison operator.
	GT     TokenType = ">"  // GT represents the greater-than comparison operator.
	EQ     TokenType = "==" // EQ represents the equality comparison operator.
	NOT_EQ TokenType = "!=" // NOT_EQ represents the inequality comparison operator.

	// Delimiters

	COMMA     TokenType = "," // COMMA represents the comma delimiter.
	SEMICOLON TokenType = ";" // SEMICOLON represents the semicolon delimiter.

	LPAREN TokenType = "(" // LPAREN represents the left parenthesis.
	RPAREN TokenType = ")" // RPAREN represents the right parenthesis.
	LBRACE TokenType = "{" // LBRACE represents the left brace.
	RBRACE TokenType = "}" // RBRACE represents the right brace.

	// Keywords

	FUNCTION TokenType = "FUNCTION" // FUNCTION represents the 'function' keyword.
	LET      TokenType = "LET"      // LET represents the 'let' keyword.
	IF       TokenType = "IF"       // IF represents the 'if' keyword.
	ELSE     TokenType = "ELSE"     // ELSE represents the 'else' keyword.
	RETURN   TokenType = "RETURN"   // RETURN represents the 'return' keyword.
	TRUE     TokenType = "TRUE"     // TRUE represents the 'true' keyword.
	FALSE    TokenType = "FALSE"    // FALSE represents the 'false' keyword.
)

// keywords maps string literals to their corresponding TokenType.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// LookUpIndent returns the TokenType for a given identifier.
// If the identifier is a recognized keyword, it returns the corresponding TokenType;
// otherwise, it defaults to IDENT.
func LookUpIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// NewToken creates a new Token of the given TokenType with the literal value derived from the byte ch.
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
