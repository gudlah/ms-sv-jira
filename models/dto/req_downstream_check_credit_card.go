package dto

type ReqDownstreamCheckCreditCard struct {
	PhoneNumber    string `json:"phoneNumber" validate:"required,numeric"`
	CardNumber     string `json:"cardNumber" validate:"required,numeric,min=16,max=16"`
	SessionId      string `json:"sessionId" validate:"required"`
	JenisTransaksi string `json:"jenisTransaksi" validate:"required"`
}
