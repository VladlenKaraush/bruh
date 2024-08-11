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
	l.skipWhitespace()
	t.Literal = string(l.ch)
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			t.Type = token.EQ
			t.Literal = "=="
		} else {
			t.Type = token.ASSIGN
		}
	case '+':
		t.Type = token.PLUS
	case '-':
		t.Type = token.MINUS
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			t.Type = token.NOT_EQ
			t.Literal = "!="
		} else {
			t.Type = token.BANG
		}
	case '*':
		t.Type = token.ASTERISK
	case '/':
		t.Type = token.SLASH
	case '>':
		t.Type = token.GT
	case '<':
		t.Type = token.LT
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
		} else if isDigit(l.ch) {
			t.Literal = l.readInt()
			t.Type = token.INT
			return t
		} else {
			t.Type = token.ILLEGAL
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readInt() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
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

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
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
