package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraSubTaskRepository(subTaskId string) (res entity.JiraSubTasks, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraSubTasks{}).
		Where("sub_task_id = ?", subTaskId).
		Scan(&res).
		Error
	return
}
