package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) UpdateSecretKeyRepository(data entity.SecretKey) error {
	return repository.Database.Save(&data).Error
}
