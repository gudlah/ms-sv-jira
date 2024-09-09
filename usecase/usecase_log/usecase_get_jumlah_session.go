package usecase_log

func (usecase *LogUsecaseImpl) GetJumlahSessionUsecase(sessionId, param string) (jumlah int, err error) {
	aEndpoint, bEndpoint := "/api/v1/bill", "/api/v1/limit"
	if param == "otp" {
		aEndpoint, bEndpoint = "/api/v1/resend-otp", "/api/v1/resend-otp"
	}

	data, err := usecase.DatabaseRepository.GetJumlahSessionRepository(sessionId, aEndpoint, bEndpoint)
	jumlah = data.Jumlah

	return
}
