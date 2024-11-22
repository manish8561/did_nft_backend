package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manish-antiersoultions/dynamic_nft_backend/modules/clients"
	"github.com/manish-antiersoultions/dynamic_nft_backend/utils/helpers"
)

/**
 * access middleware for rate
 */
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_key := c.GetHeader("X-ACCESS-KEY")
		if api_key == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "You are not authorized!",
				"success": false,
			})
			return
		}
		//retriving from db
		var retrievedClient clients.Client
		result := helpers.DB.Where(&clients.Client{API_KEY: api_key}).First(&retrievedClient)

		if result.Error != nil {
			log.Println("----DB error in middleware---------", result.Error)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "You are not authorized!",
				"success": false,
			})
			return
		}
		// for pending request waiting for approval
		if retrievedClient.Status == "pending" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Your request is waiting for approval!",
				"success": false,
			})
			return
		}
		//for non paid users
		if !retrievedClient.Is_Purchased && retrievedClient.Rate_Limits >= clients.RATE_LIMIT {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Your rate limits are exceeded!",
				"success": false,
			})
			return
		}

		//increment the request in background
		retrievedClient.Rate_Limits = retrievedClient.Rate_Limits + 1
		helpers.DB.Save(&retrievedClient)

		// Call the next middleware/handler in the chain
		c.Next()

	}
}
