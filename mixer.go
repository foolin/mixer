package mixer

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	//DictCaseAlphanumeric the alphanumeric include upper and lower:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	DictCaseAlphanumeric = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//DictUpperAlphanumeric the alphanumeric include upper:`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	DictUpperAlphanumeric = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//DictLowerAlphanumeric the alphanumeric include lower:`0123456789abcdefghijklmnopqrstuvwxyz`
	DictLowerAlphanumeric = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

	//DictUpperAlphabet the upper alphabet:`ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	DictUpperAlphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//DictLowerAlphabet the lower alphabet:`abcdefghijklmnopqrstuvwxyz`
	DictLowerAlphabet = []rune("abcdefghijklmnopqrstuvwxyz")

	//DictHex the hex alphabet and numeric:`0123456789abcdef`
	DictHex = []rune("0123456789abcdef")

	//DictNumeric the numeric:`0123456789abcdef`
	DictNumeric = []rune("0123456789")
)

const (
	AlphanumericCase = iota
	AlphanumericUpper
	AlphanumericLower
)

type Mixer struct {
	saltSeed      int64
	mapEncodeDict map[rune]rune
	mapDecodeDict map[rune]rune
}

//New create a new mixer
func New(salt string, dictChars []rune, moreDictChars ...[]rune) (*Mixer, error) {
	if len(dictChars) <= 0 {
		return nil, errors.New("at least one of `dictChars` parameters is required")
	}
	seed := sumSaltSeed(salt)
	mapEncodeTable := make(map[rune]rune, 0)
	mapEncodeTable = appendDictTable(mapEncodeTable, dictChars, seed)
	for _, v := range moreDictChars {
		mapEncodeTable = appendDictTable(mapEncodeTable, v, seed)
	}
	if len(mapEncodeTable) < 2 {
		return nil, fmt.Errorf("dict chars `%v` is not invalid", string(dictChars))
	}
	mapDecodeTable := make(map[rune]rune, 0)
	for k, v := range mapEncodeTable {
		mapDecodeTable[v] = k
	}
	return &Mixer{
		saltSeed:      seed,
		mapEncodeDict: mapEncodeTable,
		mapDecodeDict: mapDecodeTable,
	}, nil
}

//MustNew must create a new mixer
func MustNew(salt string, dictChars []rune, moreDictChars ...[]rune) *Mixer {
	mixer, err := New(salt, dictChars, moreDictChars...)
	if err != nil {
		panic(err)
	}
	return mixer
}

//MustNew must create a new mixer with alphanumeric
func NewAlphanumeric(salt string, alphanumericType int) *Mixer {
	switch alphanumericType {
	case AlphanumericUpper:
		return MustNew(salt, DictUpperAlphabet)
	case AlphanumericLower:
		return MustNew(salt, DictLowerAlphanumeric)
	case AlphanumericCase:
		return MustNew(salt, DictCaseAlphanumeric)
	}
	return MustNew(salt, DictCaseAlphanumeric)
}

//NewHex must create a new mixer with hex
func NewHex(salt string) *Mixer {
	return MustNew(salt, DictHex)
}

//NewNumeric must create a new mixer with numeric
func NewNumeric(salt string) *Mixer {
	return MustNew(salt, DictNumeric)
}

//Encode encode char array
func (m Mixer) Encode(data []rune) []rune {
	outChars := make([]rune, len(data))
	for i, c := range data {
		if v, ok := m.mapEncodeDict[c]; ok {
			outChars[i] = v
		} else {
			outChars[i] = c
		}
	}
	return randomEncode(outChars, m.saltSeed)
}

//Decode decode char array
func (m Mixer) Decode(data []rune) []rune {
	outChars := randomDecode(data, m.saltSeed)
	for i, c := range outChars {
		if rc, ok := m.mapDecodeDict[c]; ok {
			outChars[i] = rc
		} else {
			outChars[i] = c
		}
	}
	return outChars
}

//Encode encode string
func (m Mixer) EncodeString(data string) string {
	return string(m.Encode([]rune(data)))
}

//DecodeString decode string
func (m Mixer) DecodeString(data string) string {
	return string(m.Decode([]rune(data)))
}

func appendDictTable(dictMaps map[rune]rune, dictChars []rune, seed int64) map[rune]rune {
	rnChars := randomEncode(dictChars, seed)
	if dictMaps == nil {
		dictMaps = make(map[rune]rune, 0)
	}
	for i := 0; i < len(dictChars); i++ {
		key := dictChars[i]
		val := rnChars[i]
		if _, ok := dictMaps[key]; ok {
			continue
		}
		dictMaps[key] = val
	}
	return dictMaps
}

func randomEncode(chars []rune, seed int64) []rune {
	src := chars
	final := make([]rune, len(src))
	rn := rand.New(rand.NewSource(seed))
	perm := rn.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func randomDecode(chars []rune, seed int64) []rune {
	src := chars
	final := make([]rune, len(src))
	rn := rand.New(rand.NewSource(seed))
	perm := rn.Perm(len(src))
	for i, v := range perm {
		final[i] = src[v]
	}
	return final
}

func sumSaltSeed(str string) int64 {
	var sum int64
	for _, v := range []rune(str) {
		sum += int64(v)
	}
	return sum
}
