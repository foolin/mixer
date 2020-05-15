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

	mixers := []struct {
		Name  string
		Mixer *mixer.Mixer
	}{
		//create a new mixer with case sensitive alphanumeric
		{"New(salt)", mixer.New(salt)},

		//create a new mixer with lowercase  alphanumeric
		{"Newt(salt, AlphanumericLower)", mixer.Newt(salt, mixer.AlphanumericLower)},
	}

	//foreach every source
	for _, source := range sources {
		//foreach mixer for encode and decode
		for _, m := range mixers {
			//Encode source data
			encodeData := m.Mixer.EncodeString(source)

			//Decode source data
			decodeData := m.Mixer.DecodeString(encodeData)

			//Output result
			fmt.Printf("-------\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v\n-------\n",
				m.Name, source, encodeData, decodeData)
		}
	}

}
