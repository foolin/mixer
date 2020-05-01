package main

import (
	"fmt"
	"github.com/foolin/mixer"
	"log"
)

func main() {

	salt := "a1b2c3d4"

	sources := []string{
		"1234456",
		"12345abcedf",
		"48656c6c6f204d69786572",
	}

	mix := mixer.NewHex(salt)

	for _, source := range sources {

		encodeData := mix.EncodeString(source)
		decodeData := mix.DecodeString(encodeData)
		if source != decodeData {
			log.Fatalf("error: decode data not equal\nsource: %v\nencode: %v\ndecode: %v",
				source, encodeData, decodeData)
		}
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n", source, encodeData, decodeData)
	}
}
