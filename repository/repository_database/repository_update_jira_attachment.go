package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraAttachmentRepository(attachment entity.JiraAttachments) error {
	return repository.Database.
		Save(attachment).
		Error
}
