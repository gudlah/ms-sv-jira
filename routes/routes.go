package routes

import (
	"ms-sv-jira/delivery/delivery_auth"
	"ms-sv-jira/delivery/delivery_jira"
	"ms-sv-jira/delivery/delivery_jira_full"
	"ms-sv-jira/middleware"
	"ms-sv-jira/usecase/usecase_log"

	// _ "ms-sv-jira/swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterParam struct {
	LogUsecase       usecase_log.LogUsecase
	AuthDelivery     delivery_auth.AuthDelivery
	JiraDelivery     delivery_jira.JiraDelivery
	JiraFullDelivery delivery_jira_full.JiraFullDelivery
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

	// ginMode := config.Config.GinMode
	// if ginMode == "debug" {
	// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	// }
	rootRouter := router.Group("/api/v1").Use(middleware.BasicAuth(routerParam.AuthDelivery, routerParam.LogUsecase))
	{
		rootRouter.GET("/users", routerParam.JiraDelivery.GetAllUserDelivery)
		rootRouter.GET("/projects", routerParam.JiraDelivery.GetAllProjectDelivery)
		rootRouter.POST("/sprints", routerParam.JiraDelivery.GetAllSprintDelivery)
		rootRouter.GET("/priorities", routerParam.JiraDelivery.GetAllPriorityDelivery)
		rootRouter.POST("/cards", routerParam.JiraDelivery.GetAllCardDelivery)
		rootRouter.POST("/subtasks", routerParam.JiraDelivery.GetAllSubTaskDelivery)
		rootRouter.GET("/all", routerParam.JiraFullDelivery.GetAllFullDelivery)
	}
	router.NoRoute(routerParam.AuthDelivery.PageNotFoundDelivery)

	return router
}
