package clients

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* client get details
 */

// @Summary get client
// @Schemes
// @Description get client
// @Tags client
// @Accept json
// @Produce json
// @Success 200 {string} ClientDetails
// @Router /clients/ [get]
func getClientDetailsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

// @Summary add client
// @Schemes
// @Description add client email to db
// @Tags client
// @Accept json
// @Produce json
// @Param request body ClientValidator true "Client data"
// @Security ApiKeyAuthAdmin
// @Success 200 {string} success
// @Router /clients/ [post]
func addHandler(c *gin.Context) {
	var client ClientValidator
	//validating the client
	if err := c.ShouldBind(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	data, err := add(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"success": true,
	})
	return
}

// @Summary update status of client
// @Schemes
// @Description update status of client in the db
// @Tags client
// @Accept json
// @Produce json
// @Param request body ClientStatusValidator true "Client status and id"
// @Security ApiKeyAuthAdmin
// @Success 200 {string} success
// @Router /clients/ [patch]
func statusHandler(c *gin.Context) {
	var client ClientStatusValidator
	//validating the client
	if err := c.ShouldBind(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	data, err := statusUpdate(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"success": true,
	})
	return
}
// @Summary get address QRCode
// @Schemes
// @Description get QRCode data
// @Tags client
// @Accept json
// @Produce json
// @Param value path string true "Address"
// @Success 200 {string} qrcode
// @Router /clients/qrcode/{value} [get]
func getQRCodeHandler(c *gin.Context) {
	value := c.Param("value")

	// c.Header("Content-Type", "image/png")
	// c.Data(http.StatusOK, "image/png", data)
	// imgBase64 := "base64.StdEncoding.EncodeToString(data)"

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "data:image/png;base64," + value,
	})
	return
}
