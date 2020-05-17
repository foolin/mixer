package main

import (
	"crypto/md5"
	"fmt"
	"github.com/foolin/mixer"
)

func main() {

	//the source to be encrypted
	sources := []string{
		"abc012345edf",
		"0123456789abcdefghijklmnopqrstuvwxyz",
		fmt.Sprintf("%x", md5.Sum([]byte("Hello Mixer"))),
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
