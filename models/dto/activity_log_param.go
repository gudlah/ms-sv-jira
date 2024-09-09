package dto

type ActivityLogParam struct {
	IdRequest string
	IdUser    string
	IpClient  string
	Endpoint  string
	Request   string
	Response  Res
	HttpCode  int
	Kosong    interface{}
}
