package users

import (
	"github.com/gin-gonic/gin"
)

/**
* registering users routes
 */
func RegisterRoutes(router *gin.Engine) {
	routeGroup := router.Group("/users")

	routeGroup.GET("/", getUserDetailsHandler)
	routeGroup.POST("/details", getUserDetailsFromContractHandler)

	routeGroup.GET("/signature/:username", getUserSignatureHandler)
	routeGroup.GET("/qrcode/:value", getQRCodeHandler)

	routeGroup.POST("/verifying", postUserVerifyingHandler)

	//applying middleware
	routeGroup.Use(APIKeyMiddleware())
	routeGroup.POST("/decentralized", postUserDecenterlizedHandler)
	routeGroup.PUT("/decentralized", updateUserDecenterlizedHandler)

	//migrations functions
	MigrateSchema()
}
