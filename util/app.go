package util

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf(string(content))
	}
}
