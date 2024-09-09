package delivery_auth

import "ms-sv-jira/models/dto"

func (delivery *AuthDeliveryImpl) ValidateSpecialCharDelivery(kosong interface{}, input string, statusCode int, response dto.Res) (httpCode int, res dto.Res) {
	httpCode, res = delivery.AuthUsecase.ValidateSpecialCharUsecase(kosong, input, statusCode, response)
	return
}
