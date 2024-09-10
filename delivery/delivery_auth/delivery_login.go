package delivery_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
)

func (delivery *AuthDeliveryImpl) LoginDelivery(kosong interface{}, request dto.ReqLogin) (httpCode int, res dto.Res, idUser string) {
	err := delivery.Validate.Struct(request)
	if err != nil {
		httpCode, res = helpers.ResInvalidValue(kosong)
	}
	httpCode, res, idUser = delivery.AuthUsecase.LoginUsecase(kosong, request)

	return
}
