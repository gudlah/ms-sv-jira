package dto

type ServerConfig struct {
	GinMode          string `env:"GIN_MODE"`
	ServiceHost      string `env:"SERVICE_HOST"`
	ServicePort      string `env:"SERVICE_PORT"`
	BaseUrl          string `env:"BASE_URL"`
	AllowOrigin      string `env:"ALLOW_ORIGIN"`
	MinuteQueryFail  int    `env:"MINUTE_QUERY_FAIL"`
	OtpExpiredSecond uint   `env:"OTP_EXPIRED_SECOND"`
	Database         DatabaseConfig
	Jira             JiraConfig
}
