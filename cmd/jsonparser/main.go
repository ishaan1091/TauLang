package main

import (
	"fmt"
	"jsonparser/cmd/jsonparser/io"
	"jsonparser/cmd/jsonparser/lexer"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	filepath := io.ReadArgs()
	content, err := io.GetContent(filepath)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	tokens, err := lexer.Tokenize(content)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	fmt.Println(content)

	for _, t := range tokens {
		fmt.Printf("%+v\n", *t)
	}
}
