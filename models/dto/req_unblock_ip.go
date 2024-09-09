package dto

type ReqUnblockIp struct {
	IpAddress string `json:"ip_address" validate:"required"`
}
