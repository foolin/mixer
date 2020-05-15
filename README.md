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


```

Run output:
```
-------
 mixer: New(salt)
source: 123456
encode: W51Y9U
decode: 123456
-------
-------
 mixer: Newt(salt, AlphanumericLower)
source: 123456
encode: 05139e
decode: 123456
-------
-------
 mixer: New(salt)
source: 12345abcedf
encode: J51YZMDUtWS
decode: 12345abcedf
-------
-------
 mixer: Newt(salt, AlphanumericLower)
source: 12345abcedf
encode: r513spoet0n
decode: 12345abcedf
-------
-------
 mixer: New(salt)
source: 48656c6c6f204d69786572
encode: 9Ub5M9Slm19q9bq9DU15M9
decode: 48656c6c6f204d69786572
-------
-------
 mixer: Newt(salt, AlphanumericLower)
source: 48656c6c6f204d69786572
encode: 9eb5p9nlm19q9bq9oe15p9
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