package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	// "log"

	// "github.com/4kills/go-libdeflate"
)

func deflateCompress(content []byte) []byte {
	var comp bytes.Buffer
	w, _ := flate.NewWriter(&comp, flate.HuffmanOnly)
	w.Write(content)
	w.Flush()
	w.Close()
	return comp.Bytes()
}

func plantUMLBase64(input []byte) []byte {
	encoding := base64.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_")
	return []byte(encoding.EncodeToString(input))
}
