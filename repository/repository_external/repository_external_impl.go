package repository_external

import (
	"ms-sv-jira/models/dto"

	"github.com/go-resty/resty/v2"
)

type ExternalRepositoryImpl struct {
	RestyConfig *resty.Client
	JiraConfig  dto.JiraConfig
}

func NewExternalRepository(restyConfig *resty.Client, jiraConfig dto.JiraConfig) ExternalRepository {
	return &ExternalRepositoryImpl{
		RestyConfig: restyConfig,
		JiraConfig:  jiraConfig,
	}
}
