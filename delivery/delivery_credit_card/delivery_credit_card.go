package delivery_credit_card

import "github.com/gin-gonic/gin"

type CreditCardDelivery interface {
	CheckCreditCardDelivery(ginContext *gin.Context)
	CheckBillDelivery(ginContext *gin.Context)
	CheckLimitDelivery(ginContext *gin.Context)
}
