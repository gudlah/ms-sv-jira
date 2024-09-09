package delivery_mock

import "github.com/gin-gonic/gin"

type MockDelivery interface {
	CheckCreditCardDelivery(ginContext *gin.Context)
	ResendOtpDelivery(ginContext *gin.Context)
	CheckBillDelivery(ginContext *gin.Context)
	CheckLimitDelivery(ginContext *gin.Context)
}
