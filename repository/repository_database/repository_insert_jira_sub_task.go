package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraSubTaskRepository(jiraSubTasks []entity.JiraSubTasks) error {
	return repository.
		Database.
		Create(&jiraSubTasks).
		Error
}
