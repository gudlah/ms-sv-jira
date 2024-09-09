package repository_database

import (
	"gorm.io/gorm"
)

type DatabaseRepositoryImpl struct {
	Database        *gorm.DB
	MinuteQueryFail int
}

func NewDatabaseRepository(database *gorm.DB, minuteQueryFail int) DatabaseRepository {
	return &DatabaseRepositoryImpl{
		Database:        database,
		MinuteQueryFail: minuteQueryFail,
	}
}
