package dto

type ReqDownstreamCheckLimit struct {
	CardNumber  string `json:"cardNumber" validate:"required,numeric,min=16,max=16"`
	PhoneNumber string `json:"phoneNumber" validate:"required,numeric"`
	OtpCode     string `json:"otpCode" validate:"required,numeric"`
	SessionId   string `json:"sessionId" validate:"required"`
}
