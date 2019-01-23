package main

import (
	"dummy-app/util"
	"fmt"
	"os"
)

func main() {
	filePath := os.Getenv("DUMMY_PATH")
	if filePath != "" {
		util.ReadFile(filePath)
	} else {
		fmt.Printf("no file found to read!!!")
	}
}
