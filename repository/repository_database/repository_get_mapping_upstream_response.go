package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetMappingUpstreamResponse(endpoint, responseCodeBrigate string) (res entity.UpstreamResponseMapping, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.UpstreamResponseMapping{}).
		Where("endpoint = ? AND response_code_brigate = ?", endpoint, responseCodeBrigate).
		Scan(&res).
		Error

	return res, err
}
