package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertLogUpstreamRepository(log entity.UpstreamServiceRequestLog) error {
	return repository.Database.Create(log).Error
}
