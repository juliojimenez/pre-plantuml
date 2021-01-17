package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	args := os.Args[1:]
	for _, file := range args {
		fmt.Println(file)
	}
	content, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File contents: ", string(content))
}
