package dto

type ReqUpstreamSms struct {
	Rekening    string `json:"rekening"`
	Produk      string `json:"produk"`
	PhoneNumber string `json:"phoneNumber"`
	Pesan       string `json:"pesan"`
	PesanWA     string `json:"pesanWA"`
	SendTime    string `json:"sendTime"`
}
