package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertLogActivityRepository(log entity.ActivityLog) error {
	return repository.Database.Create(log).Error
}
