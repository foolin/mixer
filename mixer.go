package mixer

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

const (
	defaultSalt = "mIxODL^%&tz$AAW"
)

var (

	//StdMixer
	StdMixer = AlphanumericLowerMixer

	//AlphanumericCaseMixer the alphanumeric include upper and lower:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphanumericCaseMixer = MustNewWith(defaultSalt, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphanumericLowerMixer the alphanumeric include upper:`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphanumericUpperMixer = MustNewWith(defaultSalt, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphanumericLowerMixer the alphanumeric include lower:`0123456789abcdefghijklmnopqrstuvwxyz`
	AlphanumericLowerMixer = MustNewWith(defaultSalt, "0123456789abcdefghijklmnopqrstuvwxyz")

	//AlphabetCaseMixer the upper alphabet:`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphabetCaseMixer = MustNewWith(defaultSalt, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphabetUpperMixer the upper alphabet:`ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphabetUpperMixer = MustNewWith(defaultSalt, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphabetLowerMixer the lower alphabet:`abcdefghijklmnopqrstuvwxyz`
	AlphabetLowerMixer = MustNewWith(defaultSalt, "abcdefghijklmnopqrstuvwxyz")

	//HexCaseMixer the hex alphabet and numeric:`0123456789abcdefABCDEF`
	HexCaseMixer = MustNewWith(defaultSalt, "0123456789abcdefABCDEF")

	//HexLowerMixer the hex alphabet and numeric:`0123456789abcdef`
	HexUpperMixer = MustNewWith(defaultSalt, "0123456789ABCDEF")

	//HexLowerMixer the hex alphabet and numeric:`0123456789abcdef`
	HexLowerMixer = MustNewWith(defaultSalt, "0123456789abcdef")

	//NumericMixer the numeric:`0123456789abcdef`
	NumericMixer = MustNewWith(defaultSalt, "0123456789")

	//SymbolsMixer the symbols chars
	SymbolsMixer = MustNewWith(defaultSalt, "0123456789ABCabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+-=.")
)

var alphabetsRunes = []rune("abcdefghijklmnopqrstuvwxyz")

type Config struct {
	Salt     string //salt for random seed
	MixChars string //chars for mix
}

//Mixer a mixer instance for encode/decode
type Mixer struct {
	config         Config
	cacheSaltSeeds sync.Map //cache salt seed
	mapEncodeChars map[rune]rune
	mapDecodeChars map[rune]rune
}

//NewWithChars create a new mixer
func NewWithConfig(cfg Config) (*Mixer, error) {
	if cfg.Salt == "" {
		return nil, fmt.Errorf("salt is not allow empty")
	}
	seed := sumSaltSeed(cfg.Salt)
	var cacheSaltSeeds sync.Map
	cacheSaltSeeds.Store(cfg.Salt, seed)
	mapEncodeTable := createMapChars(cfg.MixChars, seed)
	if len(mapEncodeTable) < 2 {
		return nil, fmt.Errorf("mixChars `%v` is not invalid", cfg.MixChars)
	}
	mapDecodeTable := make(map[rune]rune, 0)
	for k, v := range mapEncodeTable {
		mapDecodeTable[v] = k
	}
	return &Mixer{
		config:         cfg,
		cacheSaltSeeds: cacheSaltSeeds,
		mapEncodeChars: mapEncodeTable,
		mapDecodeChars: mapDecodeTable,
	}, nil
}

//MustNewWithChars must create a new mixer
func MustNewWithConfig(cfg Config) *Mixer {
	mixer, err := NewWithConfig(cfg)
	if err != nil {
		panic(err)
	}
	return mixer
}

//New create a new mixer with case sensitive alphanumeric
func New() *Mixer {
	return StdMixer
}

//NewWith create a new mixer with args
func NewWith(salt string, mixChars string) (*Mixer, error) {
	return NewWithConfig(Config{
		Salt:     salt,
		MixChars: mixChars,
	})
}

//MustNewWith must create a new mixer
func MustNewWith(salt string, mixChars string) *Mixer {
	mixer, err := NewWith(salt, mixChars)
	if err != nil {
		panic(err)
	}
	return mixer
}

