package kakaoapi

import (
	"encoding/base64"
)

// EncodeBase64 encodes given bytes array into a base64-encoded string.
func EncodeBase64(bytes []byte) (encoded string) {
	return base64.StdEncoding.EncodeToString(bytes)
}

// DecodeBase64 decodes given base64-encoded string into a bytes array.
func DecodeBase64(encoded string) (decoded []byte, err error) {
	_, err = base64.StdEncoding.Decode(decoded, []byte(encoded))

	return decoded, err
}
