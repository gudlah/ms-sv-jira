package repository_external

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func (repository *ExternalRepositoryImpl) GetAllProjetRepository() (*resty.Response, error) {
	return repository.
		RestyConfig.
		SetTimeout(10 * time.Second).
		R().
		Get(repository.JiraConfig.Url + "/project")
}
