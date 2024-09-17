package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraUserRepository(userId string) (res entity.JiraUsers, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraUsers{}).
		Where("user_id = ?", userId).
		Scan(&res).
		Error
	return
}
