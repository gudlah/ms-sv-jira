package delivery_otp

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/usecase/usecase_log"
	"ms-sv-jira/usecase/usecase_otp"

	"github.com/go-playground/validator/v10"
)

type OtpDeliveryImpl struct {
	AuthDelivery delivery_auth.AuthDelivery
	OtpUsecase   usecase_otp.OtpUsecase
	LogUsecase   usecase_log.LogUsecase
	Validate     *validator.Validate
}

func NewOtpDelivery(authDelivery delivery_auth.AuthDelivery, otpUsecase usecase_otp.OtpUsecase, logUsecase usecase_log.LogUsecase, validate *validator.Validate) OtpDelivery {
	return &OtpDeliveryImpl{
		AuthDelivery: authDelivery,
		OtpUsecase:   otpUsecase,
		LogUsecase:   logUsecase,
		Validate:     validate,
	}
}
