package usecase_mock

import "ms-sv-jira/models/dto"

type MockUsecase interface {
	CheckCreditCardUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamCheckCreditCard) (int, dto.Res)
	CheckBillUsecase(kosong interface{}, bodyRequest dto.ReqDownstreamCheckBill) (int, dto.Res)
	CheckLimitUsecase(kosong interface{}, bodyRequest dto.ReqDownstreamCheckLimit) (int, dto.Res)
	ResendOtpUsecase(kosong interface{}, idRequest string, bodyRequest dto.ReqDownstreamResendOtp) (int, dto.Res)
}
