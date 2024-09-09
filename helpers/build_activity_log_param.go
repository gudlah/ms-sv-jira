package helpers

import (
	"ms-sv-jira/models/dto"

	"github.com/gin-gonic/gin"
)

func BuildActivityLogParam(idRequest, request string, httpCode int, kosong interface{}, ginContext *gin.Context, response dto.Res) dto.ActivityLogParam {
	_, activity, ipClient, _, idUser, _ := ConfigInit(ginContext)
	return dto.ActivityLogParam{
		IdRequest: idRequest,
		IdUser:    idUser,
		IpClient:  ipClient,
		Endpoint:  activity,
		Request:   request,
		Response:  response,
		HttpCode:  httpCode,
		Kosong:    kosong,
	}
}
