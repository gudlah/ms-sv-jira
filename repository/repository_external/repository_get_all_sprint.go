package repository_external

import "github.com/go-resty/resty/v2"

func (repository *ExternalRepositoryImpl) GetAllSprintRepository(boardId string) (*resty.Response, error) {
	return repository.
		RestyConfig.
		R().
		Get(repository.JiraConfig.Url + "/agile/1.0/board/" + boardId + "/sprint")
}
