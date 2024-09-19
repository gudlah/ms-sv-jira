package delivery_jira_full

import "github.com/gin-gonic/gin"

type JiraFullDelivery interface {
	GetAllFullDelivery(ginContext *gin.Context)
}
