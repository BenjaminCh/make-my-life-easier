package main

import (
	"encoding/hex"
	"flag"
	"fmt"
)

func convertHexaToBytes(key string) ([]byte, error) {
	var err error
	var k []byte

	k, err = hex.DecodeString(key)
	
	if err != nil {
		return nil, err
	}

	return k, nil
}

func main() {

	// Getting command line params
	var hexToConvert = flag.String("hexToConvert", "", "Hexa string to convert to bytes.")
	flag.Parse()

    key, err := convertHexaToBytes(*hexToConvert)

    fmt.Println("")
	fmt.Println("")

    if err != nil {
        fmt.Println("*** Error converting hexa ***")
        fmt.Println("")
        fmt.Println(err)
    } else {
        fmt.Println("*** Converted hexa ***")
        fmt.Println("")
        fmt.Println("Key (hex format) : ", fmt.Sprintf("%x", key))
        fmt.Println("Key (byte array format) : ", key)
    }

	fmt.Println("")
	fmt.Println("*********************")
}