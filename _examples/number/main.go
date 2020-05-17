package main

import (
	"fmt"
	"github.com/foolin/mixer"
	"time"
)

func main() {

	//the source to be encrypted
	sources := []int64{
		123,
		123456,
		1234567890,
		999999999,
		time.Now().UnixNano(),
	}

	//password
	password := "1q2w3e"

	//foreach every source
	for _, source := range sources {

		//Encode source data
		encodeData := mixer.EncodeNumber(password, source)

		//Decode source data
		decodeData := mixer.DecodeNumber(password, encodeData)

		//Encode source data with padding 20
		encodePaddingData := mixer.EncodeNumberPadding(password, source, 20)

		//Decode source padding data
		decodePaddingData := mixer.DecodeNumber(password, encodeData)

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\nencodePadding(20): %v\ndecodePadding(20): %v\n-------\n",
			source, encodeData, decodeData, encodePaddingData, decodePaddingData)
	}

}
