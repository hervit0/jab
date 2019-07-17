package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

type parameter struct {
	code   byte
	length uint8
	value  []byte
}

// https://golang.org/pkg/bytes/
func main() {
	ss := "AQYBEQcCYAAJAQoCAQMECIOQUWRiGUAPCgcDE3eGSBAX/QcDkHeGSBAX/gIBAMAIBgMQd4ZIEBc5CPTRwMD+wP3A"
	fmt.Println("ss:", ss)

	decoded, err := base64.StdEncoding.DecodeString(ss)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println("Array of bytes: ", decoded)
	fmt.Println("Size:", len(decoded))

	var parameters []parameter
	var valueCounter uint8
	var parameterType string
	var currentParameter parameter

	parameterType = "code"
	for _, b := range decoded {
		switch parameterType {
		case "code":
			currentParameter.code = uint8(b)
			parameterType = "length"
		case "length":
			currentParameter.length = uint8(b)
			valueCounter = currentParameter.length
			parameterType = "value"
		case "value":
			if valueCounter == 1 {
				currentParameter.value = append(currentParameter.value, b)
				parameters = append(parameters, currentParameter)

				valueCounter--
				currentParameter = parameter{}
				parameterType = "code"
			} else {
				currentParameter.value = append(currentParameter.value, b)
				valueCounter--
			}
		}
	}

	fmt.Println(parameters)

	aByte := byte(0xc0)
	asInt := uint8(aByte)
	fmt.Println(asInt)
	fmt.Println(byte(0xc0) == byte(0xfd))

	b := []byte{253}
	// b := []byte{0xc0}
	// b := []byte{0x0a, 0xc0, 0xfd}
	fmt.Println(b)

	// https://golang.org/pkg/go/types/
	var pi uint8
	buf := bytes.NewReader(b)
	errRead := binary.Read(buf, binary.LittleEndian, &pi)
	if errRead != nil {
		fmt.Println("binary.Read failed:", errRead)
	}
	fmt.Print(pi)
}
