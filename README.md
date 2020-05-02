# mixer

[![go-doc-img]][go-doc] [![travis-img]][travis] [![go-report-card-img]][go-report-card] [![Coverage Status][cov-img]][cov]


Mixer is a very simple encrypt and decrypt golang library for short strings such as id or hash.


## Features

* **Security** - Support for password(salt) encryption.
* **Symmetrical** Support decrypt from the encrypted string.
* **Equal length** - The length of the encrypted string is equal to the length of the original string.
* **Simple** - The encryption algorithm works by replacing and mixing characters.
* **Custom** Support for custom replacement characters.


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

	//NewHex(salt string, upper bool) create a new Mixer
	mix := mixer.NewHex(salt, false)

	for _, source := range sources {

		//Encode source data
		encodeData := mix.EncodeString(source)

		//Decode source data
		decodeData := mix.DecodeString(encodeData)

		//Output result
		fmt.Printf("-------\nsource: %v\nencode: %v\ndecode: %v\n-------\n", source, encodeData, decodeData)
	}
}

```

Run output:
```

-------
source: 123456
encode: 87e230
decode: 123456
-------
-------
source: 12345abcedf
encode: 47e235f9d0c
decode: 12345abcedf
-------
-------
source: 48656c6c6f204d69786572
encode: 7827858a8629541806bb08
decode: 48656c6c6f204d69786572
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