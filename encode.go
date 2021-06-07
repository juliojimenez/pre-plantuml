package main

import (
	"encoding/base64"
)

func plantUMLBase64(input []byte) []byte {
	encoding := base64.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_")
	return []byte(encoding.EncodeToString(input))
}
