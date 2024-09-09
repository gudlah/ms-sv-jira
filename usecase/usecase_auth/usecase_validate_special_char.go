package usecase_auth

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"strings"
)

func (usecase *AuthUsecaseImpl) ValidateSpecialCharUsecase(kosong interface{}, input string, statusCode int, response dto.Res) (httpCode int, res dto.Res) {
	symbols, err := usecase.DatabaseRepository.GetSymbolsRepository()
	if err != nil {
		httpCode, res = helpers.ResGeneralSystemError(kosong)
	} else {
		totalSpecialChars := 0

		for _, symbol := range symbols {
			if strings.Contains(input, symbol.Symbol) {
				totalSpecialChars += 1
			}
		}

		if totalSpecialChars > 0 {
			httpCode, res = helpers.ResInvalidValue(kosong)
		} else {
			httpCode, res = statusCode, response
		}
	}

	return
}
