package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetAllJiraProjectRepository() (res []entity.JiraProjects, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraProjects{}).
		Scan(&res).
		Error
	return
}
