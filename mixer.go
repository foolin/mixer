package mixer

import (
	"encoding/base32"
	"fmt"
	"strconv"
	"sync"
)

const (
	//DefaultSalt salt for random seed: 202002022002
	DefaultSalt          = "202002022002" //2020.02.02 20:02
	DefaultNumberPadding = 16
)

var (

	//StdMixer std mixer is alias AlphanumericCaseMixer
	StdMixer = AlphanumericCaseMixer

	//AlphanumericCaseMixer the alphanumeric include upper and lower:`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphanumericCaseMixer = NewWith(DefaultSalt, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphanumericUpperMixer the alphanumeric include upper:`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphanumericUpperMixer = NewWith(DefaultSalt, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphanumericLowerMixer the alphanumeric include lower:`0123456789abcdefghijklmnopqrstuvwxyz`
	AlphanumericLowerMixer = NewWith(DefaultSalt, "0123456789abcdefghijklmnopqrstuvwxyz")

	//AlphabetCaseMixer the upper alphabet:`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphabetCaseMixer = NewWith(DefaultSalt, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphabetUpperMixer the upper alphabet:`ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	AlphabetUpperMixer = NewWith(DefaultSalt, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//AlphabetLowerMixer the lower alphabet:`abcdefghijklmnopqrstuvwxyz`
	AlphabetLowerMixer = NewWith(DefaultSalt, "abcdefghijklmnopqrstuvwxyz")

	//HexCaseMixer the hex alphabet and numeric:`0123456789abcdefABCDEF`
	HexCaseMixer = NewWith(DefaultSalt, "0123456789abcdefABCDEF")

	//HexUpperMixer the hex alphabet and numeric:`0123456789abcdef`
	HexUpperMixer = NewWith(DefaultSalt, "0123456789ABCDEF")

	//HexLowerMixer the hex alphabet and numeric:`0123456789abcdef`
	HexLowerMixer = NewWith(DefaultSalt, "0123456789abcdef")

	//NumericMixer the numeric:`0123456789abcdef`
	NumericMixer = NewWith(DefaultSalt, "0123456789")

	//SymbolsMixer the symbols chars
	SymbolsMixer = NewWith(DefaultSalt, "0123456789ABCabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+-=.")
)

var alphabetsRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var base32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)

//Config configuration for new mixer
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

//NewWithConfig create a new mixer
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

//New create a new mixer with case sensitive alphanumeric
func New() *Mixer {
	return StdMixer
}

//NewWith create a new mixer with args
func NewWith(salt string, mixChars string) *Mixer {
	mixer, err := NewWithConfig(Config{
		Salt:     salt,
		MixChars: mixChars,
	})
	if err != nil {
		panic(err)
	}
	return mixer
}

//WithSalt create copy Mixer with new salt
func (m *Mixer) WithSalt(salt string) *Mixer {
	cfg := m.config
	cfg.Salt = salt
	mix, err := NewWithConfig(cfg)
	if err != nil {
		panic(err)
	}
	return mix
}

//Encode encode char array
func (m Mixer) Encode(password string, data []rune) []rune {
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
func (m Mixer) Decode(password string, data []rune) []rune {
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

//EncodeNumber encode int64 number
func (m Mixer) EncodeNumber(password string, value int64) string {
	return m.EncodeNumberPadding(password, value, DefaultNumberPadding)
}

//EncodeNumberPadding encode int64 number
func (m Mixer) EncodeNumberPadding(password string, value int64, paddingLen int) string {
	return m.EncodeBase32Padding(password, strconv.FormatInt(value, 10), paddingLen)
}

//DecodeNumber decode int64 number
func (m Mixer) DecodeNumber(password string, data string) (int64, error) {
	decodeString, err := m.DecodeBase32(password, data)
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseInt(decodeString, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

//EncodeID encode uint64 ID
func (m Mixer) EncodeID(password string, id uint64) string {
	return m.EncodeIDPadding(password, id, DefaultNumberPadding)
}

//EncodeIDPadding encode uint64 ID
func (m Mixer) EncodeIDPadding(password string, id uint64, paddingLen int) string {
	return m.EncodeBase32Padding(password, strconv.FormatUint(id, 10), paddingLen)
}

//DecodeID decode uint64 ID
func (m Mixer) DecodeID(password string, data string) (uint64, error) {
	decodeString, err := m.DecodeBase32(password, data)
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(decodeString, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// DecodeBase32 decode base32 data
func (m Mixer) DecodeBase32(password string, data string) (string, error) {
	decodeRunes := m.Decode(password, []rune(data))
	baseRunes := make([]rune, 0)
	for _, r := range decodeRunes {
		if r != '0' {
			baseRunes = append(baseRunes, r)
		} else {
			break
		}
	}
	value, err := base32NoPadding.DecodeString(string(baseRunes))
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// EncodeBase32 encode base32 data
func (m Mixer) EncodeBase32(password string, value string) string {
	return m.EncodeBase32Padding(password, value, 0)
}

// EncodeBase32Padding encode base32 data
func (m Mixer) EncodeBase32Padding(password string, value string, paddingLen int) string {
	//base32 chars range: [A-Z] [2-7]
	runes := []rune(base32NoPadding.EncodeToString([]byte(value)))
	numLen := len(runes)
	if paddingLen > 0 && numLen < paddingLen {
		seed := m.getSeed(password)
		pdSize := paddingLen - numLen - 1
		runes = append(runes, '0') //zero to split padding chars
		if pdSize > 0 {
			runes = append(runes, randomAlphabets(pdSize, seed)...)
		}
	}
	return string(m.Encode(password, runes))
}

//EncodeString encode string
func (m Mixer) EncodeString(password, data string) string {
	return string(m.Encode(password, []rune(data)))
}

//DecodeString decode string
func (m Mixer) DecodeString(password, data string) string {
	return string(m.Decode(password, []rune(data)))
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

func randomAlphabets(n int, seed int64) []rune {
	chars := make([]rune, n)
	size := len(alphabetsRunes)
	rn := NewLGC(seed)
	for i := 0; i < n; i++ {
		chars[i] = alphabetsRunes[rn.Intn(size)]
	}
	return chars
}
