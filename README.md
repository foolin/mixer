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


```golang

//EncodeString encode string use global default mixer
func EncodeString(password string, data string) string

//DecodeString encode global default mixer
func DecodeString(password string, data string) string

//EncodeNumber encode global default mixer
func EncodeNumber(password string, value int64) string

//EncodeNumberPadding  encode padding default number mixer
func EncodeNumberPadding(password string, value int64, paddingLen int) string

//DecodeNumber decode default number mixer
func DecodeNumber(password string, data string) int64

```

See [Mixer Godoc](https://pkg.go.dev/github.com/foolin/mixer)

# Usage

- Number(ID) Example:

```golang

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

```

Run output:
```
-------
source: 123
encode: 0kanys350waz3juw
decode: 123
encodePadding(20): 0ka8yss50waz3oulw3nj
decodePadding(20): 123
-------
-------
source: 123456
encode: skiwyapn30j5z0u7
decode: 123456
encodePadding(20): ski3yaan30j5zluw7pw0
decodePadding(20): 123456
-------
-------
source: 1234567890
encode: wkidy1pbens30au7
decode: 1234567890
encodePadding(20): wkijy10bens30zu57pda
decodePadding(20): 1234567890
-------
-------
source: 999999999
encode: aeeeeeewes0n53ee
decode: 999999999
encodePadding(20): aeezeejwes0n5we0eee3
decodePadding(20): 999999999
-------
-------
source: 1589704043390707700
encode: k1dbibbk7dbbdeudeb7
decode: 1589704043390707700
encodePadding(20): k1dbibbk7dbbdwudeb7e
decodePadding(20): 1589704043390707700
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


# Todos

- Other languages Implementations
[ ] Java
[ ] PHP
[ ] Javascript

[go-doc]: https://pkg.go.dev/github.com/foolin/mixer
[go-doc-img]: https://godoc.org/github.com/foolin/mixer?status.svg
[travis]: https://travis-ci.org/foolin/mixer
[travis-img]: https://travis-ci.org/foolin/mixer.svg?branch=master&t=mixer
[go-report-card]: https://goreportcard.com/report/github.com/foolin/mixer
[go-report-card-img]: https://goreportcard.com/badge/github.com/foolin/mixer
[cov-img]: https://codecov.io/gh/foolin/mixer/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/foolin/mixer