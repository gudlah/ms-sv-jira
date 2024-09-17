package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraCardRepository(card entity.JiraCards) error {
	return repository.Database.
		Save(card).
		Error
}
