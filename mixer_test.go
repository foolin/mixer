package mixer

import (
	"encoding/hex"
	"testing"
	"time"
)

var salt = "123456"
var password = "1q2w3e4r"
var testMixers = []struct {
	Name  string
	Mixer *Mixer
}{
	{"StdMixer", StdMixer.WithSalt(salt)},
	{"AlphanumericCaseMixer", AlphanumericCaseMixer},
	{"AlphanumericCaseMixerWithSalt", AlphanumericCaseMixer.WithSalt(salt)},
	{"AlphanumericUpperMixer", AlphanumericUpperMixer},
	{"AlphanumericLowerMixer", AlphanumericLowerMixer},
	{"AlphabetCaseMixer", AlphabetCaseMixer},
	{"AlphabetUpperMixer", AlphabetUpperMixer},
	{"AlphabetLowerMixer", AlphabetLowerMixer},
	{"HexCaseMixer", HexCaseMixer},
	{"HexUpperMixer", HexUpperMixer},
	{"HexLowerMixer", HexLowerMixer},
	{"NumericMixer", NumericMixer},
	{"SymbolsMixer", SymbolsMixer},
	{"myDefineMixer", NewWith(salt, "0123456789ABCabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_@!0123456789ABCEF~!@#$%^&*()_+,./\\{}<>[]|")},
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
		"^%&tz$",
	}

	for _, source := range sources {
		for _, m := range testMixers {
			encodeData := m.Mixer.EncodeString(password, source)
			decodeData := m.Mixer.DecodeString(password, encodeData)
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

func TestMixer_Config(t *testing.T) {
	t.Logf("%#v", StdMixer.Config())
}

func TestNumber(t *testing.T) {
	sources := []int64{
		1,
		12,
		123,
		123456789,
		1234567890,
		-1234567890,
	}
	for _, source := range sources {
		for _, m := range testMixers {
			encodeData := m.Mixer.EncodeNumber(password, source)
			decodeData, err := m.Mixer.DecodeNumber(password, encodeData)
			if err != nil {
				t.Fatal(err)
			}
			if source != decodeData {
				t.Fatalf("error: decode data not equal\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v",
					m.Name, source, encodeData, decodeData)
			}
			t.Logf("-------\n mixer: %v\nsource: %v\nencode: %v\ndecode: %v", m.Name, source, encodeData, decodeData)
		}
	}
}

func TestID(t *testing.T) {
	sources := []uint64{
		1,
		12,
		123,
		123456789,
		1234567890,
		uint64(time.Now().Unix()),
	}
	for _, source := range sources {
		for _, m := range testMixers {
			encodeData := m.Mixer.EncodeID(password, source)
			decodeData, err := m.Mixer.DecodeID(password, encodeData)
			if err != nil {
				t.Fatal(err)
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

func TestEncodeEmptyPassowrd(t *testing.T) {
	mix := New()
	source := "0123456789abcdefgHIJKLMN"
	encode := mix.EncodeString("", source)
	decode := mix.DecodeString("", encode)

	t.Logf("\n------------\nsource: %v\nencode: %v\ndecode: %v\n------------\n",
		source, encode, decode)
}
