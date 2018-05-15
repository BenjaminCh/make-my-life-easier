package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/benjaminch/openrtb-pricers/doubleclick"
	pricerhelper "github.com/benjaminch/openrtb-pricers/helpers"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Commands / Params names
var (
	// Prices encryption / decryption
	priceEncryptionCmdValue  = "price-encryption"
	priceDecryptionCmdValue  = "price-decryption"
	algorithmNameArgValue    = "algorithm"
	encryptionKeyArgValue    = "ekey"
	integrityKeyArgValue     = "ikey"
	keysTypeArgValue         = "keystype"
	keysAreBase64ArgValue    = "keysb64"
	priceScaleFactorArgValue = "scale"
	seedArgValue             = "seed"
	priceArgValue            = "price"

	// Hex key generator
	hexKeyGeneratorCmdValue = "hex-key-generator"
	keyLengthArgValue       = "keylength"
)

// Commands and Flags
var (
	// Common flags
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()

	// Commands

	// Encryption
	priceEncryptionCmd                 = kingpin.Command(priceEncryptionCmdValue, "Encrypt a price using a given algorithm.")
	priceEncryptionEncryptionKeyArg    = priceEncryptionCmd.Arg(encryptionKeyArgValue, "Encryption key").Required().String()
	priceEncryptionIntegrityKeyArg     = priceEncryptionCmd.Arg(integrityKeyArgValue, "Integrity key").Required().String()
	priceEncryptionPriceArg            = priceEncryptionCmd.Arg(priceArgValue, "Price to encrypt").Required().Float()
	priceEncryptionAlgorithmNameArg    = priceEncryptionCmd.Flag(algorithmNameArgValue, "Name of the algorithm").Default("google").Enum("google")
	priceEncryptionKeysTypeArg         = priceEncryptionCmd.Flag(keysTypeArgValue, "Keys type").Default("utf-8").Enum("utf-8", "hexa")
	priceEncryptionKeysAreBase64Arg    = priceEncryptionCmd.Flag(keysAreBase64ArgValue, "Keys are base 64.").Default("false").Bool()
	priceEncryptionPriceScaleFactorArg = priceEncryptionCmd.Flag(priceScaleFactorArgValue, "Price scale factor").Default("1000000").Float()
	priceEncryptionSeedArg             = priceEncryptionCmd.Flag(seedArgValue, "Seed").Default("").String()

	// Decryption
	priceDecryptionCmd                 = kingpin.Command(priceDecryptionCmdValue, "Decrypt a price using a given algorithm.")
	priceDecryptionEncryptionKeyArg    = priceDecryptionCmd.Arg(encryptionKeyArgValue, "Encryption key").Required().String()
	priceDecryptionIntegrityKeyArg     = priceDecryptionCmd.Arg(integrityKeyArgValue, "Integrity key").Required().String()
	priceDecryptionPriceArg            = priceDecryptionCmd.Arg(priceArgValue, "Price to decrypt").Required().String()
	priceDecryptionAlgorithmNameArg    = priceDecryptionCmd.Flag(algorithmNameArgValue, "Name of the algorithm").Default("google").Enum("google")
	priceDecryptionKeysTypeArg         = priceDecryptionCmd.Flag(keysTypeArgValue, "Keys type").Default("utf-8").Enum("utf-8", "hexa")
	priceDecryptionKeysAreBase64Arg    = priceDecryptionCmd.Flag(keysAreBase64ArgValue, "Keys are base 64.").Default("false").Bool()
	priceDecryptionPriceScaleFactorArg = priceDecryptionCmd.Flag(priceScaleFactorArgValue, "Price scale factor").Default("1000000").Float()

	// Hex key generator
	hexKeyGeneratorCmd          = kingpin.Command(hexKeyGeneratorCmdValue, "Generate a random hex key of the specified length (by default: 64)")
	hexKeyGeneratorKeyLengthArg = hexKeyGeneratorCmd.Arg(keyLengthArgValue, "Key Length").Default(32).Int64()
)

func main() {
	var cmd = kingpin.Parse()

	if *verbose {
		fmt.Println("Verbose activated")
	}

	switch cmd {

	// Price Encryption
	// TODO: Refactor this code, put it in helpers
	case priceEncryptionCmdValue:
		var keysType pricerhelper.KeyDecodingMode
		var pricer *doubleclick.DoubleClickPricer
		var err error

		keysType, err = pricerhelper.ParseKeyDecodingMode(*priceEncryptionKeysTypeArg)
		if err != nil {
			fmt.Println("Error while trying to parse keys type: ", err)
			return
		}

		pricer, err = doubleclick.NewDoubleClickPricer(
			*priceEncryptionEncryptionKeyArg,
			*priceEncryptionIntegrityKeyArg,
			*priceEncryptionKeysAreBase64Arg,
			keysType,
			*priceEncryptionPriceScaleFactorArg,
			*verbose,
		)
		if err != nil {
			fmt.Println("Error while creating pricer: ", err)
			return
		}

		var encryptedPrice string
		encryptedPrice, err = pricer.Encrypt(*priceEncryptionSeedArg, *priceEncryptionPriceArg, *verbose)
		if err != nil {
			fmt.Println("Error while trying to encrypt the price: ", err)
			return
		}
		fmt.Println("Encrypted price: ", encryptedPrice)
		break

	// Price Decryption
	// TODO: Refactor this code, put it in helpers
	case priceDecryptionCmdValue:
		var keysType pricerhelper.KeyDecodingMode
		var pricer *doubleclick.DoubleClickPricer
		var err error

		keysType, err = pricerhelper.ParseKeyDecodingMode(*priceDecryptionKeysTypeArg)
		if err != nil {
			fmt.Println("Error while trying to parse keys type: ", err)
			return
		}

		pricer, err = doubleclick.NewDoubleClickPricer(
			*priceDecryptionEncryptionKeyArg,
			*priceDecryptionIntegrityKeyArg,
			*priceDecryptionKeysAreBase64Arg,
			keysType,
			*priceDecryptionPriceScaleFactorArg,
			*verbose,
		)
		if err != nil {
			fmt.Println("Error while creating pricer: ", err)
			return
		}

		var decryptedPrice float64
		decryptedPrice, err = pricer.Decrypt(*priceDecryptionPriceArg, *verbose)
		if err != nil {
			fmt.Println("Error while trying to decrypt the price: ", err)
			return
		}
		fmt.Println("Decrypted price: ", decryptedPrice)
		break

	// Hex key generator
	// TODO: Refactor this code, put it in helpers
	case hexKeyGeneratorCmdValue:
		key := make([]byte, *hexKeyGeneratorKeyLengthArg)
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
		break

	default:
		fmt.Println("No command specified")
	}
}
