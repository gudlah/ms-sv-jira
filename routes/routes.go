package routes

import (
	"ms-sv-jira/config"
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/usecase/usecase_log"

	_ "ms-sv-jira/swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterParam struct {
	LogUsecase   usecase_log.LogUsecase
	AuthDelivery delivery_auth.AuthDelivery
}

func NewRouter(routerParam RouterParam) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	ginMode := config.Config.GinMode
	if ginMode == "debug" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	}
	// router.POST("/api/v1/login", routerParam.AuthDelivery.LoginDelivery)
	// rootRouter := router.Group("/api/v1").Use(middleware.JWTauth(routerParam.LogUsecase, routerParam.JwtUsecase, routerParam.AuthDelivery))
	// {
	// 	rootRouter.POST("/verif-credit-phone", routerParam.CreditCardDelivery.CheckCreditCardDelivery)
	// 	rootRouter.POST("/resend-otp", routerParam.OtpDelivery.ResendOtpDelivery)
	// 	rootRouter.POST("/bill", routerParam.CreditCardDelivery.CheckBillDelivery)
	// 	rootRouter.POST("/limit", routerParam.CreditCardDelivery.CheckLimitDelivery)
	// 	rootRouter.POST("/unblock-ip", routerParam.AuthDelivery.UnblockIpDelivery)

	// 	rootRouter.POST("/verif-credit-phone/mock", routerParam.MockDelivery.CheckCreditCardDelivery)
	// 	rootRouter.POST("/resend-otp/mock", routerParam.MockDelivery.ResendOtpDelivery)
	// 	rootRouter.POST("/bill/mock", routerParam.MockDelivery.CheckBillDelivery)
	// 	rootRouter.POST("/limit/mock", routerParam.MockDelivery.CheckLimitDelivery)
	// }
	router.NoRoute(routerParam.AuthDelivery.PageNotFoundDelivery)

	return router
}
