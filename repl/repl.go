package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"taulang/io"
	"taulang/lexer"
)

func StartREPL(logger *log.Logger) {
	fmt.Println("Welcome to TauLang REPL!")
	fmt.Println("Type 'exit' to quit.")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				fmt.Println("Exiting REPL. Goodbye!")
				break
			}
			ExecuteInput(input, logger)
		}
	}
	if err := scanner.Err(); err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}
}

func ExecuteInput(input string, logger *log.Logger) {
	// Lexical Analysis
	l, err := lexer.NewLexer(input)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

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

	fmt.Println(input)
}
