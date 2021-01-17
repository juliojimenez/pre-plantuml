package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"encoding/hex"
	"strings"
)

func findFiles(re string) []string {
	libRegEx, e := regexp.Compile(re)
	if e != nil {
		log.Fatal(e)
	}
	
	var files []string
	e = filepath.Walk(".", func(filePath string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			files = append(files, filePath)
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
	return files
}

func readFileContentString(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}	
	return string(content)
}

func readFileContentBytes(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}	
	return content
}

func hexEncodedURL(content []byte) string {
	encodedStr := hex.EncodeToString(content)
	return fmt.Sprintf("http://www.plantuml.com/plantuml/png/~h%s", encodedStr)
}

func replaceLineInFile(filePath string, searchString string, replaceString string) {
	libRegEx, e := regexp.Compile(searchString)
	if e != nil {
		log.Fatal(e)
	}
	content := readFileContentString(filePath)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if libRegEx.MatchString(line) {
			libRegEx.ReplaceAllString(line, replaceString)
			lines[i] = replaceString
		}
	}
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	diagramFiles := findFiles(".*\\.(pu)")
	for _, diagramFile := range diagramFiles {
		diagramContent := readFileContentBytes(diagramFile)
		url := hexEncodedURL(diagramContent)
		fmt.Println(url)
		markdownFiles := findFiles(".*\\.md")
		for _, markdownFile := range markdownFiles {
			searchImageStub := fmt.Sprintf("\\!\\[%s\\]\\(.*\\)", diagramFile)
			fmt.Println("Search Image Stub: ", searchImageStub)
			replaceImageStub := fmt.Sprintf("![%s](%s)", diagramFile, url)
			fmt.Println("Replace Image Stub: ", replaceImageStub)
			replaceLineInFile(markdownFile, searchImageStub, replaceImageStub)
		}
	}

}
