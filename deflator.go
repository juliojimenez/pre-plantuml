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
	flateWrite.Close()
	return dataBuffer.Bytes()
}

func mapByteToBase64() []byte {
	base64Map := make([]rune, 64)
	i := 0
	for rLower := 'a'; rLower <= 'z'; rLower++ {
		base64Map = append(base64Map, rLower)
		i++
	}
	for rUpper := 'A'; rUpper <= 'Z'; rUpper++ {
		base64Map = append(base64Map, rUpper)
		i++
	}
	base64Map = append(base64Map, '+')
	base64Map = append(base64Map, '/')
	return []byte(string(base64Map))
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
		ip++
		i0 := in[ip] & 0xff
		ip++
		var i1 byte
		if ip < iEnd {
			i1 = in[ip] & 0xff
		} else {
			i1 = 0
		}
		ip++
		var i2 byte
		if ip < iEnd {
			i2 = in[ip] & 0xff
		} else {
			i2 = 0
		}
		o0 := i0 >> 2
		o1 := ((i0 & 3) << 4) | (i1 >> 4)
		o2 := ((i1 & 0xf) << 2) | (i2 >> 6)
		o3 := i2 & 0x3f
		map1 := mapByteToBase64()
		op++
		out[op] = map1[o0]
		op++
		out[op] = map1[o1]
		if op < oDataLen {
			out[op] = map1[o2]
		} else {
			out[op] = byte('=')
		}
		op++
		if op < oDataLen {
			out[op] = map1[o3]
		} else {
			out[op] = byte('=')
		}
		op++
	}
	return out
}