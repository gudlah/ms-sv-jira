package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraColumnRepository(column entity.JiraColumns) error {
	return repository.Database.
		Save(column).
		Error
}
