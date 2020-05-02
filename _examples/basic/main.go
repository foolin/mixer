package main

import (
	"fmt"
	"github.com/foolin/mixer"
)

func main() {

	//password
	salt := "a1b2c3d4"

	//the source to be encrypted
	sources := []string{
		"123456",
		"12345abcedf",
		"48656c6c6f204d69786572",
	}

	//NewHex(salt string, upper bool) create a new Mixer
	mix := mixer.NewHex(salt, false)

	for _, source := range sources {

		//Encode source data
		encodeData := mix.EncodeString(source)

		//Decode source data
		decodeData := mix.DecodeString(encodeData)

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n", source, encodeData, decodeData)
	}
}
