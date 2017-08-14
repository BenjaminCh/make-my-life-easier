package main

import (
	"regexp"
	"flag"
	"fmt"
)

func convertToBase64WebSafe(base64 string) (string) {

	var base64WebSafe string

    base64WebSafe = base64
    base64WebSafe = regexp.MustCompile("\\+").ReplaceAllLiteralString(base64WebSafe, "-")
    base64WebSafe = regexp.MustCompile("\\/").ReplaceAllLiteralString(base64WebSafe, "_")
    base64WebSafe = regexp.MustCompile("=+$").ReplaceAllLiteralString(base64WebSafe, "")

	return base64WebSafe
}

func convertToBase64(base64WebSafe string) (string) {

	var base64 string

    base64 = base64WebSafe
    base64 = regexp.MustCompile("-").ReplaceAllLiteralString(base64, "+")
    base64 = regexp.MustCompile("_").ReplaceAllLiteralString(base64, "/")
    
    base64 = (base64 + "==")[0:(3*len(base64)%4)];

	return base64
}

func main() {

	// Getting command line params
    var base64ToEdit = flag.String("value", "", "Base 64 to edit.")
    var action = flag.String("action", "base64websafe", "Action to perform: base64websafe, base64 or isbase64websafe. Default : base64websafe")
    flag.Parse()

    if !(*action == "base64websafe" || *action == "base64" || *action == "isbase64websafe") {
		fmt.Println("Action should be either : 'base64websafe', 'base64' or 'isbase64websafe'")
		return
	}

    var base64, base64WebSafe string

    if *action == "base64websafe" || *action == "isbase64websafe" {
        base64 = *base64ToEdit
        base64WebSafe = convertToBase64WebSafe(*base64ToEdit)
    } else if *action == "base64" {
        base64WebSafe = *base64ToEdit
        base64 = convertToBase64(*base64ToEdit)
    }

    fmt.Println("")
    fmt.Println("******************************************")
	fmt.Println("")
    if *action == "base64websafe" || *action == "base64" {
        fmt.Println("***** Converted " + *action + " *****")
        fmt.Println("")
        fmt.Println("Base64 : ", base64)
        fmt.Println("Base64 Web Safe : ", base64WebSafe)
    } else if *action == "isbase64websafe" {
        fmt.Println("***** Base64 Url Safe Check *****")
        fmt.Println("")
        if base64 == base64WebSafe {
            fmt.Println("Your base64 is web safe :)")
        } else {
            fmt.Println("Your base64 is NOT web safe :(")
        }
    }
    fmt.Println("")
	fmt.Println("******************************************")
}