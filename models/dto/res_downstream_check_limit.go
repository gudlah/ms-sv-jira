package dto

type ResDownstreamCheckLimit struct {
	CreditLimit              string `json:"creditLimit"`
	CashLimit                string `json:"cashLimit"`
	CardNumber               string `json:"cardNumber"`
	AvailCredit              string `json:"availCredit"`
	AvailCreditIndicator     string `json:"availCreditIndicator"`
	AvailableCredit          string `json:"availableCredit"`
	AvailableCreditIndicator string `json:"availableCreditIndicator"`
	DailyCashLimit           string `json:"dailyCashLimit"`
	AvailCashRemain          string `json:"availCashRemain"`
	AvailCashRemainIndicator string `json:"availCashRemainIndicator"`
}
