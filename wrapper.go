package mixer

var globalStdStringMixer = AlphanumericCaseMixer
var globalStdNumberMixer = AlphanumericUpperMixer

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
func DecodeNumber(password string, data string) (int64, error) {
	return globalStdNumberMixer.DecodeNumber(password, data)
}

//EncodeID encode global default mixer
func EncodeID(password string, value uint64) string {
	return globalStdNumberMixer.EncodeID(password, value)
}

//EncodeIDPadding  encode padding default number mixer
func EncodeIDPadding(password string, value uint64, paddingLen int) string {
	return globalStdNumberMixer.EncodeIDPadding(password, value, paddingLen)
}

//DecodeID decode default number mixer
func DecodeID(password string, data string) (uint64, error) {
	return globalStdNumberMixer.DecodeID(password, data)
}

//EncodeBase32 encode global default mixer
func EncodeBase32(password string, value string) string {
	return globalStdNumberMixer.EncodeBase32(password, value)
}

//EncodeBase32Padding  encode padding default number mixer
func EncodeBase32Padding(password string, value string, paddingLen int) string {
	return globalStdNumberMixer.EncodeBase32Padding(password, value, paddingLen)
}

//DecodeBase32 decode default number mixer
func DecodeBase32(password string, data string) (string, error) {
	return globalStdNumberMixer.DecodeBase32(password, data)
}
