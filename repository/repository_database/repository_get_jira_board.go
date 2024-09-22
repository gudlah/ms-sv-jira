package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraBoardRepository(boardId int) (res entity.JiraBoards, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraBoards{}).
		Where("board_id = ?", boardId).
		Scan(&res).
		Error
	return
}
