package mixer

import (
	"encoding/hex"
	"testing"
)

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
	salt := "123456"
	alphanumericAndUpperMixer, _ := New(salt, CharsNumeric, CharsLowerAlphabet, CharsUpperAlphabet)
	myCharsChars, _ := New(salt, "0123456789ABCabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_@!", "0123456789ABCEF&^%")
	sources := []string{
		"HelloMixer",
		"Hello@Mixer!",
		"abc4d3f69575dd4123",
		"68656c6c6f20776f726c6421",
		"46653FD803893E4F75696240258265D2",
	}

	mixers := []struct {
		Name  string
		Mixer *Mixer
	}{
		{"AlphanumericCase", NewAlphanumeric(salt, AlphanumericCase)},
		{"AlphanumericUpper", NewAlphanumeric(salt, AlphanumericUpper)},
		{"AlphanumericLower", NewAlphanumeric(salt, AlphanumericLower)},
		{"HexUpper", NewHex(salt, true)},
		{"HexLower", NewHex(salt, false)},
		{"Numeric", NewNumeric(salt)},
		{"alphanumericAndUpperMixer", alphanumericAndUpperMixer},
		{"myCharsChars", myCharsChars},
	}
	for _, source := range sources {
		for _, m := range mixers {
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
