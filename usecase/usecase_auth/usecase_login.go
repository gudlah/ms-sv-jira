package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (usecase *AuthUsecaseImpl) LoginUsecase(kosong interface{}, loginRequest dto.ReqLogin) (httpCode int, res dto.Res, idUser string) {
	user, err := usecase.DatabaseRepository.GetUser(loginRequest.Username)
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else if user.Id == "" {
		httpCode, res = helpers.ResInvalidCredential(kosong)
	} else {
		httpCode, res = usecase.Authentication(kosong, loginRequest, user.Password)
	}
	idUser = user.Id
	return
}
