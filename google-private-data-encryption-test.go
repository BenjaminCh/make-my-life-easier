package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	pricers "github.com/benjaminch/openrtb-pricers"
)

func main() {

	// Getting command line params
	var encryptionKey = flag.String("encryptionkey", "", "Encryption key")
	var integrityKey = flag.String("integritykey", "", "Integrity key")
	var keyDecodingMode = flag.String("keyDecodingMode", "hexa", "Key decoding mode : hexa or utf-8, default is hexa")
	var initializationVector = flag.String("seed", fmt.Sprintf("%d", time.Now()), "Seed for initialization vector, default is current timestamp")
	var priceToEncrypt = flag.String("price", "0", "Price to encrypt, default = 0.0")
	var scaleFactor = flag.Float64("scalefactor", 1000000, "What scale factor to apply on the price for encryption? Default is micros (1000000)")
	var debug = flag.Bool("debug", false, "Debug traces for middle steps, default = false")
	var mode = flag.String("mode", "all", "Specify what to do: 'all' : encrypt / decrypt, 'encrypt' : encrypt only, 'decrypt' : decrypt only")
	flag.Parse()

	if *encryptionKey == "" {
		fmt.Println("Encryption Key is mandatory !")
		return
	}
	if *integrityKey == "" {
		fmt.Println("Integrity Key is mandatory !")
		return
	}
	var keyDecoding pricers.KeyDecodingMode
	if !(*keyDecodingMode == "hexa" || *keyDecodingMode == "utf-8") {
		fmt.Println("KeyDecodingMode should be either : 'hexa' or 'utf-8'")
		return
	} else {
		keyDecoding = *keyDecodingMode
	}
	if !(*mode == "all" || *mode == "decrypt" || *mode == "encrypt") {
		fmt.Println("Mode should be either : 'all', 'encrypt' or 'decrypt'")
		return
	}

	var priceToEncryptTrimed string
	var pricesToTest []string

	priceToEncryptTrimed = strings.Replace(*priceToEncrypt, " ", "", -1)
	if strings.Contains(priceToEncryptTrimed, ",") {
		pricesToTest = strings.Split(priceToEncryptTrimed, ",")
	} else {
		pricesToTest = []string{priceToEncryptTrimed}
	}

	// Create the DoubleClick Pricer
	var pricer *Pricer
	var err error
	pricer, err = pricers.NewDoubleClickPricer(
		*encryptionKey,
		*integrityKey,
		false,         // TODO : Handle this case allowing to pass it with cmd line
		keyDecoding,
		*scaleFactor,
		*debug),
	)

	var encryptedPrice string
	for _, priceToTest := range pricesToTest {

		fmt.Println(fmt.Sprintf("\nInitial price: %s", priceToTest))

		if *mode == "all" || *mode == "encrypt" {
			price, err := strconv.ParseFloat(priceToTest, 64)
			if err != nil {
				fmt.Println("Error trying to parse price to encrypt, cannot convert %s to float.", priceToTest)
				return
			}
			encryptedPrice, err = pricer.Encrypt(*initializationVector, price, *debug)
			if err != nil {
				err = errors.New("Encryption failed. Error : %s", err)
				fmt.Println(err)
				return
			}
			fmt.Println("Encrypted price:", encryptedPrice)
		}

		if *mode == "all" || *mode == "decrypt" {
			fmt.Println("Encrypted price:", encryptedPrice)
			var decryptedPrice float64
			decryptedPrice, err = pricer.Decrypt(
				encryptedPrice
				*debug,
			)
			if err != nil {
				err = errors.New("Decryption failed. Error : %s", err)
				fmt.Println(err)
				return
			}
			fmt.Println("Decrypted price:", decryptedPrice)
		}
	}
}
