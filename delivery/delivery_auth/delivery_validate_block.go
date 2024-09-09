package delivery_auth

import "ms-sv-jira/models/dto"

func (delivery *AuthDeliveryImpl) ValidateBlockDelivery(ipClient string, httpCode int, res dto.Res, kosong interface{}) (httpCodeHasil int, resHasil dto.Res) {
	return delivery.AuthUsecase.ValidateBlockUsecase(ipClient, httpCode, res, kosong)
}
