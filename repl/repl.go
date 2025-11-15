package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"taulang/evaluator"
	"taulang/io"
	"taulang/lexer"
	"taulang/object"
	"taulang/parser"
)

func StartREPL(logger *log.Logger) {
	logger.Println("Welcome to TauLang REPL!")
	logger.Println("Type 'exit' to quit.")
	logger.Println("")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				logger.Println("Exiting REPL. Goodbye!")
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
	p := parser.NewParser(l)

	program := p.Parse()
	errors := p.Errors()

	if errors != nil && len(errors) != 0 {
		logger.Println("encountered errors while parsing: ")
		for _, e := range errors {
			logger.Println(e)
		}
		logger.Print("\n\n")
	}

	env := object.NewEnvironment()
	output := evaluator.Eval(program, env)

	logger.Println(output.Inspect())
}
