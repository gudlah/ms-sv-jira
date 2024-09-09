package dto

type BrigateConfig struct {
	EsbUrl                 string `env:"ESB_URL"`
	BrigateUrl             string `env:"BRIGATE_URL"`
	Username               string `env:"BRIGATE_USERNAME"`
	Password               string `env:"BRIGATE_PASSWORD"`
	ChannelId              string `env:"ESB_CHANNEL_ID"`
	GetCreditCardServiceId string `env:"ESB_GET_CREDIT_CARD_SERVICE_ID"`
	SmsProductNameLimit    string `env:"SMS_PRODUCT_NAME_LIMIT"`
	SmsProductNameBilling  string `env:"SMS_PRODUCT_NAME_BILLING"`
	SmsAppName             string `env:"SMS_APP_NAME"`
}
