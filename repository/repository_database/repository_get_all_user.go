package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetAllJiraUsersRepository() (res []entity.JiraUsers, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraUsers{}).
		Scan(&res).
		Error
	return
}
