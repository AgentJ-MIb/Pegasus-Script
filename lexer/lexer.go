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

func (lexer *Lexer) LineContent(line int) string {
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

func Error(lexer *Lexer, char byte) {
	lineContent := lexer.LineContent(lexer.Line)

	fmt.Printf(inc.Red+"Error"+inc.Reset+": [line: "+inc.Blue+"%d"+inc.Reset+", column: "+inc.Cyan+"%d"+inc.Reset+"] "+inc.Yellow+"Unexpected character '%c'\n"+inc.Reset, lexer.Line, lexer.Column, char)

	fmt.Printf(" %d | %s\n", lexer.Line, lineContent)
	caretPosition := lexer.Column - 1
	if caretPosition < 0 {
		caretPosition = 0
	}
	fmt.Printf("   | %s^\n", strings.Repeat(" ", caretPosition))
	lexer.ReadChar()
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsAlpha(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '$'
}

func (lexer *Lexer) SkipWhiteSpace() {
	for lexer.Char == ' ' || lexer.Char == '\t' || lexer.Char == '\n' || lexer.Char == '\r' || lexer.Char == '#' {
		lexer.ReadChar()
	}
}

func New(input string) *Lexer {
	lexer := &Lexer{Content: input}
	lexer.ReadChar()
	lexer.Line = 1
	return lexer
}

func (lexer *Lexer) ReadString() string {
	position := lexer.Position + 1
	for {
		lexer.ReadChar()
		if lexer.Char == '"' {
			break
		}
	}
	str := lexer.Content[position:lexer.Position]
	lexer.ReadChar()
	return str
}

func (lexer *Lexer) ReadChar() {
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

func (lexer *Lexer) ReadIdentifier() string {
	pos := lexer.Position
	for IsAlpha(lexer.Char) {
		lexer.ReadChar()
	}

	return lexer.Content[pos:lexer.Position]
}

func (lexer *Lexer) readNumber() (string, bool) {
	position := lexer.Position
	isFloat := false
	for IsDigit(lexer.Char) {
		lexer.ReadChar()
	}
	if lexer.Char == '.' {
		lexer.ReadChar()
		for IsDigit(lexer.Char) {
			lexer.ReadChar()
		}
		isFloat = true
	}
	return lexer.Content[position:lexer.Position], isFloat
}

func (lexer *Lexer) Peek() byte {
	if lexer.ReadPosition >= len(lexer.Content) {
		return 0
	}

	return lexer.Content[lexer.ReadPosition]
}

func NewToken(tokenType TokenType, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (lexer *Lexer) Consume() Token {
	lexer.SkipWhiteSpace()

	var tok Token

	if lexer.Char == 0 {
		return Token{Type: EOF, Literal: "EOF"}
	}

	if tokenType := GetOperatorType(lexer.Char); tokenType != ERROR {
		if tokenType == TYPE_STRING {
			tok.Type = TYPE_STRING
			tok.Literal = lexer.ReadString()
			return tok
		}
		tok = NewToken(tokenType, lexer.Char)
		if lexer.Peek() == '=' {
			lexer.ReadChar()
			tok.Type = GetAssignmentType(lexer.Char)
			tok.Literal = TokenTypeToString(tok.Type)
		}
		lexer.ReadChar()
		return tok
	}

	switch {
	case IsAlpha(lexer.Char):
		tok.Literal = lexer.ReadIdentifier()
		tok.Type = ToKeywords(tok.Literal)
		return tok
	case IsDigit(lexer.Char):
		literal, isFloat := lexer.readNumber()
		tok.Literal = literal
		tok.Type = TYPE_INT
		if isFloat {
			tok.Type = TYPE_FLOAT
			return tok
		}
		return tok
	default:
		tok = NewToken(ERROR, lexer.Char)
		Error(lexer, lexer.Char)
	}

	lexer.ReadChar()
	return tok
}
