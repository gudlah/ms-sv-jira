package delivery_otp

import "github.com/gin-gonic/gin"

type OtpDelivery interface {
	ResendOtpDelivery(ginContext *gin.Context)
}
