package repository_database

import "ms-sv-jira/models/entity"

func (repository *DatabaseRepositoryImpl) GetSymbolsRepository() (res []entity.Symbols, err error) {
	err = repository.Database.
		Select("*").
		Model(&entity.Symbols{}).
		Scan(&res).
		Error
	return
}
