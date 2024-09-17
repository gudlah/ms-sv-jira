package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraColumnRepository(jiraColumns []entity.JiraColumns) error {
	return repository.
		Database.
		Create(&jiraColumns).
		Error
}
