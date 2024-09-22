package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraPrioritiesRepository(JiraPriorities []entity.JiraPriorities) error {
	return repository.
		Database.
		Create(&JiraPriorities).
		Error
}
