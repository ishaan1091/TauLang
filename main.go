package main

import (
	"fmt"
	"log"
	"os"
	"taulang/io"
	"taulang/lexer"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	filepath := io.ReadArgs()
	content, err := io.GetContent(filepath)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	l := lexer.NewLexer(content)

	for {
		token, err := l.NextToken()
		if err != nil {
			io.OutputFatalErrorAndExit(logger, err)
		}
		if token.Type == "EOF" {
			break
		}
		fmt.Printf("%+v\n", token)
	}

	fmt.Println(content)
}
