package main

import (
	"bytes"
	"github.com/andybalholm/brotli"
)

func brotliCompress(content []byte) ([]byte, error) {
	var comp bytes.Buffer
	w := brotli.NewWriterLevel(&comp, 11)
	_, writeErr := w.Write(content)
	if writeErr != nil {
		return nil, writeErr
	}
	_ = w.Close()
	return comp.Bytes(), nil
}
