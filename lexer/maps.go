package lexer

import "fmt"

var operatorMap = map[byte]TokenType{
	'=': ASSIGN,
	'!': NOT,
	'+': PLUS,
	'-': MINUS,
	'*': MULTIPLY,
	'/': DIVIDE,
	'<': LT,
	'>': GT,
	'|': PIPE,
	'(': OPEN_PAREN,
	')': CLOSE_PAREN,
	'[': OPEN_BRACKET,
	']': CLOSE_BRACKET,
	'{': OPEN_BRACE,
	'}': CLOSE_BRACE,
	'.': DOT,
	'?': QUESTION,
	':': COLON,
	';': SEMI_COLON,
	',': COMMA,
	'"': TYPE_STRING,
}

var assignmentMap = map[byte]TokenType{
	'=': EQUALS,
	'+': PLUS_ASSIGN,
	'-': MINUS_ASSIGN,
	'*': MULTIPLY_ASSIGN,
	'/': DIVIDE_ASSIGN,
	'<': LESS_THAN_EQUALS,
	'>': GREATER_THAN_EQUALS,
	'|': LOGICAL_OR,
}

var keywordMap = map[string]TokenType{
	"int":      TYPE_INT,
	"float":    TYPE_FLOAT,
	"bool":     TYPE_BOOLEAN,
	"string":   TYPE_STRING,
	"try":      KEYWORDS_TRY,
	"catch":    KEYWORDS_CATCH,
	"import":   KEYWORDS_IMPORT,
	"true":     KEYWORDS_TRUE,
	"false":    KEYWORDS_FALSE,
	"function": KEYWORDS_FUNCTION,
	"if":       KEYWORDS_IF,
	"let":      KEYWORDS_VAR,
	"const":    KEYWORDS_CONST,
	"else":     KEYWORDS_ELSE,
	"return":   KEYWORDS_RETURN,
	"for":      KEYWORDS_FOR,
	"switch":   KEYWORDS_SWITCH,
	"case":     KEYWORDS_CASE,
	"default":  KEYWORDS_DEFAULT,
}

func TokenTypeToString(t TokenType) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case MINUS:
		return "MINUS"
	case PLUS:
		return "PLUS"
	case MULTIPLY:
		return "MULTIPLY"
	case DIVIDE:
		return "DIVIDE"
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	case OPEN_BRACE:
		return "OPEN_BRACE"
	case CLOSE_BRACE:
		return "CLOSE_BRACE"
	case DOT:
		return "DOT"
	case QUESTION:
		return "QUESTION"
	case COLON:
		return "COLON"
	case SEMI_COLON:
		return "SEMI_COLON"
	case COMMA:
		return "COMMA"
	case NOT:
		return "NOT"
	case PIPE:
		return "PIPE"
	case TILDE:
		return "TILDE"
	case AMSPERSAND:
		return "AMSPERSAND"
	case GT:
		return "GT"
	case LT:
		return "LT"
	case LOGICAL_OR:
		return "LOGICAL_OR"
	case LOGICAL_AND:
		return "LOGICAL_AND"
	case EQUALS:
		return "EQUALS"
	case NOT_EQUALS:
		return "NOT_EQUALS"
	case LESS_THAN_EQUALS:
		return "LESS_THAN_EQUALS"
	case GREATER_THAN_EQUALS:
		return "GREATER_THAN_EQUALS"
	case ASSIGN:
		return "ASSIGN"
	case PLUS_ASSIGN:
		return "PLUS_ASSIGN"
	case MINUS_ASSIGN:
		return "MINUS_ASSIGN"
	case MULTIPLY_ASSIGN:
		return "MULTIPLY_ASSIGN"
	case DIVIDE_ASSIGN:
		return "DIVIDE_ASSIGN"
	case IDENTIFIER:
		return "IDENTIFIER"
	case TYPE_STRING:
		return "TYPE_STRING"
	case TYPE_INT:
		return "TYPE_INT"
	case TYPE_FLOAT:
		return "TYPE_FLOAT"
	case TYPE_BOOLEAN:
		return "TYPE_BOOLEAN"
	case KEYWORDS_FUNCTION:
		return "KEYWORDS_FUNCTION"
	case KEYWORDS_IF:
		return "KEYWORDS_IF"
	case KEYWORDS_ELSE:
		return "KEYWORDS_ELSE"
	case KEYWORDS_RETURN:
		return "KEYWORDS_RETURN"
	case KEYWORDS_VAR:
		return "KEYWORDS_VAR"
	case KEYWORDS_CONST:
		return "KEYWORDS_CONST"
	case KEYWORDS_SWITCH:
		return "KEYWORDS_SWITCH"
	case KEYWORDS_FOR:
		return "KEYWORDS_FOR"
	case KEYWORDS_CASE:
		return "KEYWORDS_CASE"
	case KEYWORDS_DEFAULT:
		return "KEYWORDS_DEFAULT"
	case KEYWORDS_TRUE:
		return "KEYWORDS_TRUE"
	case KEYWORDS_FALSE:
		return "KEYWORDS_FALSE"
	case KEYWORDS_TRY:
		return "KEYWORDS_TRY"
	case KEYWORDS_CATCH:
		return "KEYWORDS_CATCH"
	case KEYWORDS_IMPORT:
		return "KEYWORDS_IMPORT"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", t)
	}
}

func ToKeywords(value string) TokenType {
	if val, ok := keywordMap[value]; ok {
		return val
	}
	return IDENTIFIER
}

func GetOperatorType(char byte) TokenType {
	if tokenType, ok := operatorMap[char]; ok {
		return tokenType
	}
	return ERROR
}

func GetAssignmentType(char byte) TokenType {
	if tokenType, ok := assignmentMap[char]; ok {
		return tokenType
	}
	return ERROR
}
