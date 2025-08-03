package commons

type DataType string

const (
	STRING  DataType = "STRING"
	INTEGER DataType = "INTEGER"
	BOOLEAN DataType = "BOOLEAN"
)

type Token struct {
	Value string
	Type  DataType
}
