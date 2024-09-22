package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraProjectRepository(jiraProject []entity.JiraProjects) error {
	return repository.
		Database.
		Create(&jiraProject).
		Error
}
