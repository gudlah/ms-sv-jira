package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"golang.org/x/crypto/bcrypt"
)

func (usecase *AuthUsecaseImpl) Authentication(kosong interface{}, loginRequest dto.ReqLogin, password string) (httpCode int, res dto.Res) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginRequest.Password))
	if err != nil {
		httpCode, res = helpers.ResInvalidCredential(kosong)
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(loginRequest.Password))
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully login", kosong)
		}
	}
	return
}
