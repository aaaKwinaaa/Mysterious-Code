package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt := Reverse(string(sd))
	fmt.Print("Your secret is : " + whatIsIt)

}

func Reverse(word string) (result string) {
	for _, s := range word {
		result = string(s) + result
	}
	return
}
