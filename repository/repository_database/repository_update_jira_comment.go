package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraCommentRepository(comment entity.JiraComments) error {
	return repository.Database.
		Save(comment).
		Error
}
