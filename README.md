# mixer

[![go-doc-img]][go-doc] [![travis-img]][travis] [![go-report-card-img]][go-report-card] [![Coverage Status][cov-img]][cov]


Mixer is a very simple encrypt and decrypt golang library for short strings such as id or hash.


## Features

* **Security** - Support for password(salt) encryption.
* **Symmetrical** Support decrypt from the encrypted string.
* **Equal length** - The length of the encrypted string is equal to the length of the original string.
* **Simple** - The encryption algorithm works by replacing and mixing characters.
* **Custom** Support for custom replacement characters.
* **LCG algorithm** Use Linear Congruential Generator (LCG) algorithm to generate pseudorandom numbers.
* **ID Number Padding** Support number padding for ID number(default is padding 16 characters)

## Install

```bash
go get -u github.com/foolin/mixer
```

Or get the specified version:
```bash
go get github.com/foolin/mixer@{version}
```
The {version} release list: <https://github.com/foolin/mixer/releases>

# Docs

See [Mixer Godoc](https://pkg.go.dev/github.com/foolin/mixer)

# Usage

- ID Example:

```golang

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


```

Run output:
```
-------
source: 0
encode: 0BAN5S350WAZ3JQW
decode: 0
encodePadding(20): 0BA85SS50WAZ3OQLW3NJ
decodePadding(20): 0
-------
-------
source: 123
encode: S0QWFABN30J5Z0QH
decode: 123
encodePadding(20): S0Q3FAAN30J5ZLQWHBW0
decodePadding(20): 123
-------
-------
source: 123456
encode: B0Q8FC4OQ3NASWQH
decode: 123456
encodePadding(20): B0Q0FC5OQ3NASJQ0H48W
decodePadding(20): 123456
-------
-------
source: 1234567890
encode: K0Q8FC4OQSL33ZQH
decode: 1234567890
encodePadding(20): K0QAFCWOQSL333QBH48Z
decodePadding(20): 1234567890
-------
-------
source: 999999999
encode: 77JLF0SFNS9JBZNZ
decode: 999999999
encodePadding(20): 77J3F0AFNS9JBNNWZSLZ
decodePadding(20): 999999999
-------
-------
source: 9223372036854775808
encode: 059Y4Z77Q0052GNN5W0HFQWZFHSHM50
decode: 9223372036854775808
encodePadding(20): 059Y4Z77Q0052GNN5W0HFQWZFHSHM50
decodePadding(20): 9223372036854775808
-------
-------
source: 18446744073709551615
encode: YL5KS2O7QL89GCQQW48H9QWZFCWZMW7H
decode: 18446744073709551615
encodePadding(20): YL5KS2O7QL89GCQQW48H9QWZFCWZMW7H
decodePadding(20): 18446744073709551615
-------

```

- String(Hash) Example:


```golang

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


```

Run output:
```

-------
source: abc012345edf
encode: hruC87FfRv5N
decode: abc012345edf
-------
-------
source: 0123456789abcdefghijklmnopqrstuvwxyz
encode: 5B6oSK3Q1wfrvHNqjRp87LhlucJ0Xn2CFaIt
decode: 0123456789abcdefghijklmnopqrstuvwxyz
-------
-------
source: 51f5df9e0e802ed54465638ba31158c2
encode: 1rppCfufN1hFR8R7RvQNh5fRuF1uC7vR
decode: 51f5df9e0e802ed54465638ba31158c2
-------

```


# Todo

###  Other languages Implementations

- [ ] Java Implementation
- [ ] PHP Implementation
- [ ] Javascript Implementation

[go-doc]: https://pkg.go.dev/github.com/foolin/mixer
[go-doc-img]: https://godoc.org/github.com/foolin/mixer?status.svg
[travis]: https://travis-ci.org/foolin/mixer
[travis-img]: https://travis-ci.org/foolin/mixer.svg?branch=master&t=mixer
[go-report-card]: https://goreportcard.com/report/github.com/foolin/mixer
[go-report-card-img]: https://goreportcard.com/badge/github.com/foolin/mixer
[cov-img]: https://codecov.io/gh/foolin/mixer/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/foolin/mixer