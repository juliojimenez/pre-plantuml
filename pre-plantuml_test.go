package main

import (
	"testing"
)

// TestFindFiles Test
func TestFindFiles(t *testing.T) {
	var mockFs fileSystem = mockFS{}
	_, e := findFiles(mockFs, ".*\\.(pu|puml|plantuml|iuml|wsd)")
	if e != nil {
		t.Fatal(e)
	}
}

func TestReadFileContentString(t *testing.T) {
	var mockFS fileSystem = mockFS{}
	result := readFileContentString(mockFS, "example.pu")
	if result != "Test String" {
		t.Fatal()
	}
}
