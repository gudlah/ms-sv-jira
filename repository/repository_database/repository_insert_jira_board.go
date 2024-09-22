package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertJiraBoardRepository(jiraBoards []entity.JiraBoards) error {
	return repository.
		Database.
		Create(&jiraBoards).
		Error
}
