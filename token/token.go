package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF TokenType = iota

	MINUS    // -
	PLUS     // +
	MULTIPLY // *
	DIVIDE   // /

	OPEN_PAREN          // (
	CLOSE_PAREN         // )
	OPEN_BRACKET        // [
	CLOSE_BRACKET       // ]
	OPEN_BRACE          // {
	CLOSE_BRACE         // }
	DOT                 // .
	QUESTION            // ?
	COLON               // :
	SEMI_COLON          // ;
	COMMA               // ,
	NOT                 // !
	PIPE                // |
	TILDE               // ~
	AMSPERSAND          // &
	GT                  // >
	LT                  // <
	LOGICAL_OR          // ||
	LOGICAL_AND         // &&
	EQUALS              // ==
	NOT_EQUALS          // !=
	LESS_THAN_EQUALS    // <=
	GREATER_THAN_EQUALS // >=
	ASSIGN              // =
	PLUS_ASSIGN         // +=
	MINUS_ASSIGN        // -=
	MULTIPLY_ASSIGN     // *=
	DIVIDE_ASSIGN       // /=

	IDENTIFIER
	TYPE_STRING  // "Hello"
	TYPE_INT     // 10
	TYPE_FLOAT   // 10.0
	TYPE_BOOLEAN // True or False

	KEYWORDS_FUNCTION // function
	KEYWORDS_IF       // if
	KEYWORDS_ELSE     // else
	KEYWORDS_RETURN   // return
	KEYWORDS_VAR      // let
	KEYWORDS_CONST    // const
	KEYWORDS_SWITCH   // switch
	KEYWORDS_FOR      // for
	KEYWORDS_CASE     // case
	KEYWORDS_DEFAULT  // default
	KEYWORDS_TRUE     // true
	KEYWORDS_FALSE    // false
	KEYWORDS_TRY      // try
	KEYWORDS_CATCH    // catch
	KEYWORDS_IMPORT   // import

)

func ToKeywords(value string) TokenType {
	keys := map[string]TokenType{
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
	if val, ok := keys[value]; ok {
		return val
	}
	return IDENTIFIER
}