//WithSalt create copy Mixer with new salt
func (m *Mixer) WithSalt(salt string) *Mixer {
	cfg := m.config
	cfg.Salt = salt
	return MustNewWithConfig(cfg)
}

//Encode encode char array
func (m Mixer) Encode(data []rune) []rune {
	return m.Encodep("", data)
}

//Encodep encode char array
func (m Mixer) Encodep(password string, data []rune) []rune {
	seed := m.getSeed(password)
	outChars := make([]rune, len(data))
	for i, c := range data {
		if v, ok := m.mapEncodeChars[c]; ok {
			outChars[i] = v
		} else {
			outChars[i] = c
		}
	}
	return randomEncode(outChars, seed)
}

//Decode decode char array
func (m Mixer) Decode(data []rune) []rune {
	return m.Decodep("", data)
}

//Decodep decode char array
func (m Mixer) Decodep(password string, data []rune) []rune {
	seed := m.getSeed(password)
	outChars := randomDecode(data, seed)
	for i, c := range outChars {
		if rc, ok := m.mapDecodeChars[c]; ok {
			outChars[i] = rc
		} else {
			outChars[i] = c
		}
	}
	return outChars
}

//EncodeInt64 encode string
func (m Mixer) EncodeNumber(value int64) string {
	return m.EncodeNumberp("", value)
}

//EncodeInt64p encode string
func (m Mixer) EncodeNumberp(password string, value int64) string {
	return m.EncodeNumberPaddingp(password, value, 16)
}

//EncodeInt64p encode string
func (m Mixer) EncodeNumberPadding(value int64, paddingLen int) string {
	return m.EncodeNumberPaddingp("", value, paddingLen)
}

//EncodeInt64p encode string
func (m Mixer) EncodeNumberPaddingp(password string, value int64, paddingLen int) string {
	runes := []rune(strconv.FormatInt(value, 10))
	numLen := len(runes)
	if numLen < paddingLen {
		runes = append(runes, randomAlphabets(paddingLen-numLen)...)
	}
	return string(m.Encodep(password, runes))
}

//DecodeNumber decode string
func (m Mixer) DecodeNumber(data string) int64 {
	return m.DecodeNumberp("", data)
}

//DecodeNumberp decode string
func (m Mixer) DecodeNumberp(password string, data string) int64 {
	decodeRunes := m.Decodep(password, []rune(data))
	numRunes := make([]rune, 0)
	for _, r := range decodeRunes {
		if r >= '0' && r <= '9' {
			numRunes = append(numRunes, r)
		} else {
			break
		}
	}
	val, _ := strconv.ParseInt(string(numRunes), 10, 64)
	return val
}

//EncodeString encode string
func (m Mixer) EncodeString(data string) string {
	return m.EncodeStringp("", data)
}

//EncodeStringp encode string
func (m Mixer) EncodeStringp(password, data string) string {
	return string(m.Encodep(password, []rune(data)))
}

//DecodeString decode string
func (m Mixer) DecodeString(data string) string {
	return m.DecodeStringp("", data)
}

//DecodeStringp decode string
func (m Mixer) DecodeStringp(password, data string) string {
	return string(m.Decodep(password, []rune(data)))
}

//Config return current Config
func (m Mixer) Config() Config {
	return m.config
}

func (m Mixer) getSeed(password string) int64 {
	if password == "" {
		password = m.config.Salt
	}
	seed, ok := m.cacheSaltSeeds.Load(password)
	if ok {
		return seed.(int64)
	}
	saltSeed := sumSaltSeed(password)
	m.cacheSaltSeeds.Store(password, saltSeed)
	return saltSeed
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

func createMapChars(chars string, seed int64) map[rune]rune {
	dictChars := uniqueChars(chars)
	rnChars := randomEncode(dictChars, seed)
	dictMaps := make(map[rune]rune, 0)
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

func randomAlphabets(n int) []rune {
	chars := make([]rune, n)
	size := len(alphabetsRunes)
	for i := 0; i < n; i++ {
		chars[i] = alphabetsRunes[rand.Intn(size)]
	}
	return chars
}
