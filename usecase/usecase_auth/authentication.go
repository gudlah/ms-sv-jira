package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"

	"golang.org/x/crypto/bcrypt"
)

func (usecase *AuthUsecaseImpl) Authentication(kosong interface{}, loginRequest dto.ReqLogin, user entity.Users) (httpCode int, res dto.Res) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		httpCode, res = helpers.ResInvalidCredential(kosong)
	} else {
		data, err := usecase.JWTUsecase.GenerateJWTUsecase(user.Id)
		if err != nil {
			httpCode, res = helpers.ResGeneralSystemError(kosong)
		} else {
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully login", data)
		}
	}
	return
}
