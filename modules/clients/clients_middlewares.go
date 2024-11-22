package clients

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/**
 * access middleware for rate
 */
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ADMIN_ACCESS_KEY := os.Getenv("ADMIN_ACCESS_KEY")
		api_key := c.GetHeader("ADMIN-ACCESS-KEY")

		if api_key != ADMIN_ACCESS_KEY {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "You are not authorized!",
				"success": false,
			})
			return
		}
		// Call the next middleware/handler in the chain
		c.Next()

	}
}
