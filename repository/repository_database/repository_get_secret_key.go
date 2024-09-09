package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetSecretKeyRepository(jenisTransaksi, phoneNumber string) (res entity.SecretKey, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.SecretKey{}).
		Where("phone_number = ? AND type = ?", phoneNumber, jenisTransaksi).
		Scan(&res).
		Error

	return res, err
}
