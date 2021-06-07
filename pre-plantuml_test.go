package main

import (
	"bytes"
	"fmt"
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

func TestReadFileContentBytes(t *testing.T) {
	var mockFS fileSystem = mockFS{}
	result := readFileContentBytes(mockFS, "example.pu")
	if bytes.Compare(result, []byte(`Test String`)) != 0 {
		t.Fatal()
	}
}

func TestHexEncodedURL(t *testing.T) {
	result := hexEncodedURL([]byte(`Hello World!`))
	if result != "http://www.plantuml.com/plantuml/png/~h48656c6c6f20576f726c6421" {
		t.Fatal()
	}
}

func TestReplaceLineInFile(t *testing.T) {
	var mockFS fileSystem = mockFS{}
	result := replaceLineInFile(mockFS, "README.md", "\\!\\[example.pu\\]\\(.*\\)", "![example.pu](https://example.com)")
	if result != true {
		t.Fatal()
	}
}

func TestDeflateEncodedURL(t *testing.T) {
	result := deflateEncodedURL([]byte("@startuml\nAlice -> Bob: Authentication Request\nBob --> Alice: Authentication Response\nAlice -> Bob: Another authentication Request\nAlice <-- Bob: Another authentication Response\n@enduml"))
	fmt.Printf(result)
	if result != "http://www.plantuml.com/plantuml/png/1C3HZSCW40JGVwgO1cZ0Ebd19RW3u4PY9RARcA7_lDTIVRJVCvLfdSWdhcW7ojQWotgLXUFcTtCfNT6GyuaohVD0sHfqMQ-oSDnSd_35LAPr8f-ueXqe7XfyKBS6NTQhB1mtlvjBgKphn5_EkA8TA1uQV52t1btMgomSDzSdV36zwF_xFNy0003__m400F__" {
		t.Fatal()
	}
}

func TestBrotliEncodedURL(t *testing.T) {
	var fs fileSystem = osFS{}
	result := brotliEncodedURL(readFileContentBytes(fs, "diagrams/src/test.pu"))
	fmt.Printf(" %s ", result)
}
