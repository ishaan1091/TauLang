package main

import (
	"fmt"
	"jsonparser/cmd/jsonparser/io"
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

	fmt.Println(content)
}
