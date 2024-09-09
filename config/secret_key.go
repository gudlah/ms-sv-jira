package config

import (
	"errors"
	"ms-sv-jira/models/dto"
	"os"
)

func SecretKey(path string) (secretKey dto.SecretKey, errUtama error) {
	var errString string
	publicKey, err := os.ReadFile(path + "/pubkey.pem")
	if err != nil {
		errString = "public key not found"
	}

	privateKey, err := os.ReadFile(path + "/privkey.pem")
	if err != nil {
		if errString != "" {
			errString = "private and public key not found"
		} else {
			errString = "private key not found"
		}
	}
	secretKey = dto.SecretKey{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	if errString != "" {
		errUtama = errors.New(errString)
	}

	return
}
