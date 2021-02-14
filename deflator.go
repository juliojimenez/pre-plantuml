package main

import (
	"bytes"
	"compress/flate"
	"fmt"
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
	flateWrite.Close()
	return dataBuffer.Bytes()
}

// https://github.com/plantuml/plantuml/blob/master/src/net/sourceforge/plantuml/code/Base64Coder.java
func deflateBase64ishEncode(in []byte, iOff int, iLen int) []byte {
	oDataLen := (iLen * 4 + 2) / 3
	oLen := ((iLen + 2) / 3) * 4
	out := make([]byte, oLen)
	ip := iOff
	iEnd := iOff + iLen
	op := 0
	for ip < iEnd {

	}
	return out
	fmt.Println(oDataLen)
	fmt.Println(oLen)
}