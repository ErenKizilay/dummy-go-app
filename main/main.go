package main

import (
	"fmt"
	"github.com/ErenKizilay/dummy-go-app/util"
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
