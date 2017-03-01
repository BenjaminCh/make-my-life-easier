package main

import (
	"fmt"
	"flag"
	"crypto/rand"
	"encoding/hex"
)

func main() {

	// Getting command line params
	var keylength = flag.Int("keylength", 64, "Key length, could be 16, 32, 64 ... 1024 bytes. Default value is 128.")
	flag.Parse()

	key := make([]byte, *keylength)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error creating the key : ", err)
		return
	}
	keyHexa, _ := hex.DecodeString(fmt.Sprintf("%x", key))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("*** Generated key ***")
	fmt.Println("")
	fmt.Println("Key (hex format) : ", fmt.Sprintf("%x", key))
	fmt.Println("Key (byte array format) : ", []byte(keyHexa))
	fmt.Println("")
	fmt.Println("*********************")
}
