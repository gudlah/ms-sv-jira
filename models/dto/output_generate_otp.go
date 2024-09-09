package dto

type OutputGenerateOtp struct {
	SecretKey string `json:"secretKey"`
	OtpCode   string `json:"otpCode"`
	ExpiredAt string `json:"expiredAt"`
}
