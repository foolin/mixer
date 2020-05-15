package mixer

import (
	"errors"
	"fmt"
)

const (
	//CharsCaseAlphanumeric the alphanumeric include upper and lower:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	CharsCaseAlphanumeric = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//CharsUpperAlphanumeric the alphanumeric include upper:`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	CharsUpperAlphanumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//CharsLowerAlphanumeric the alphanumeric include lower:`0123456789abcdefghijklmnopqrstuvwxyz`
	CharsLowerAlphanumeric = "0123456789abcdefghijklmnopqrstuvwxyz"

	//CharsUpperAlphabet the upper alphabet:`ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	CharsUpperAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//CharsLowerAlphabet the lower alphabet:`abcdefghijklmnopqrstuvwxyz`
	CharsLowerAlphabet = "abcdefghijklmnopqrstuvwxyz"

	//CharsUpperHex the hex alphabet and numeric:`0123456789abcdef`
	CharsUpperHex = "0123456789ABCDEF"

	//CharsLowerHex the hex alphabet and numeric:`0123456789abcdef`
	CharsLowerHex = "0123456789abcdef"

	//CharsNumeric the numeric:`0123456789abcdef`
	CharsNumeric = "0123456789"
)

//AlphanumericType alphanumeric type
type AlphanumericType int

const (
	//AlphanumericCase AlphanumericType is CharsCaseAlphanumeric
	AlphanumericCase AlphanumericType = iota

	//AlphanumericUpper AlphanumericType is CharsUpperAlphanumeric
	AlphanumericUpper

	//AlphanumericLower AlphanumericType is CharsLowerAlphanumeric
	AlphanumericLower
)

//Mixer a mixer instance for encode/decode
type Mixer struct {
	saltSeed       int64
	mapEncodeChars map[rune]rune
	mapDecodeChars map[rune]rune
}

//NewWithChars create a new mixer
func NewWithChars(salt string, chars string, candidateChars ...string) (*Mixer, error) {
	if chars == "" {
		return nil, errors.New("at least one of `dictChars` parameters is required")
	}
	seed := sumSaltSeed(salt)
	mapEncodeTable := make(map[rune]rune, 0)
	mapEncodeTable = appendChars(mapEncodeTable, chars, seed)
	for _, v := range candidateChars {
		mapEncodeTable = appendChars(mapEncodeTable, v, seed)
	}
	if len(mapEncodeTable) < 2 {
		return nil, fmt.Errorf("dict chars `%v` is not invalid", chars)
	}
	mapDecodeTable := make(map[rune]rune, 0)
	for k, v := range mapEncodeTable {
		mapDecodeTable[v] = k
	}
	return &Mixer{
		saltSeed:       seed,
		mapEncodeChars: mapEncodeTable,
		mapDecodeChars: mapDecodeTable,
	}, nil
}

//MustNewWithChars must create a new mixer
func MustNewWithChars(salt string, chars string, candidateChars ...string) *Mixer {
	mixer, err := NewWithChars(salt, chars, candidateChars...)
	if err != nil {
		panic(err)
	}
	return mixer
}

//New create a new mixer with case sensitive alphanumeric
func New(salt string) *Mixer {
	return Newt(salt, AlphanumericCase)
}

//Newt create a new mixer with alphanumeric and AlphanumericType
func Newt(salt string, alphanumericType AlphanumericType) *Mixer {
	switch alphanumericType {
	case AlphanumericUpper:
		return MustNewWithChars(salt, CharsUpperAlphanumeric)
	case AlphanumericLower:
		return MustNewWithChars(salt, CharsLowerAlphanumeric)
	case AlphanumericCase:
		return MustNewWithChars(salt, CharsCaseAlphanumeric)
	}
	return MustNewWithChars(salt, CharsCaseAlphanumeric)
}

//Encode encode char array
func (m Mixer) Encode(data []rune) []rune {
	outChars := make([]rune, len(data))
	for i, c := range data {
		if v, ok := m.mapEncodeChars[c]; ok {
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
		if rc, ok := m.mapDecodeChars[c]; ok {
			outChars[i] = rc
		} else {
			outChars[i] = c
		}
	}
	return outChars
}

//EncodeString encode string
func (m Mixer) EncodeString(data string) string {
	return string(m.Encode([]rune(data)))
}

//DecodeString decode string
func (m Mixer) DecodeString(data string) string {
	return string(m.Decode([]rune(data)))
}

func uniqueChars(chars string) []rune {
	// Use map to record duplicates as we find them.
	mapCheck := make(map[rune]bool)
	list := make([]rune, 0)
	for _, v := range []rune(chars) {
		if mapCheck[v] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			mapCheck[v] = true
			// Append to list slice.
			list = append(list, v)
		}
	}
	// Return the new slice.
	return list
}

func appendChars(dictMaps map[rune]rune, chars string, seed int64) map[rune]rune {
	dictChars := uniqueChars(chars)
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
	rn := NewLGC(seed)
	perm := rn.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func randomDecode(chars []rune, seed int64) []rune {
	src := chars
	final := make([]rune, len(src))
	rn := NewLGC(seed)
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
