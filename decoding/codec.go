package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/ugorji/go/codec"
)

type A struct {
	I int
	S string
}
type B float64

var v1 A
var v2 *A = &v1
var v3 int = 9
var v4 bool = false
var v5 interface{} = v3
var v6 interface{} = nil
var v7 B
var v8 *B = &v7

func main() {
	ss := "AQYBEQcCYAAJAQoCAQMECIOQUWRiGUAPCgcDE3eGSBAX/QcDkHeGSBAX/gIBAMAIBgMQd4ZIEBc5CPTRwMD+wP3A"
	fmt.Println("ss:", ss)

	decoded, _ := base64.StdEncoding.DecodeString(ss)
	fmt.Println("Array of bytes: ", decoded)
	// fmt.Println("Size:", len(decoded))

	// var b []byte = make([]byte, 0, 64)
	// fmt.Println(b)

	// var h codec.Handle = new(codec.JsonHandle)
	// var enc *codec.Encoder = codec.NewEncoderBytes(&b, h)
	// enc.Encode(v1)
	// fmt.Println(b)

	value := []byte{3, 19, 119, 134, 72, 16, 23}
	var h codec.Handle = new(codec.JsonHandle)
	var dec *codec.Decoder = codec.NewDecoderBytes(value, h)
	dec.Decode(v2)
	// var h codec.Handle = new(codec.JsonHandle)
	// var dec *codec.Decoder = codec.NewDecoderBytes(cc, h)
	// dec.Decode(v2)
	// fmt.Println(b)

	fmt.Println("\nTest:")
	test()

	fmt.Println("\nLol:")
	lol()
}

func test() {
	var m = map[string]*A{"1": &A{I: 123, S: "oe"}, "2": &A{I: 2, S: "two"}}
	fmt.Printf("before: %v\n", m)
	var b = []byte(`{"1": {"I":111}, "3": {"I": 333, "S": "blah"} }`)
	var err error = codec.NewDecoderBytes(b, new(codec.JsonHandle)).Decode(&m)
	fmt.Println(err)
	fmt.Printf(" after: %v\n", m)
	for k, v := range m {
		fmt.Printf("\t%v: %v\n", k, v)
	}
}

func lol() {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	r := bytes.NewReader(b)

	var data struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(b)
	fmt.Println(data.PI)
	fmt.Println(data.Uate)
	fmt.Printf("% x\n", data.Mine)
	fmt.Println(data.Too)

	x := make(map[string]int)
	x["lol"] = 12
	v, ok := x["fweafewa"]
	fmt.Println(v)
	fmt.Println(ok)
}
