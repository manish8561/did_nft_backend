package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/manish-antiersoultions/dynamic_nft_backend/docs"
	"github.com/manish-antiersoultions/dynamic_nft_backend/modules/clients"
	"github.com/manish-antiersoultions/dynamic_nft_backend/modules/users"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	helpers "github.com/manish-antiersoultions/dynamic_nft_backend/utils/helpers"
	"github.com/manish-antiersoultions/dynamic_nft_backend/utils/middlewares"
)

/**
 * cors common function for * n
 */
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			// c.AbortWithStatus(204)
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	}
}

func init() {
	// file for local environment
	// err := godotenv.Load("./config/.env")
	err := godotenv.Load()

	//change file for prod environnment
	// err := godotenv.Load("config/prod.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//intializing data
	err = helpers.MySQL_Initialize()
	if err != nil {
		log.Println("error with db: ", err)
	}
}

// @title Ancrypto NFT Backend Swagger API
// @version 1.0
// @description This is a ancrypto-sdk backend server for user identity.
// @termsOfService https://www.ancrypto.io/terms-of-service/

// @contact.name Ancrypto Support
// @contact.url https://www.ancrypto.io/service-support/
// @contact.email support@ancrypto.io

// // @license.name Apache 2.0
// // @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-ACCESS-KEY
//	@description				Access key for creating identity
//
//	@securityDefinitions.apikey	ApiKeyAuthAdmin
//	@in							header
//	@name						ADMIN-ACCESS-KEY
//	@description				Access key for creating client
//
// @schemes http https
func main() {
	PORT := os.Getenv("PORT")
	MODE := os.Getenv("GIN_MODE")
	if MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin.
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Apply the rate limiting middleware to all routes
	r.Use(middlewares.RateLimitMiddleware(20, time.Minute))

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Listening server on port: " + PORT,
		})
	})
	//registering modules routes
	users.RegisterRoutes(r)
	clients.RegisterRoutes(r)

	// listen and serve on 0.0.0.0:8080
	// can be overriden with the PORT env var
	r.Run(":" + PORT)
}
