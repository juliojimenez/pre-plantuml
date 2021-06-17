package main

import (
	"bytes"
	"compress/flate"
)

func deflateCompress(content []byte) ([]byte, error) {
	var comp bytes.Buffer
	w, _ := flate.NewWriter(&comp, flate.HuffmanOnly)
	_, writeErr := w.Write(content)
	if writeErr != nil {
		return nil, writeErr
	}
	_ = w.Flush()
	_ = w.Close()
	return comp.Bytes(), nil
}
