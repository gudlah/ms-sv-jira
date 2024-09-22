package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraCommentRepository(jiraComments []entity.JiraComments) error {
	return repository.
		Database.
		Create(&jiraComments).
		Error
}
