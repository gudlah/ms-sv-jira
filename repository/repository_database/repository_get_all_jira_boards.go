package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetAllJiraBoardsRepository(projectId string) (res []entity.JiraBoards, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraBoards{}).
		Where("project_id = ?", projectId).
		Scan(&res).
		Error
	return
}
