package config

import (
	"ms-sv-jira/helpers"
	"ms-sv-jira/models/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase(dbConfig dto.DatabaseConfig) (database *gorm.DB) {
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	helpers.PanicIfError(err)
	return
}
