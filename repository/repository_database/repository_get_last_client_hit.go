package repository_database

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) GetLastClientHitRepository(clientIp string) (res dto.ResQueryGetLastHit, err error) {
	err = repository.Database.
		Select("COUNT(id) AS jumlah").
		Model(&entity.ActivityLog{}).
		Where("client_ip = ? AND response_code = '1001' AND created_at > (? - INTERVAL ? MINUTE)", clientIp, helpers.Now(), repository.MinuteQueryFail).
		Scan(&res).
		Error
	return
}
