package main

import (
	"bytes"
	"github.com/andybalholm/brotli"
)

func brotliCompress(content []byte) []byte {
	var comp bytes.Buffer
	w := brotli.NewWriterLevel(&comp, 0)
	w.Write(content)
	w.Close()
	return comp.Bytes()
}
