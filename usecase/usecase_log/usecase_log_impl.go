package usecase_log

import (
	"ms-sv-jira/repository/repository_database"
)

type LogUsecaseImpl struct {
	DatabaseRepository repository_database.DatabaseRepository
}

func NewLogUsecase(databaseRepository repository_database.DatabaseRepository) LogUsecase {
	return &LogUsecaseImpl{
		DatabaseRepository: databaseRepository,
	}
}
