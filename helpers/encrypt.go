package helpers

import "encoding/base64"

func Encrypt(kata string) string {
	return base64.StdEncoding.EncodeToString([]byte(kata))
}
