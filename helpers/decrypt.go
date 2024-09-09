package helpers

import "encoding/base64"

func Decrypt(hash string) (string, error) {
	kata, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return "", err
	} else {
		return string(kata), err
	}
}
