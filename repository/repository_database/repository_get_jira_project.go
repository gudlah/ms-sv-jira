package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraProjectRepository(projectId string) (res entity.JiraProjects, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraProjects{}).
		Where("project_id = ?", projectId).
		Scan(&res).
		Error
	return
}
