package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraSprintRepository(sprint entity.JiraSprints) error {
	return repository.Database.
		Save(sprint).
		Error
}
