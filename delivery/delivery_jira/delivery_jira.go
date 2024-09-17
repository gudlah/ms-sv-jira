package delivery_jira

import "github.com/gin-gonic/gin"

type JiraDelivery interface {
	GetAllJiraUserDelivery(ginContext *gin.Context)
	GetAllJiraProjectDelivery(ginContext *gin.Context)
}
