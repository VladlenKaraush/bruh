package lexer

import "bruh/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func (l *Lexer) NextToken() token.Token {
	t := token.Token{}
	t.Literal = string(l.ch)
	switch l.ch {
	case '=':
		t.Type = token.ASSIGN
	case '+':
		t.Type = token.PLUS
	case ',':
		t.Type = token.COMMA
	case ';':
		t.Type = token.SEMICOLON
	case '(':
		t.Type = token.LPAREN
	case ')':
		t.Type = token.RPAREN
	case '{':
		t.Type = token.LBRASE
	case '}':
		t.Type = token.RBRASE
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else {
			t.Type = token.ILLEGAL
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
