package mixer

import (
	"encoding/hex"
	"testing"
)

func TestMixer(t *testing.T) {
	salt := "123456"
	alphanumericAndUpperMixer, _ := New(salt, DictLowerAlphanumeric, DictLowerAlphabet)
	myDictChars, _ := New(salt, []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_!*()%$!@"))
	mixers := map[string]*Mixer{
		"AlphanumericCase":          NewAlphanumeric(salt, AlphanumericCase),
		"AlphanumericUpper":         NewAlphanumeric(salt, AlphanumericUpper),
		"AlphanumericLower":         NewAlphanumeric(salt, AlphanumericLower),
		"Hex":                       NewHex(salt),
		"Numeric":                   NewNumeric(salt),
		"alphanumericAndUpperMixer": alphanumericAndUpperMixer,
		"myDictChars":               myDictChars,
	}
	sources := []string{
		"HelloMixer",
		"Hello@Mixer!",
		"abc4d3f69575dd4123",
		"68656c6c6f20776f726c6421",
		"46653FD803893E4F75696240258265D2",
	}

	for _, source := range sources {
		for name, mixer := range mixers {
			encodeData := mixer.EncodeString(source)
			decodeData := mixer.DecodeString(encodeData)
			if source != decodeData {
				t.Fatalf("error: decode data not equal\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v",
					name, source, encodeData, decodeData)
			}
			t.Logf("-------\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v", name, source, encodeData, decodeData)
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
