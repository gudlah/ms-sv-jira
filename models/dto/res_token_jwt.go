package dto

type ResTokenJWT struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
