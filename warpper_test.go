package mixer

import "testing"

const testWrapperPassword = "mixer123456"

func TestEncodeNumber(t *testing.T) {
	source := int64(1234567890)
	encode := EncodeNumber(testWrapperPassword, source)
	decode := DecodeNumber(testWrapperPassword, encode)
	encodePadding := EncodeNumberPadding(testWrapperPassword, source, 32)
	decodePadding := DecodeNumber(testWrapperPassword, encodePadding)

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
