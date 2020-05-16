package mixer

var globalStdStringMixer = AlphanumericCaseMixer
var globalStdNumberMixer = AlphanumericLowerMixer

func EncodeString(password string, data string) string {
	return globalStdStringMixer.EncodeString(password, data)
}

func DecodeString(password string, data string) string {
	return globalStdStringMixer.DecodeString(password, data)
}

func EncodeNumber(password string, value int64) string {
	return globalStdNumberMixer.EncodeNumber(password, value)
}

func EncodeNumberPadding(password string, value int64, paddingLen int) string {
	return globalStdNumberMixer.EncodeNumberPadding(password, value, paddingLen)
}

func DecodeNumber(password string, data string) int64 {
	return globalStdNumberMixer.DecodeNumber(password, data)
}
