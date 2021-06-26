package lexer

import "github.com/nixii/monkey/token"

type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // current reading position
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if len(lex.input) <= lex.readPosition {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition++
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.ch {
	case '=':
		tok = newToken(token.ASSIGN, lex.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}

	lex.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lex *Lexer) readIdentifier() string {
	firstPos := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[firstPos:lex.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
