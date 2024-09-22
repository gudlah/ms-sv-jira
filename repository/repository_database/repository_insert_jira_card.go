package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraCardRepository(jiraCards []entity.JiraCards) error {
	return repository.
		Database.
		Create(&jiraCards).
		Error
}
