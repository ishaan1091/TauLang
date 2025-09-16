package token

type TokenType string

const (
	// identifiers + literals
	NUMBER     TokenType = "NUMBER"
	STRING     TokenType = "STRING"
	IDENTIFIER TokenType = "IDENTIFIER"

	// operators
	ASSIGNMENT     TokenType = "ASSIGNMENT"     // ne_bana_diye
	BANG           TokenType = "BANG"           // !
	EQUALS         TokenType = "EQUALS"         // ==
	NOT_EQUALS     TokenType = "NOT_EQUALS"     // !=
	GREATER_THAN   TokenType = "GREATER_THAN"   // >
	LESSER_THAN    TokenType = "LESSER_THAN"    // <
	GREATER_EQUALS TokenType = "GREATER_EQUALS" // >=
	LESSER_EQUALS  TokenType = "LESSER_EQUALS"  // <=
	ADDITION       TokenType = "ADDITION"       // +
	SUBTRACTION    TokenType = "SUBTRACTION"    // -
	MULTIPLICATION TokenType = "MULTIPLICATION" // *
	DIVISION       TokenType = "DIVISION"       // /

	// keywords
	LET      TokenType = "LET"
	FUNCTION TokenType = "FUNCTION"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	WHILE    TokenType = "WHILE"
	BREAK    TokenType = "BREAK"
	CONTINUE TokenType = "CONTINUE"

	// delimiters
	COMMA     TokenType = "COMMA"     // ,
	SEMICOLON TokenType = "SEMICOLON" // ;
	COLON     TokenType = "COLON"     // :
	QUOTES    TokenType = "QUOTES"    // "

	// parenthesis
	LEFT_BRACE    TokenType = "LEFT_BRACE"    // {
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"   // }
	LEFT_BRACKET  TokenType = "LEFT_BRACKET"  // [
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET" // ]
	LEFT_PAREN    TokenType = "LEFT_PAREN"    // (
	RIGHT_PAREN   TokenType = "RIGHT_PAREN"   // )

	// special tokens
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}

var Keywords = map[string]TokenType{
	"sun_liyo_tau":         LET,
	"rasoi_mein_bata_diye": FUNCTION,
	"agar_maan_lo":         IF,
	"na_toh":               ELSE,
	"laadle_ye_le":         RETURN,
	"saccha":               TRUE,
	"jhootha":              FALSE,
	"jab_tak":              WHILE,
	"rok_diye":             BREAK,
	"jaan_de":              CONTINUE,
	"ne_bana_diye":         ASSIGNMENT,
}

var ReverseKeywords = map[TokenType]string{
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

func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}
