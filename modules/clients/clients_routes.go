package clients

import (
	"github.com/gin-gonic/gin"
)

/**
* registering module routes
 */
func RegisterRoutes(router *gin.Engine) {
	routeGroup := router.Group("/clients")

	routeGroup.GET("/", getClientDetailsHandler)
	routeGroup.GET("/qrcode/:value", getQRCodeHandler)

	routeGroup.Use(AdminMiddleware())
	routeGroup.POST("/", addHandler)
	routeGroup.PATCH("/", statusHandler)

	//migrations functions
	MigrateSchema()
}
