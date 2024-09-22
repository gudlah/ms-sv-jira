package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraSprintRepository(sprintId int) (res entity.JiraSprints, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraSprints{}).
		Where("sprint_id = ?", sprintId).
		Scan(&res).
		Error
	return
}
