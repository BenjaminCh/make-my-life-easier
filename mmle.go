package main

import (
	"fmt"

	"github.com/benjaminch/openrtb-pricers/doubleclick"
	pricerhelper "github.com/benjaminch/openrtb-pricers/helpers"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Commands / Params names
var (
	priceEncryptionCmdValue  = "price-encryption"
	algorithmNameArgValue    = "algorithm"
	encryptionKeyArgValue    = "ekey"
	integrityKeyArgValue     = "ikey"
	keysTypeArgValue         = "keystype"
	keysAreBase64ArgValue    = "keysb64"
	priceScaleFactorArgValue = "scale"
	seedArgValue             = "seed"
	priceArgValue            = "price"
)

// Commands and Flags
var (
	// Common flags
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()

	// Commands

	// Encryption
	priceEncryptionCmd  = kingpin.Command(priceEncryptionCmdValue, "Encrypt a price using a given algorithm.")
	algorithmNameArg    = priceEncryptionCmd.Arg(algorithmNameArgValue, "Name of the algorithm").Required().Enum("google")
	encryptionKeyArg    = priceEncryptionCmd.Arg(encryptionKeyArgValue, "Encryption key").Required().String()
	integrityKeyArg     = priceEncryptionCmd.Arg(integrityKeyArgValue, "Integrity key").Required().String()
	keysTypeArg         = priceEncryptionCmd.Arg(keysTypeArgValue, "Keys type").Default("utf-8").Enum("utf-8", "hexa")
	keysAreBase64Arg    = priceEncryptionCmd.Arg(keysAreBase64ArgValue, "Keys are base 64.").Bool()
	priceScaleFactorArg = priceEncryptionCmd.Arg(priceScaleFactorArgValue, "Price scale factor").Float()
	seedArg             = priceEncryptionCmd.Arg(seedArgValue, "Seed").Default("").String()
	priceArg            = priceEncryptionCmd.Arg(priceArgValue, "Price to encrypt").Float()
)

func main() {
	var cmd = kingpin.Parse()

	if *verbose {
		fmt.Println("Verbose activated")
	}

	switch cmd {
	case priceEncryptionCmdValue:
		var keysType pricerhelper.KeyDecodingMode
		var pricer *doubleclick.DoubleClickPricer
		var err error

		keysType, err = pricerhelper.ParseKeyDecodingMode(*keysTypeArg)
		if err != nil {
			fmt.Println("Error while trying to parse keys type: ", err)
			return
		}

		pricer, err = doubleclick.NewDoubleClickPricer(
			*encryptionKeyArg,
			*integrityKeyArg,
			*keysAreBase64Arg,
			keysType,
			*priceScaleFactorArg,
			*verbose,
		)
		if err != nil {
			fmt.Println("Error while creating pricer: ", err)
			return
		}

		var encryptedPrice string
		encryptedPrice, err = pricer.Encrypt(*seedArg, *priceArg, *verbose)
		if err != nil {
			fmt.Println("Error while trying to encrypt the price: ", err)
			return
		}
		fmt.Println("Encrypted price: ", encryptedPrice)
	default:
		fmt.Println("No command specified")
	}
}
