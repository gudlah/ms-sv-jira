package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraSprintRepository(JiraSprints []entity.JiraSprints) error {
	return repository.
		Database.
		Create(&JiraSprints).
		Error
}
