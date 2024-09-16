package delivery_jira

import "github.com/gin-gonic/gin"

type JiraDelivery interface {
	GetAllUserDelivery(ginContext *gin.Context)
	GetAllProjectDelivery(ginContext *gin.Context)
	GetAllSprintDelivery(ginContext *gin.Context)
	GetAllPriorityDelivery(ginContext *gin.Context)
	GetAllCardDelivery(ginContext *gin.Context)
	GetAllSubTaskDelivery(ginContext *gin.Context)
}
