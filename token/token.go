package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// idenitfiers + literals
	IDENT = "IDENT" // add, x, y, ...
	INT   = "INT"   // 123

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRASE = "{"
	RBRASE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if t_type, ok := keywords[ident]; ok {
		return t_type
	} else {
		return IDENT
	}
}
