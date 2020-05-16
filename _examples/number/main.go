package main

import (
	"fmt"
	"github.com/foolin/mixer"
)

func main() {

	//the source to be encrypted
	sources := []int64{
		123,
		456,
		123456,
		1234567890,
	}

	mix := mixer.AlphanumericLowerMixer

	//foreach every source
	for _, source := range sources {
		//Encode source data
		encodeData := mix.EncodeNumber(source)

		//Decode source data
		decodeData := mix.DecodeNumber(encodeData)

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n",
			source, encodeData, decodeData)
	}

}
