package main

import (
	"fmt"
	"github.com/foolin/mixer"
)

func main() {

	//the source to be encrypted
	sources := []string{
		"123456",
		"12345abcedf",
		"0123456789abcdefghijklmnopqrstuvwxyz",
	}

	//password
	password := "1q2w3e"

	//foreach every source
	for _, source := range sources {
		//Encode source data
		encodeData := mixer.EncodeString(password, source)

		//Decode source data
		decodeData := mixer.DecodeString(password, encodeData)

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n",
			source, encodeData, decodeData)
	}

}
