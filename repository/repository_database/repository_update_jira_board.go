package repository_database

import (
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) UpdateJiraBoardRepository(board entity.JiraBoards) error {
	return repository.Database.
		Save(board).
		Error
}
