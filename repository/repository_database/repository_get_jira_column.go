package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraColumnRepository(columnId string) (res entity.JiraColumns, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraColumns{}).
		Where("column_id = ?", columnId).
		Scan(&res).
		Error
	return
}
