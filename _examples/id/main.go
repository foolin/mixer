package main

import (
	"fmt"
	"github.com/foolin/mixer"
)

func main() {

	//the source to be encrypted
	sources := []uint64{
		0,
		123,
		123456,
		1234567890,
		999999999,
		9223372036854775808,
		18446744073709551615,
	}

	//password
	password := "1q2w3e"

	//foreach every source
	for _, source := range sources {

		//Encode source data
		encodeData := mixer.EncodeID(password, source)

		//Decode source data
		decodeData, err := mixer.DecodeID(password, encodeData)
		if err != nil {
			panic(err)
		}

		//Encode source data with padding 20
		encodePaddingData := mixer.EncodeIDPadding(password, source, 20)

		//Decode source padding data
		decodePaddingData, err := mixer.DecodeID(password, encodeData)
		if err != nil {
			panic(err)
		}

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\nencodePadding(20): %v\ndecodePadding(20): %v\n-------\n",
			source, encodeData, decodeData, encodePaddingData, decodePaddingData)
	}

}
