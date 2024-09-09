package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertSecretKeyRepository(data entity.SecretKey) error {
	return repository.Database.Create(data).Error
}
