package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraUserRepository(jiraUser []entity.JiraUsers) error {
	return repository.
		Database.
		Create(&jiraUser).
		Error
}
