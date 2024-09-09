package dto

type ActivityLogParam struct {
	IdRequest string
	IdUser    string
	IpClient  string
	SessionId string
	Activity  string
	Request   string
	Response  Res
	HttpCode  int
	Kosong    interface{}
}
