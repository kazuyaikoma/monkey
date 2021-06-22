package lexer

import (
	"github.com/nixii/monkey/token"
	"testing"
)

func TestNextToken(test *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	inp := New(input)

	for i, testToken := range tests {
		tok := inp.NextToken()

		if tok.Type != testToken.expectedType {
			test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, testToken.expectedType, tok.Type)
		}

		if tok.Literal != testToken.expectedLiteral {
			test.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, testToken.expectedLiteral, tok.Literal)
		}
	}
}
