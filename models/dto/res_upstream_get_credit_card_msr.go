package dto

type ResUpstreamGetCreditCardMsr struct {
	Response ResponseUpstreamGetCreditCardMsr `json:"response"`
}

type ResponseUpstreamGetCreditCardMsr struct {
	ErrorCode                 string `json:"errorCode"`
	ResponseCode              string `json:"responseCode"`
	ResponseMessage           string `json:"responseMessage"`
	BillName                  string `json:"billName"`
	BillAmount                string `json:"billAmount"`
	BillMinAmount             string `json:"billMinAmount"`
	BillDueDate               string `json:"billDueDate"`
	CreditLimit               string `json:"creditLimit"`
	CashLimit                 string `json:"cashLimit"`
	AvailCredit               string `json:"availCredit"`
	AvailCreditIndicator      string `json:"availCreditIndicator"`
	BillSign                  string `json:"billSign"`
	BlockCode                 string `json:"blockCode"`
	CardNo                    string `json:"cardNo"`
	ShortName                 string `json:"shortName"`
	BillingCycle              string `json:"billingCycle"`
	DailyCashLimit            string `json:"dailyCashLimit"`
	CurrentBalance            string `json:"currentBalance"`
	CurrentBalanceIndicator   string `json:"currentBalanceIndicator"`
	AvailableCredit           string `json:"availableCredit"`
	AvailableCreditIndicator  string `json:"availableCreditIndicator"`
	AvailCashRemain           string `json:"availCashRemain"`
	AvailCashRemainIndicator  string `json:"availCashRemainIndicator"`
	TotalDueAmount            string `json:"totalDueAmount"`
	TotalDueAmountIndicator   string `json:"totalDueAmountIndicator"`
	StatementBalance          string `json:"statementBalance"`
	StatementBalanceIndicator string `json:"statementBalanceIndicator"`
	AmountLastAdvance         string `json:"amountLastAdvance"`
	AmountLastAdvanceIndx     string `json:"amountLastAdvanceIndx"`
	AmountLastPayment         string `json:"amountLastPayment"`
	AmountLastPaymentIndx     string `json:"amountLastPaymentIndx"`
	AmountLastPurchase        string `json:"amountLastPurchase"`
	AmountLastPurchaseIndx    string `json:"amountLastPurchaseIndx"`
	DateOpened                string `json:"dateOpened"`
	DateClosed                string `json:"dateClosed"`
	BlockedDate               string `json:"blockedDate"`
	CardExpiryDate            string `json:"cardExpiryDate"`
	LastStatementDate         string `json:"lastStatementDate"`
	PaymentDueDate            string `json:"paymentDueDate"`
	LastRetailPurchaseDate    string `json:"lastRetailPurchaseDate"`
	LastCashAdvanceDate       string `json:"lastCashAdvanceDate"`
	PaymentDate               string `json:"paymentDate"`
	TrfAccount                string `json:"trfAccount"`
	DateIntoCollect           string `json:"dateIntoCollect"`
	CmStatus                  string `json:"cmStatus"`
	TodayPayment              string `json:"todayPayment"`
	TodayPaymentInd           string `json:"todayPaymentInd"`
}
