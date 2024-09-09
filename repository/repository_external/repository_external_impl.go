package repository_external

import (
	"ms-sv-jira/models/dto"

	"github.com/go-resty/resty/v2"
)

type ExternalRepositoryImpl struct {
	RestyConfig   *resty.Client
	BrigateConfig dto.BrigateConfig
}

func NewExternalRepository(restyConfig *resty.Client, brigateConfig dto.BrigateConfig) ExternalRepository {
	return &ExternalRepositoryImpl{
		RestyConfig:   restyConfig,
		BrigateConfig: brigateConfig,
	}
}
