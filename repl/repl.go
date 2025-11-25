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
	env := object.NewEnvironment()
	for {
		fmt.Print(">> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				logger.Println("Exiting REPL. Goodbye!")
				break
			}
			executeInputWithEnvironment(input, logger, env)
		}
	}
	if err := scanner.Err(); err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}
}

func ExecuteInput(input string, logger *log.Logger) {
	env := object.NewEnvironment()
	executeInputWithEnvironment(input, logger, env)
}

func executeInputWithEnvironment(input string, logger *log.Logger, env object.Environment) {
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

	output := evaluator.Eval(program, env)

	if output == evaluator.NULL {
		logger.Println("")
	} else {
		logger.Println(output.Inspect())
	}
}
