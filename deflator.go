package main

import (
	"bytes"
	"compress/flate"
	"log"
)

func deflateCompress(content []byte) []byte {
	var dataBuffer bytes.Buffer
	flateWrite, err := flate.NewWriter(&dataBuffer, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}
	_, error := flateWrite.Write(content)
	if error != nil {
		log.Fatal(error)
	}
	return dataBuffer.Bytes()
}