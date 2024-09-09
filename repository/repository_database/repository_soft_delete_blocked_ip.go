package repository_database

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) SoftDeleteBlockedIpRepository(id string) error {
	return repository.Database.
		Model(&entity.IpBlockeds{}).
		Where("id = ?", id).
		Update("deleted_at", helpers.Now()).
		Error
}
