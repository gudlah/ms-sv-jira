package repository_database

import (
	"ms-sv-jira/models/dto"
	"ms-sv-jira/models/entity"
)

func (repository *DatabaseRepositoryImpl) GetJumlahSessionRepository(sessionId, aEndpoint, bEndpoint string) (res dto.JumlahLog, err error) {
	err = repository.Database.
		Select("COUNT(id) AS jumlah").
		Model(&entity.ActivityLog{}).
		Where("session_id = ? AND activity IN (?, ?)", sessionId, aEndpoint, bEndpoint).
		Scan(&res).
		Error

	return res, err
}
