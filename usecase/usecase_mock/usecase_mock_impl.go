package usecase_mock

import (
	"ms-sv-jira/usecase/usecase_log"
	"ms-sv-jira/usecase/usecase_otp"
)

type MockUsecaseImpl struct {
	LogUsecase usecase_log.LogUsecase
	OtpUsecase usecase_otp.OtpUsecase
}

func NewMockUsecase(logUsecase usecase_log.LogUsecase, otpUsecase usecase_otp.OtpUsecase) MockUsecase {
	return &MockUsecaseImpl{
		LogUsecase: logUsecase,
		OtpUsecase: otpUsecase,
	}
}
