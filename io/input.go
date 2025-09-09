package io

import (
	"fmt"
	"os"
)

func ReadArgs() string {
	// Read args from command line
	args := os.Args[1:]

	// Read filepath if provided
	var filepath string

	if len(args) > 0 {
		filepath = args[0]
	}

	return filepath
}

func GetContentFromFilepath(filepath string) (string, error) {
	// Read from filepath if given
	if filepath != "" {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return "", fmt.Errorf("failed to read file : %v, %v", filepath, err)
		}

		return string(content), err
	}

	return "", nil
}
