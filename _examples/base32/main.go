package main

import (
	"fmt"
	"github.com/foolin/mixer"
)

func main() {

	sources := []string{
		"!@#$#$%^&*()",
		"/foolin/Mixer#Install",
		"我是Mixer",
	}

	//password
	password := "1q2w3e"

	//foreach every source
	for _, source := range sources {

		//Encode source data
		encodeData := mixer.EncodeBase32(password, source)

		//Decode source data
		decodeData, err := mixer.DecodeBase32(password, encodeData)
		if err != nil {
			panic(err)
		}

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n",
			source, encodeData, decodeData)
	}

}
