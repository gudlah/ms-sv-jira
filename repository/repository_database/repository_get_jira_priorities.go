package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetJiraPrioritiesRepository(priorityId int) (res entity.JiraPriorities, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.JiraPriorities{}).
		Where("priority_id = ?", priorityId).
		Scan(&res).
		Error
	return
}
