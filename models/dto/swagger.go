package dto

type SwaggerLoginSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         *struct {
		Token     string `json:"token" example:"QmVhcmVyIGV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUp3WlhKdWNpSTZJa=="`
		ExpiredAt string `json:"expiredAt" example:"2023-12-28  02:05:27 +07"`
	} `json:"data"`
}

type SwaggerCheckCreditCardSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         *struct {
		Name        string `json:"name" example:"Satrio Pinayungan"`
		PhoneNumber string `json:"phoneNumber" example:"081208120812"`
		CardNumber  string `json:"cardNumber" example:"6189********0002"`
	} `json:"data"`
}

type SwaggerResendOtpSuccess struct {
	Success      bool      `json:"success" example:"true"`
	ResponseCode string    `json:"responseCode" example:"0000"`
	Message      string    `json:"message" example:"Successfully"`
	Data         *struct{} `json:"data"`
}

type SwaggerResendOtpReached struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"1010"`
	Message      string    `json:"message" example:"Request limit reached"`
	Data         *struct{} `json:"data"`
}

type SwaggerCheckBillSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         *struct {
		BillName                  string `json:"billName" example:"Satrio Pinayungan"`
		BillAmount                string `json:"billAmount" example:"10"`
		BillMinAmount             string `json:"billMinAmount" example:"0"`
		BillDueDate               string `json:"billDueDate" example:"20130620"`
		CardNumber                string `json:"cardNumber" example:"6189********0002"`
		StatementBalance          string `json:"statementBalance" example:"0000000010"`
		StatementBalanceIndicator string `json:"statementBalanceIndicator" example:"+"`
		TotalDueAmount            string `json:"totalDueAmount" example:"0000000010"`
		TotalDueAmountIndicator   string `json:"totalDueAmountIndicator" example:" "`
	} `json:"data"`
}

type SwaggerCheckLimitBillReached struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"1009"`
	Message      string    `json:"message" example:"OTP limit reached"`
	Data         *struct{} `json:"data"`
}

type SwaggerCheckLimitSuccess struct {
	Success      bool   `json:"success" example:"true"`
	ResponseCode string `json:"responseCode" example:"0000"`
	Message      string `json:"message" example:"Successfully"`
	Data         *struct {
		CreditLimit              string `json:"creditLimit" example:"9000000"`
		CashLimit                string `json:"cashLimit" example:"0"`
		CardNumber               string `json:"cardNumber" example:"6189********0002"`
		AvailCredit              string `json:"availCredit" example:"8619230"`
		AvailCreditIndicator     string `json:"availCreditIndicator" example:"+"`
		AvailableCredit          string `json:"availableCredit" example:"8619230"`
		AvailableCreditIndicator string `json:"availableCreditIndicator" example:"+"`
		DailyCashLimit           string `json:"dailyCashLimit" example:"8619230"`
		AvailCashRemain          string `json:"availCashRemain" example:"8619230"`
		AvailCashRemainIndicator string `json:"availCashRemainIndicator" example:"+"`
	} `json:"data"`
}

type SwaggerInvalidValue struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"1002"`
	Message      string    `json:"message" example:"Invalid value"`
	Data         *struct{} `json:"data"`
}

type SwaggerInvalidCredential struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"1001"`
	Message      string    `json:"message" example:"Invalid credential"`
	Data         *struct{} `json:"data"`
}

type SwaggerGeneralSystemError struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"2001"`
	Message      string    `json:"message" example:"General system error"`
	Data         *struct{} `json:"data"`
}

type SwaggerBackendError struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"2002"`
	Message      string    `json:"message" example:"Backend error"`
	Data         *struct{} `json:"data"`
}

type SwaggerIpBlocked struct {
	Success      bool      `json:"success" example:"false"`
	ResponseCode string    `json:"responseCode" example:"2004"`
	Message      string    `json:"message" example:"IP blocked"`
	Data         *struct{} `json:"data"`
}
