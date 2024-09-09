package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) InsertIpBlockRepository(ipBlock entity.IpBlockeds) error {
	return repository.
		Database.
		Create(&ipBlock).
		Error
}
