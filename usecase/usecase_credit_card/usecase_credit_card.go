package usecase_credit_card

import "ms-sv-jira/models/dto"

type CreditCardUsecase interface {
	CheckCreditCardUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamCheckCreditCard) (int, dto.Res)
	CheckLimitUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamCheckLimit) (int, dto.Res)
	CheckBillUsecase(idRequest string, kosong interface{}, bodyRequest dto.ReqDownstreamCheckBill) (int, dto.Res)
}
