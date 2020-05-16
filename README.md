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

See [Mixer](https://pkg.go.dev/github.com/foolin/mixer)

# Usage


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
		456,
		123456,
		1234567890,
		time.Now().UnixNano(),
	}

	//password
	password := "1q2w3e"

	//foreach every source
	for _, source := range sources {

		//Encode source data, default padding 16 characters
		encodeData := mixer.EncodeNumber(password, source)

		//Decode source data
		decodeData := mixer.DecodeNumber(password, encodeData)

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n",
			source, encodeData, decodeData)
	}

}

```

Run output:
```

-------
source: 123
encode: 5kc8y0x8qwtr59ug
decode: 123
-------
-------
source: 456
encode: zp2niz2rn5gmnr7l
decode: 456
-------
-------
source: 123456
encode: hkigy6pv3axvj5u7
decode: 123456
-------
-------
source: 1234567890
encode: lkidy1pberg8svu7
decode: 1234567890
-------
-------
source: 1589642163813684200
encode: 11pbiubkpp1k7uuye7y
decode: 1589642163813684200
-------

```

[go-doc]: https://pkg.go.dev/github.com/foolin/mixer
[go-doc-img]: https://godoc.org/github.com/foolin/mixer?status.svg
[travis]: https://travis-ci.org/foolin/mixer
[travis-img]: https://travis-ci.org/foolin/mixer.svg?branch=master&t=mixer
[go-report-card]: https://goreportcard.com/report/github.com/foolin/mixer
[go-report-card-img]: https://goreportcard.com/badge/github.com/foolin/mixer
[cov-img]: https://codecov.io/gh/foolin/mixer/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/foolin/mixer