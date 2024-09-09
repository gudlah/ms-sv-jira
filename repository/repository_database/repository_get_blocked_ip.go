package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetBlockedIpRepository(ipClient string) (res entity.IpBlockeds, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.IpBlockeds{}).
		Where("client_ip = ? AND deleted_at = ''", ipClient).
		Scan(&res).
		Error
	return
}
