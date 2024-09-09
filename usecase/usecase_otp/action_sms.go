package usecase_otp

import (
	"encoding/json"
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
	"strconv"
)

func (usecase *OtpUsecaseImpl) SmsAction(kosong interface{}, idRequest, phoneNumber, otp, jenisTransaksi string) (httpCode int, res dto.Res) {
	menit := strconv.Itoa(int(usecase.OtpExpiredSecond) / 60)
	pesan := "(RAHASIA) Jangan Berikan kepada pihak lain / yang mengaku BRI. Kode OTP Anda adalah " + otp + " (berlaku " + menit + " menit). Info Contact BRI 1500017"
	reqUpstream := dto.ReqUpstreamSms{
		Rekening:    "0",
		PhoneNumber: phoneNumber,
		Pesan:       pesan,
		SendTime:    helpers.Now(),
	}
	if jenisTransaksi == "limit" {
		reqUpstream.Produk = usecase.BrigateConfig.SmsProductNameBilling
	} else if jenisTransaksi == "bill" {
		reqUpstream.Produk = usecase.BrigateConfig.SmsProductNameLimit
	} else {
		httpCode, res = helpers.ResInvalidValue(kosong)
		return
	}

	reqUpstreamString, _ := json.Marshal(reqUpstream)
	logUpstream := entity.UpstreamServiceRequestLog{
		Id:               helpers.GenerateUUID(),
		IdRequest:        idRequest,
		RequestPayload:   string(reqUpstreamString),
		RequestTimestamp: helpers.Now(),
		Service:          "ESB",
		Action:           "F10 - API BRI Notification (Pengiriman Notif SMS Blast V2)",
	}

	resUpstream, err := usecase.ExternalRepository.SmsRepository(reqUpstream)
	logUpstream.Url = resUpstream.Request.URL
	logUpstream.ResponseTimestamp = helpers.Now()

	if err != nil {
		logUpstream.ResponsePayload = err.Error()
		logUpstream.IsSuccess = 0
		httpCode, res = helpers.ResBackendError(kosong)
	} else {
		logUpstream.ResponsePayload = resUpstream.String()

		resStruct := dto.ResUpstreamSms{}
		json.Unmarshal(resUpstream.Body(), &resStruct)

		if resStruct.ResponseCode != "00" {
			logUpstream.IsSuccess = 0
			httpCode, res = helpers.ResBackendError(kosong)
		} else {
			logUpstream.IsSuccess = 1
			httpCode, res = helpers.ResSuccess(true, "0000", "Successfully", kosong)
		}
	}

	paramInsertLogUpstream := helpers.BuildParamInsertLogUpstream(logUpstream, httpCode, res, kosong)
	httpCode, res = usecase.LogUsecase.InsertLogUpstreamUsecase(paramInsertLogUpstream)

	return
}
