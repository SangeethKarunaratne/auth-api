package config

import (
	"fmt"
	"os"
)

func read(file string) []byte {

	content, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return content
}
