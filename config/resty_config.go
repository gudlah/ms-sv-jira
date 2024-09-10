package config

import (
	"crypto/tls"
	"ms-sv-jira/models/dto"
	"time"

	"github.com/go-resty/resty/v2"
)

func RestyConfig(jiraConfig dto.JiraConfig) *resty.Client {
	return resty.
		New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetDisableWarn(true).
		SetBasicAuth(jiraConfig.Email, jiraConfig.Token).
		SetTimeout(10 * time.Second)
}
