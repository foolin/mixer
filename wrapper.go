package mixer

var globalStdStringMixer = AlphanumericCaseMixer
var globalStdNumberMixer = AlphanumericLowerMixer

//EncodeString encode string use global default mixer
func EncodeString(password string, data string) string {
	return globalStdStringMixer.EncodeString(password, data)
}

//DecodeString encode global default mixer
func DecodeString(password string, data string) string {
	return globalStdStringMixer.DecodeString(password, data)
}

//EncodeNumber encode global default mixer
func EncodeNumber(password string, value int64) string {
	return globalStdNumberMixer.EncodeNumber(password, value)
}

//EncodeNumberPadding  encode padding default number mixer
func EncodeNumberPadding(password string, value int64, paddingLen int) string {
	return globalStdNumberMixer.EncodeNumberPadding(password, value, paddingLen)
}

//DecodeNumber decode default number mixer
func DecodeNumber(password string, data string) int64 {
	return globalStdNumberMixer.DecodeNumber(password, data)
}
