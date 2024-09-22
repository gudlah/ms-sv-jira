package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraSubTaskRepository(subTask entity.JiraSubTasks) error {
	return repository.Database.
		Save(subTask).
		Error
}
