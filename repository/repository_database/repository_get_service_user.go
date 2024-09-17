package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) GetServiceUserRepository(username string) (res entity.ServiceUsers, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.ServiceUsers{}).
		Where("username = ? AND is_active = 1", username).
		Scan(&res).
		Error

	return res, err
}
