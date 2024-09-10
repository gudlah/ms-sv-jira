package main

import (
	"ms-sv-jira/config"
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/delivery/delivery_jira"
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/repository/repository_external"
	"ms-sv-jira/routes"
	"ms-sv-jira/usecase/usecase_auth"
	"ms-sv-jira/usecase/usecase_jira"
	"ms-sv-jira/usecase/usecase_log"
	"os"

	"github.com/go-playground/validator/v10"
)

type RouterParam struct {
	LogUsecase   usecase_log.LogUsecase
	AuthDelivery delivery_auth.AuthDelivery
	JiraDelivery delivery_jira.JiraDelivery
}

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	konfigurasi := config.Config
	database := config.ConnectDatabase(konfigurasi.Database)
	restyConfig := config.RestyConfig(konfigurasi.Jira)
	validate := validator.New()

	databaseRepository := repository_database.NewDatabaseRepository(database, konfigurasi.MinuteQueryFail)
	externalRepository := repository_external.NewExternalRepository(restyConfig, konfigurasi.Jira)

	logUsecase := usecase_log.NewLogUsecase(databaseRepository)

	authUsecase := usecase_auth.NewAuthUsecase(databaseRepository, logUsecase)
	authDelivery := delivery_auth.NewAuthDelivery(authUsecase, logUsecase, validate)

	jiraUsecase := usecase_jira.NewJiraUsecase(externalRepository, logUsecase)
	jiraDelivery := delivery_jira.NewJiraDelivery(authDelivery, jiraUsecase, logUsecase, validate)

	routerParam := RouterParam{
		LogUsecase:   logUsecase,
		AuthDelivery: authDelivery,
		JiraDelivery: jiraDelivery,
	}

	routes.
		NewRouter(routes.RouterParam(routerParam)).
		Run(konfigurasi.ServiceHost + ":" + konfigurasi.ServicePort)
}
