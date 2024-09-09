package dto

type ReqDownstreamResendOtp struct {
	PhoneNumber    string `json:"phoneNumber" validate:"required,numeric"`
	SessionId      string `json:"sessionId" validate:"required"`
	JenisTransaksi string `json:"jenisTransaksi" validate:"required"`
}
