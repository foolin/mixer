package mixer

import "testing"

const testWrapperPassword = "mixer123456"

func TestEncodeNumber(t *testing.T) {
	source := int64(1234567890)
	encode := EncodeNumber(testWrapperPassword, source)
	decode, err := DecodeNumber(testWrapperPassword, encode)
	if err != nil {
		t.Fatal(err)
	}
	encodePadding := EncodeNumberPadding(testWrapperPassword, source, 32)
	decodePadding, err := DecodeNumber(testWrapperPassword, encodePadding)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("\n------------\nsource: %v\nencode: %v\ndecode: %v\nencodePadding: %v\ndecodePadding: %v\n------------\n",
		source, encode, decode, encodePadding, decodePadding)

}

func TestEncodeID(t *testing.T) {
	source := uint64(9223372036854775808)
	encode := EncodeID(testWrapperPassword, source)
	decode, err := DecodeID(testWrapperPassword, encode)
	if err != nil {
		t.Fatal(err)
	}
	encodePadding := EncodeIDPadding(testWrapperPassword, source, 32)
	decodePadding, err := DecodeID(testWrapperPassword, encodePadding)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("\n------------\nsource: %v\nencode: %v\ndecode: %v\nencodePadding: %v\ndecodePadding: %v\n------------\n",
		source, encode, decode, encodePadding, decodePadding)

}

func TestEncodeString(t *testing.T) {
	source := "0123456789abcdefgHIJKLMN"
	encode := EncodeString(testWrapperPassword, source)
	decode := DecodeString(testWrapperPassword, encode)

	t.Logf("\n------------\nsource: %v\nencode: %v\ndecode: %v\n------------\n",
		source, encode, decode)
}
