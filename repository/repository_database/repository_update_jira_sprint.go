package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraPrioritiesRepository(priority entity.JiraPriorities) error {
	return repository.Database.
		Save(priority).
		Error
}
