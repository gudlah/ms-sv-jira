package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraProjectRepository(project entity.JiraProjects) error {
	return repository.Database.
		Save(project).
		Error
}
