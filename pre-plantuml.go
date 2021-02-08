package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func findFiles(fs fileSystem, re string) ([]string, error) {
	libRegEx, e := regexp.Compile(re)
	if e != nil {
		return nil, e
	}

	var files []string
	e = fs.Walk(".", func(filePath string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			files = append(files, filePath)
		}
		return nil
	})
	if e != nil {
		return nil, e
	}
	return files, nil
}

func readFileContentString(fs fileSystem, filePath string) string {
	content, err := fs.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func readFileContentBytes(fs fileSystem, filePath string) []byte {
	content, err := fs.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func hexEncodedURL(content []byte) string {
	encodedStr := hex.EncodeToString(content)
	return fmt.Sprintf("http://www.plantuml.com/plantuml/png/~h%s", encodedStr)
}

func replaceLineInFile(fs fileSystem, filePath string, searchString string, replaceString string) bool {
	libRegEx, e := regexp.Compile(searchString)
	if e != nil {
		log.Fatal(e)
	}
	content := readFileContentString(fs, filePath)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if libRegEx.MatchString(line) {
			libRegEx.ReplaceAllString(line, replaceString)
			lines[i] = replaceString
		}
	}
	output := strings.Join(lines, "\n")
	err := fs.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Updated image tags in ", filePath)
	return true
}

func main() {
	var fs fileSystem = osFS{}
	diagramFiles, e := findFiles(fs, ".*\\.(pu|puml|plantuml|iuml|wsd)")
	if e != nil {
		log.Fatal(e)
	}
	for _, diagramFile := range diagramFiles {
		diagramContent := readFileContentBytes(fs, diagramFile)
		fmt.Println("Hex Encoding Diagram for: ", diagramFile)
		url := hexEncodedURL(diagramContent)
		markdownFiles, e := findFiles(fs, ".*\\.md")
		if e != nil {
			log.Fatal(e)
		}
		for _, markdownFile := range markdownFiles {
			fmt.Println("Working on ", markdownFile)
			searchImageStub := fmt.Sprintf("\\!\\[%s\\]\\(.*\\)", diagramFile)
			replaceImageStub := fmt.Sprintf("![%s](%s)", diagramFile, url)
			replaceLineInFile(fs, markdownFile, searchImageStub, replaceImageStub)
		}
	}
}
