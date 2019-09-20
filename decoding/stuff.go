package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
)

type parameter struct {
	code   byte
	length uint8
	value  []byte
}

type cli struct {
	name                    string
	oddness                 bool
	addressNature           uint
	presentationRestriction uint
	screening               uint
	digits                  uint
}

// https://golang.org/pkg/bytes/
func main2() {
	i, err := strconv.ParseInt("00000011", 7, 64)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println("Array of bytes: ", i)

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
			currentParameter.code = b
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
	fmt.Println(bytes.Contains([]byte{0xc0, 0xfd}, []byte{0xfa}))

	knownClis := []byte{0xfd, 0xc0, 0x0a}
	clisParameters := make([]parameter, 0)
	for _, parameter := range parameters {
		if bytes.Contains(knownClis, []byte{parameter.code}) {
			clisParameters = append(clisParameters, parameter)
		}
	}
	fmt.Println(clisParameters)
	// [{10 7 [3 19 119 134 72 16 23]} {253 7 [3 144 119 134 72 16 23]} {192 8 [6 3 16 119 134 72 16 23]}]

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
