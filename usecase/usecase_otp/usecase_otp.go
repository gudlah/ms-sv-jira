package usecase_otp

import "ms-sv-jira/models/dto"

type OtpUsecase interface {
	GenerateOtpUsecase(kosong interface{}, idRequest, jenisTransaksi, phoneNumber string) (httpCode int, res dto.Res)
	ResendOtpUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamResendOtp) (int, dto.Res)
	VerifOtpUsecase(kosong interface{}, param dto.ParamVerifOtp) (int, dto.Res)
}
