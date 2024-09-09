package dto

type ParamVerifOtp struct {
	PhoneNumber    string `json:"phoneNumber"`
	OtpCode        string `json:"otpCode"`
	JenisTransaksi string `json:"jenisTransaksi"`
}
