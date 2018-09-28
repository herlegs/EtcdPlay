package logutil

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return ""
	}
	return string(b)
}

func WriteFile(file, content string) {
	err := ioutil.WriteFile(file, []byte(content), 0644)
	if err != nil {
		fmt.Printf("error writing file: %v\n", err)
	}
}
