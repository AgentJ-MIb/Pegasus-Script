package lexer

import (
	"fmt"
	"plutonium/inc"
	"strings"
)

type Lexer struct {
	Content      string
	Position     int
	ReadPosition int
	Char         byte
	Line         int
	Column       int
}

func (lexer *Lexer) lineContent(line int) string {
	start := 0
	cLine := 1

	for cLine != line && start < len(lexer.Content) {
		if lexer.Content[start] == '\n' {
			cLine++
		}
		start++
	}
	end := start
	for end < len(lexer.Content) && lexer.Content[end] != '\n' {
		end++
	}
	return lexer.Content[start:end]
}

func error(lexer *Lexer, char byte) {
	lineContent := lexer.lineContent(lexer.Line)

	fmt.Printf(inc.Red+"Error"+inc.Reset+": [line: "+inc.Blue+"%d"+inc.Reset+", column: "+inc.Cyan+"%d"+inc.Reset+"] "+inc.Yellow+"Unexpected character '%c'\n"+inc.Reset, lexer.Line, lexer.Column, char)

	fmt.Printf(" %d | %s\n", lexer.Line, lineContent)
	caretPosition := lexer.Column - 1
	if caretPosition < 0 {
		caretPosition = 0
	}
	fmt.Printf("   | %s^\n", strings.Repeat(" ", caretPosition))
	lexer.readChar()
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
	lexer.Line = 1
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

	if lexer.Char == '\n' {
		lexer.Line++
		lexer.Column = 0
	} else {
		lexer.Column++
	}

	lexer.Position = lexer.ReadPosition
	lexer.ReadPosition++
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

func newToken(tokenType TokenType, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (lexer *Lexer) Consume() Token {
	lexer.skipWhiteSpace()

	var tok Token

	if lexer.Char == 0 {
		return Token{Type: EOF, Literal: "EOF"}
	}

	if tokenType := GetOperatorType(lexer.Char); tokenType != ERROR {
		if tokenType == TYPE_STRING {
			tok.Type = TYPE_STRING
			tok.Literal = lexer.readString()
			return tok
		}
		tok = newToken(tokenType, lexer.Char)
		if lexer.peek() == '=' {
			lexer.readChar()
			tok.Type = GetAssignmentType(lexer.Char)
			tok.Literal = TokenTypeToString(tok.Type) // This will use the TokenTypeToString function
		}
		lexer.readChar()
		return tok
	}

	switch {
	case isAlpha(lexer.Char):
		tok.Literal = lexer.readIdentifier()
		tok.Type = ToKeywords(tok.Literal)
		return tok
	case isDigit(lexer.Char):
		literal, isFloat := lexer.readNumber()
		tok.Literal = literal
		tok.Type = TYPE_INT
		if isFloat {
			tok.Type = TYPE_FLOAT
			return tok
		}
		return tok
	default:
		tok = newToken(ERROR, lexer.Char)
		error(lexer, lexer.Char)
	}

	lexer.readChar()
	return tok
}
