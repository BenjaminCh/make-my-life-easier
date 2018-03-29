package hexbytesconverter

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
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

func convertBytesToHexa(key []byte) string {
	var k string

	k = hex.EncodeToString(key)

	return k
}

func parseBytesArrayFromInput(bytesArrayString string) ([]byte, error) {
	var err error
	var bytesArray []byte

	bytesArrayString = strings.TrimLeft(bytesArrayString, "[")
	bytesArrayString = strings.TrimRight(bytesArrayString, "]")
	bytesArrayString = strings.Replace(bytesArrayString, " ", "", -1)

	var bytesItems = strings.Split(bytesArrayString, ",")
	var itemInt int64

	for _, item := range bytesItems {
		itemInt, err = strconv.ParseInt(item, 10, 64)

		if err != nil {
			return []byte{}, err
		}

		bytesArray = append(bytesArray, byte(itemInt))
	}

	return bytesArray, err
}

func main() {

	// Getting command line params
	var value = flag.String("value", "", "Value to convert.")
	var action = flag.String("action", "hextobytes", "Action to perform: hextobytes or bytestohex. Default : hextobytes")
	flag.Parse()

	if !(*action == "hextobytes" || *action == "bytestohex") {
		fmt.Println("Action should be either : 'bytestohex', 'hextobytes'")
		return
	}

	var hexa string
	var bytes []byte
	var err error

	if *action == "hextobytes" {
		hexa = *value
		bytes, err = convertHexaToBytes(hexa)
	} else if *action == "bytestohex" {

		bytes, err = parseBytesArrayFromInput(*value)
		if err != nil {
			fmt.Println("Error while parsing byte array. Should be formatted like : [12,15,49]")
			return
		}
		hexa = convertBytesToHexa(bytes)
	}

	fmt.Println("")
	fmt.Println("")

	if err != nil {
		fmt.Println("*** Error converting hexa ***")
		fmt.Println("")
		fmt.Println(err)
	} else {
		fmt.Println("*** Converted ***")
		fmt.Println("")
		fmt.Println("Value (hex format) : ", hexa)
		fmt.Println("Value (byte array format) : ", bytes)
	}

	fmt.Println("")
	fmt.Println("*********************")
}
