package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraUserRepository(user entity.JiraUsers) error {
	return repository.Database.
		Save(user).
		Error
}
