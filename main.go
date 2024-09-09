package main

import (
	"ms-sv-jira/config"
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/helpers"
	"ms-sv-jira/repository/repository_database"
	"ms-sv-jira/routes"
	"ms-sv-jira/usecase/usecase_auth"
	"ms-sv-jira/usecase/usecase_jwt"
	"ms-sv-jira/usecase/usecase_log"
	"os"

	"github.com/go-playground/validator/v10"
)

type RouterParam struct {
	LogUsecase   usecase_log.LogUsecase
	AuthDelivery delivery_auth.AuthDelivery
}

func main() {
	secretKey, err := config.SecretKey("ssl")
	if err != nil {
		helpers.PanicIfError(err)
	}

	os.Setenv("TZ", "Asia/Jakarta")
	konfigurasi := config.Config
	database := config.ConnectDatabase(konfigurasi.Database)
	validate := validator.New()

	// restyConfig := resty.
	// 	New().
	// 	SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
	// 	SetDisableWarn(true)

	databaseRepository := repository_database.NewDatabaseRepository(database, konfigurasi.MinuteQueryFail)
	// externalRepository := repository_external.NewExternalRepository(restyConfig, konfigurasi.Brigate)

	logUsecase := usecase_log.NewLogUsecase(databaseRepository)

	jwtUsecase := usecase_jwt.NewJWTUsecase(validate, secretKey)

	authUsecase := usecase_auth.NewAuthUsecase(databaseRepository, logUsecase, jwtUsecase)
	authDelivery := delivery_auth.NewAuthDelivery(authUsecase, logUsecase, validate)

	routerParam := RouterParam{
		LogUsecase:   logUsecase,
		AuthDelivery: authDelivery,
	}

	routes.
		NewRouter(routes.RouterParam(routerParam)).
		Run(konfigurasi.ServiceHost + ":" + konfigurasi.ServicePort)
}
