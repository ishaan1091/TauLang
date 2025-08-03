package io

import (
	"log"
	"os"
)

func OutputFatalErrorAndExit(logger *log.Logger, err error) {
	logger.Fatal(err)
	os.Exit(1)
}
