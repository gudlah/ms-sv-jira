package dto

type ResDownstreamGetAllUser struct {
	AccountID    string `json:"accountId"`
	AccountType  string `json:"accountType"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	Locale       string `json:"locale"`
	EmailAddress string `json:"emailAddress"`
}
