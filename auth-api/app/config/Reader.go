package config

import (
	"fmt"
	"io/ioutil"
)

func read(file string) []byte {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return content
}
