package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"flag"
	"time"
	"strconv"
	"strings"
)

func applyScaleFactor(price float64, scaleFactor float64, isDebugMode bool) [8]byte {
	scaledPrice := [8]byte{}
	binary.BigEndian.PutUint64(scaledPrice[:], uint64(price*scaleFactor))

	if isDebugMode == true {
		fmt.Println(fmt.Sprintf("Micro price bytes: %v", scaledPrice))
	}

	return scaledPrice
}

func createHmac(key string) (hash.Hash, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return hmac.New(sha1.New, k), nil
}

func hmacSum(hmac hash.Hash, buf []byte) []byte {
	hmac.Reset()
	hmac.Write(buf)
	return hmac.Sum(nil)
}

func Encrypt(encryptionKey, integrityKey, seed string, price float64, scaleFactor float64, isDebugMode bool) string {
	encodingFun, _ := createHmac(encryptionKey)
	integrityFun, _ := createHmac(integrityKey)

	// Result
	var (
		iv        [16]byte
		encoded   [8]byte
		signature [4]byte
	)

	if isDebugMode == true {
		fmt.Println("Encryption key : ", encryptionKey)
		encryptionKeyHexa, _ := hex.DecodeString(encryptionKey)
		fmt.Println("Encryption key (bytes) : ", []byte(encryptionKeyHexa))
		fmt.Println("Integrity key : ", integrityKey)
		integrityKeyHexa, _ := hex.DecodeString(integrityKey)
		fmt.Println("Integrity key (bytes) : ", []byte(integrityKeyHexa))
	}

	data := applyScaleFactor(price, scaleFactor, isDebugMode)

	// Create Initialization Vector from seed
	sum := md5.Sum([]byte(seed))
	copy(iv[:], sum[:])
	if isDebugMode == true {
		fmt.Println("Seed : ", seed)
		fmt.Println("Initialization vector : ", iv)
	}

	//pad = hmac(e_key, iv), first 8 bytes
	pad := hmacSum(encodingFun, iv[:])[:8]
	if isDebugMode == true {
		fmt.Println("// pad = hmac(e_key, iv), first 8 bytes")
		fmt.Println("Pad : ", pad)
	}

	// enc_data = pad <xor> data
	for i := range data {
		encoded[i] = pad[i] ^ data[i]
	}
	if isDebugMode == true {
		fmt.Println("// enc_data = pad <xor> data")
		fmt.Println("Encoded price bytes : ", encoded)
	}

	// signature = hmac(i_key, data || iv), first 4 bytes
	sig := hmacSum(integrityFun, append(data[:], iv[:]...))[:4]
	copy(signature[:], sig[:])
	if isDebugMode == true {
		fmt.Println("// signature = hmac(i_key, data || iv), first 4 bytes")
		fmt.Println("Signature : ", sig)
	}

	// final_message = WebSafeBase64Encode( iv || enc_price || signature )
	return base64.URLEncoding.EncodeToString(append(append(iv[:], encoded[:]...), signature[:]...))
}

func Decrypt(encryptionKey, integrityKey, encodedPrice string, scaleFactor float64) float64 {
	encodingFun, _ := createHmac(encryptionKey)
	integrityFun, _ := createHmac(integrityKey)

	// Decode base64
	decoded, _ := base64.URLEncoding.DecodeString(encodedPrice)

	// Get elements
	var (
		iv         [16]byte
		p          [8]byte
		signature  [4]byte
		priceMicro [8]byte
	)
	copy(iv[:], decoded[0:16])
	copy(p[:], decoded[16:24])
	copy(signature[:], decoded[24:28])

	// pad = hmac(e_key, iv)
	pad := hmacSum(encodingFun, iv[:])[:8]

	// priceMicro = p <xor> pad
	for i := range p {
		priceMicro[i] = pad[i] ^ p[i]
	}

	// conf_sig = hmac(i_key, data || iv)
	sig := hmacSum(integrityFun, append(priceMicro[:], iv[:]...))[:4]

	// success = (conf_sig == sig)
	for i := range sig {
		if sig[i] != signature[i] {
			panic("Failed to decrypt")
		}
	}
	price := float64(binary.BigEndian.Uint64(priceMicro[:])) / scaleFactor
	return price
}

func main() {

	// Getting command line params
	var encryptionKey = flag.String("encryptionkey", "", "Encryption key")
	var integrityKey = flag.String("integritykey", "", "Integrity key")
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

	var encryptedPrice string 
	for _, priceToTest := range pricesToTest {
		
		fmt.Println(fmt.Sprintf("\nInitial price: %s", priceToTest))

		if *mode == "all" || *mode == "encrypt" {
			price, err := strconv.ParseFloat(priceToTest, 64)
			if err != nil {
				fmt.Println("Error trying to parse price to encrypt, cannot convert %s to float.", priceToTest)
				return
			}
			encryptedPrice = Encrypt(*encryptionKey, *integrityKey, *initializationVector, price, *scaleFactor, *debug)
			fmt.Println("Encrypted price:", encryptedPrice)
		}

		if *mode == "all" || *mode == "decrypt" {
			if encryptedPrice == "" && *mode == "decrypt" {
				encryptedPrice = priceToTest[:]
			}
			decryptedPrice := Decrypt(*encryptionKey, *integrityKey, encryptedPrice, *scaleFactor)
			fmt.Println("Decrypted price:", decryptedPrice)
		}
	}
}
