package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, file := range args {
		fmt.Println(file)
	}
}
