package dto

type ResDownstreamGetAllUser struct {
	AccountID    string `json:"accountId"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	Locale       string `json:"locale"`
	EmailAddress string `json:"emailAddress"`
}
