package dto

type ResDownstreamCheckBill struct {
	BillName                  string `json:"billName"`
	BillAmount                string `json:"billAmount"`
	BillMinAmount             string `json:"billMinAmount"`
	BillDueDate               string `json:"billDueDate"`
	CardNumber                string `json:"cardNumber"`
	StatementBalance          string `json:"statementBalance"`
	StatementBalanceIndicator string `json:"statementBalanceIndicator"`
	TotalDueAmount            string `json:"totalDueAmount"`
	TotalDueAmountIndicator   string `json:"totalDueAmountIndicator"`
}
