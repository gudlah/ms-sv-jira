package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraAttachmentRepository(jiraAttachment []entity.JiraAttachments) error {
	return repository.
		Database.
		Create(&jiraAttachment).
		Error
}
