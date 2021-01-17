package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

func findFiles(re string) []string {
	libRegEx, e := regexp.Compile(re)
	if e != nil {
		log.Fatal(e)
	}
	
	var files []string
	e = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			files = append(files, info.Name())
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
	return files
}

func main() {
	content, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File contents: ", string(content))
	fmt.Println(findFiles(".*\\.(pu)"))
}
