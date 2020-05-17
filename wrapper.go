package mixer

var globalStdCaseMixer = AlphanumericCaseMixer
var globalStdUpperMixer = AlphanumericUpperMixer

//EncodeString encode string use global default mixer
func EncodeString(password string, data string) string {
	return globalStdCaseMixer.EncodeString(password, data)
}

//DecodeString encode global default mixer
func DecodeString(password string, data string) string {
	return globalStdCaseMixer.DecodeString(password, data)
}

//EncodeNumber encode global default mixer
func EncodeNumber(password string, value int64) string {
	return globalStdUpperMixer.EncodeNumber(password, value)
}

//EncodeNumberPadding  encode padding default number mixer
func EncodeNumberPadding(password string, value int64, paddingLen int) string {
	return globalStdUpperMixer.EncodeNumberPadding(password, value, paddingLen)
}

//DecodeNumber decode default number mixer
func DecodeNumber(password string, data string) (int64, error) {
	return globalStdUpperMixer.DecodeNumber(password, data)
}

//EncodeID encode global default mixer
func EncodeID(password string, value uint64) string {
	return globalStdUpperMixer.EncodeID(password, value)
}

//EncodeIDPadding  encode padding default number mixer
func EncodeIDPadding(password string, value uint64, paddingLen int) string {
	return globalStdUpperMixer.EncodeIDPadding(password, value, paddingLen)
}

//DecodeID decode default number mixer
func DecodeID(password string, data string) (uint64, error) {
	return globalStdUpperMixer.DecodeID(password, data)
}

//EncodeBase32 encode global default mixer
func EncodeBase32(password string, value string) string {
	return globalStdUpperMixer.EncodeBase32(password, value)
}

//EncodeBase32Padding  encode padding default number mixer
func EncodeBase32Padding(password string, value string, paddingLen int) string {
	return globalStdUpperMixer.EncodeBase32Padding(password, value, paddingLen)
}

//DecodeBase32 decode default number mixer
func DecodeBase32(password string, data string) (string, error) {
	return globalStdUpperMixer.DecodeBase32(password, data)
}
