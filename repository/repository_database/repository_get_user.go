package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) GetUser(username string) (res entity.Users, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.Users{}).
		Where("username = ? AND is_active = 1", username).
		Scan(&res).
		Error

	return res, err
}
