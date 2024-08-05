package lexer

import (
	"log"
	"plutonium/token"
)

type Lexer struct {
	Content      string
	Position     int
	ReadPosition int
	Char         byte
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '$'
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.Char == ' ' || lexer.Char == '\t' || lexer.Char == '\n' || lexer.Char == '\r' || lexer.Char == '#' {
		lexer.readChar()
	}
}

func New(input string) *Lexer {
	lexer := &Lexer{Content: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readString() string {
	position := lexer.Position + 1
	for {
		lexer.readChar()
		if lexer.Char == '"' {
			break
		}
	}
	str := lexer.Content[position:lexer.Position]
	lexer.readChar()
	return str
}

func (lexer *Lexer) readChar() {
	if lexer.ReadPosition >= len(lexer.Content) {
		lexer.Char = 0
	} else {
		lexer.Char = lexer.Content[lexer.ReadPosition]
	}

	lexer.Position = lexer.ReadPosition
	lexer.ReadPosition += 1
}

func (lexer *Lexer) readIdentifier() string {
	pos := lexer.Position
	for isAlpha(lexer.Char) {
		lexer.readChar()
	}

	return lexer.Content[pos:lexer.Position]
}

func (lexer *Lexer) readNumber() (string, bool) {
	position := lexer.Position
	isFloat := false
	for isDigit(lexer.Char) {
		lexer.readChar()
	}
	if lexer.Char == '.' {
		lexer.readChar()
		for isDigit(lexer.Char) {
			lexer.readChar()
		}
		isFloat = true
	}
	return lexer.Content[position:lexer.Position], isFloat
}

func (lexer *Lexer) peek() byte {
	if lexer.ReadPosition >= len(lexer.Content) {
		return 0
	}

	return lexer.Content[lexer.ReadPosition]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (lexer *Lexer) Consume() token.Token {
	var tok token.Token

	lexer.skipWhiteSpace()

	switch lexer.Char {
	case '=':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.EQUALS, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.ASSIGN, lexer.Char)
		}
	case '!':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQUALS, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.NOT, lexer.Char)
		}
	case '+':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.PLUS_ASSIGN, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.PLUS, lexer.Char)
		}
	case '-':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.MINUS_ASSIGN, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.MINUS, lexer.Char)
		}
	case '*':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.MULTIPLY_ASSIGN, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.MULTIPLY, lexer.Char)
		}
	case '/':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.DIVIDE_ASSIGN, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.DIVIDE, lexer.Char)
		}
	case '<':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.LESS_THAN_EQUALS, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.LT, lexer.Char)
		}
	case '>':
		if lexer.peek() == '=' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.GREATER_THAN_EQUALS, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.GT, lexer.Char)
		}
	case '|':
		if lexer.peek() == '|' {
			char := lexer.Char
			lexer.readChar()
			tok = token.Token{Type: token.LOGICAL_OR, Literal: string(char) + string(lexer.Char)}
		} else {
			tok = newToken(token.PIPE, lexer.Char)
		}
	case '(':
		tok = newToken(token.OPEN_PAREN, lexer.Char)
	case ')':
		tok = newToken(token.CLOSE_PAREN, lexer.Char)
	case '[':
		tok = newToken(token.OPEN_BRACKET, lexer.Char)
	case ']':
		tok = newToken(token.CLOSE_BRACKET, lexer.Char)
	case '{':
		tok = newToken(token.OPEN_BRACE, lexer.Char)
	case '}':
		tok = newToken(token.CLOSE_BRACE, lexer.Char)
	case '.':
		tok = newToken(token.DOT, lexer.Char)
	case '?':
		tok = newToken(token.QUESTION, lexer.Char)
	case ':':
		tok = newToken(token.COLON, lexer.Char)
	case ';':
		tok = newToken(token.SEMI_COLON, lexer.Char)
	case ',':
		tok = newToken(token.COMMA, lexer.Char)
	case '"':
		tok.Type = token.TYPE_STRING
		tok.Literal = lexer.readString()
		return tok
	case 0:
		tok.Literal = "EOF"
		tok.Type = token.EOF
	default:
		switch {
		case isAlpha(lexer.Char):
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.ToKeywords(tok.Literal)
			return tok
		case isDigit(lexer.Char):
			literal, isFloat := lexer.readNumber()
			tok.Literal = literal
			tok.Type = token.TYPE_INT
			if isFloat {
				tok.Type = token.TYPE_FLOAT
				return tok
			}
			return tok
		default:
			log.Fatalln("Illegal Token")
		}
	}

	lexer.readChar()
	return tok
}
