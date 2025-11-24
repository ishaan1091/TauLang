package token

type Type string

const (
	// identifiers + literals
	NUMBER     Type = "NUMBER"
	STRING     Type = "STRING"
	IDENTIFIER Type = "IDENTIFIER"

	// operators
	ASSIGNMENT     Type = "ASSIGNMENT"     // ne_bana_diye
	BANG           Type = "BANG"           // !
	EQUALS         Type = "EQUALS"         // ==
	NOT_EQUALS     Type = "NOT_EQUALS"     // !=
	GREATER_THAN   Type = "GREATER_THAN"   // >
	LESSER_THAN    Type = "LESSER_THAN"    // <
	GREATER_EQUALS Type = "GREATER_EQUALS" // >=
	LESSER_EQUALS  Type = "LESSER_EQUALS"  // <=
	ADDITION       Type = "ADDITION"       // +
	SUBTRACTION    Type = "SUBTRACTION"    // -
	MULTIPLICATION Type = "MULTIPLICATION" // *
	DIVISION       Type = "DIVISION"       // /

	// keywords
	LET      Type = "LET"
	FUNCTION Type = "FUNCTION"
	IF       Type = "IF"
	ELSE     Type = "ELSE"
	RETURN   Type = "RETURN"
	TRUE     Type = "TRUE"
	FALSE    Type = "FALSE"
	WHILE    Type = "WHILE"
	BREAK    Type = "BREAK"
	CONTINUE Type = "CONTINUE"

	// delimiters
	COMMA     Type = "COMMA"     // ,
	SEMICOLON Type = "SEMICOLON" // ;
	COLON     Type = "COLON"     // :

	// parenthesis
	LEFT_BRACE    Type = "LEFT_BRACE"    // {
	RIGHT_BRACE   Type = "RIGHT_BRACE"   // }
	LEFT_BRACKET  Type = "LEFT_BRACKET"  // [
	RIGHT_BRACKET Type = "RIGHT_BRACKET" // ]
	LEFT_PAREN    Type = "LEFT_PAREN"    // (
	RIGHT_PAREN   Type = "RIGHT_PAREN"   // )

	// special tokens
	EOF     Type = "EOF"
	ILLEGAL Type = "ILLEGAL"
)

type Token struct {
	Type    Type
	Literal string
}

var Keywords = map[string]Type{
	"sun_liyo_tau":  LET,
	"tau_ka_jugaad": FUNCTION,
	"agar_maan_lo":  IF,
	"na_toh":        ELSE,
	"laadle_ye_le":  RETURN,
	"saccha":        TRUE,
	"jhootha":       FALSE,
	"jab_tak":       WHILE,
	"rok_diye":      BREAK,
	"jaan_de":       CONTINUE,
	"ne_bana_diye":  ASSIGNMENT,
}

var ReverseKeywords = map[Type]string{
	LET:        "let",
	FUNCTION:   "func",
	IF:         "if",
	ELSE:       "else",
	RETURN:     "return",
	TRUE:       "true",
	FALSE:      "false",
	WHILE:      "while",
	BREAK:      "break",
	CONTINUE:   "continue",
	ASSIGNMENT: "=",
}

func GetTokenForIdentifierOrKeyword(value string) Token {
	if tok, ok := Keywords[value]; ok {
		return NewToken(tok, ReverseKeywords[tok])
	}
	return NewToken(IDENTIFIER, value)
}

func NewToken(tokenType Type, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}
