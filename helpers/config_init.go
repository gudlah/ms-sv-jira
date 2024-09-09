package helpers

import "github.com/gin-gonic/gin"

func ConfigInit(ginContext *gin.Context) (idRequest, activity, ipClient, transaksi, idUser string, kosong interface{}) {
	idRequest = GenerateUUID()
	activity = ginContext.Request.URL.String()
	ipClient = ginContext.ClientIP()
	transaksi = ginContext.Request.Method + ": " + activity
	idUser = ginContext.GetString("user_id")
	kosong = EmptyObject()
	return
}
