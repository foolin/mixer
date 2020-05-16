package mixer

import (
	"encoding/hex"
	"testing"
	"time"
)

var alphanumericAndUpperMixer = New(salt)
var myCharsChars, _ = NewWithChars(salt, "0123456789ABCabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_@!", "0123456789ABCEF&^%")
var salt = "123456"
var testMixers = []struct {
	Name  string
	Mixer *Mixer
}{
	{"AlphanumericCase", Newt(salt, AlphanumericCase)},
	{"AlphanumericUpper", Newt(salt, AlphanumericUpper)},
	{"AlphanumericLower", Newt(salt, AlphanumericLower)},
	{"HexUpper", MustNewWithChars(salt, CharsUpperHex)},
	{"HexLower", MustNewWithChars(salt, CharsLowerHex)},
	{"Numeric", New(salt)},
	{"alphanumericAndUpperMixer", alphanumericAndUpperMixer},
	{"myCharsChars", myCharsChars},
}

func TestMixer(t *testing.T) {
	runTest(t, true)
}

func TestMixerTimes(t *testing.T) {
	//10,000 tests
	for times := 0; times <= 10000; times++ {
		runTest(t, false)
	}
}

func runTest(t *testing.T, isLog bool) {
	sources := []string{
		"HelloMixer",
		"Hello@Mixer!",
		"abc4d3f69575dd4123",
		"68656c6c6f20776f726c6421",
		"46653FD803893E4F75696240258265D2",
	}

	for _, source := range sources {
		for _, m := range testMixers {
			encodeData := m.Mixer.EncodeString(source)
			decodeData := m.Mixer.DecodeString(encodeData)
			if source != decodeData {
				t.Fatalf("error: decode data not equal\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v",
					m.Name, source, encodeData, decodeData)
			}
			if isLog {
				t.Logf("-------\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v", m.Name, source, encodeData, decodeData)
			}
		}
	}
}

func TestHex(t *testing.T) {
	t.Logf("%v", hex.EncodeToString([]byte("Hello Mixer")))
}

func TestMixer_EncodeInt64(t *testing.T) {
	sources := []int64{
		1,
		12,
		123,
		123456789,
		time.Now().Unix(),
	}
	for _, source := range sources {
		for _, m := range testMixers {
			encodeData := m.Mixer.EncodeInt64(source)
			decodeData, err := m.Mixer.DecodeInt64(encodeData)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if source != decodeData {
				t.Fatalf("error: decode data not equal\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v",
					m.Name, source, encodeData, decodeData)
			}
			t.Logf("-------\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v", m.Name, source, encodeData, decodeData)
		}
	}
}

func TestRandomize(t *testing.T) {
	chars := []rune("abcdefg")
	t.Logf("%v", chars)
	for i := 0; i < 10; i++ {
		rnChars := randomEncode(chars, 12345)
		t.Logf("src:%v|out:%v", chars, rnChars)
		revChars := randomDecode(rnChars, 12345)
		t.Logf("src:%v|rev:%v", chars, revChars)
	}

}
