package main

import (
	"log"
	"os"
	"taulang/io"
	"taulang/repl"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	filepath := io.ReadArgs()
	content, err := io.GetContentFromFilepath(filepath)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	if content != "" {
		repl.ExecuteInput(content, logger)
	} else {
		repl.StartREPL(logger)
	}
}
