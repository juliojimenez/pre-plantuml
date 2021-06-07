package main

import (
	"bytes"
	"compress/flate"
)

func deflateCompress(content []byte) []byte {
	var comp bytes.Buffer
	w, _ := flate.NewWriter(&comp, flate.HuffmanOnly)
	w.Write(content)
	w.Flush()
	w.Close()
	return comp.Bytes()
}
