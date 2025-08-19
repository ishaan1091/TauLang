package commons

type TokenType string

const (
	STRING  TokenType = "STRING"
	INTEGER TokenType = "INTEGER"
	BOOLEAN TokenType = "BOOLEAN"

	COLON         TokenType = "COLON"
	COMMA         TokenType = "COMMA"
	LEFT_BRACE    TokenType = "LEFT_BRACE"
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"
	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"

	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}
