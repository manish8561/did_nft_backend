package users

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* user get details
 */

// @Summary get user
// @Schemes
// @Description get user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} UserDetails
// @Router /users/ [get]
func getUserDetailsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

// @Summary post user data to blockchian and IPFS
// @Schemes
// @Description post user data to blockchian and IPFS
// @Tags user
// @Accept json
// @Produce json
// @Param request body UserDecentrializeValidator true "User Decenterlize Data"
// @Security ApiKeyAuth
// @Success 200 {string} TransactionHash
// @Router /users/decentralized [post]
func postUserDecenterlizedHandler(c *gin.Context) {
	var user UserDecentrializeValidator

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	//validating the user
	if err := user.UserDecentrializeValidation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	data, err := postUserDecenterlized(&user, c.ClientIP())
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

// @Summary put user data to blockchian and IPFS
// @Schemes
// @Description put user data to blockchian and IPFS
// @Tags user
// @Accept json
// @Produce json
// @Param request body UserDecentrializeValidator true "User Decenterlize Data"
// @Security ApiKeyAuth
// @Success 200 {string} TransactionHash
// @Router /users/decentralized [put]
func updateUserDecenterlizedHandler(c *gin.Context) {
	var user UserDecentrializeValidator

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	//validating the user
	if err := user.UserDecentrializeValidation(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	data, err := updateUserDecenterlized(&user, c.ClientIP())
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

// @Summary get user signature
// @Schemes
// @Description get user signature
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {string} Signature
// @Router /users/signature/{username} [get]
func getUserSignatureHandler(c *gin.Context) {
	username := c.Param("username")

	data, err := getUserSignature(username)
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

// @Summary get user address
// @Schemes
// @Description get user adress
// @Tags user
// @Accept json
// @Produce json
// @Param request body UserValidator true "User data"
// @Success 200 {string} Wallet Address
// @Router /users/details [post]
func getUserDetailsFromContractHandler(c *gin.Context) {
	var user UserValidator
	//validating the user
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	data, err := getUserDetailsFromContract(&user)
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
// @Tags user
// @Accept json
// @Produce json
// @Param value path string true "Address"
// @Success 200 {string} qrcode
// @Router /users/qrcode/{value} [get]
func getQRCodeHandler(c *gin.Context) {
	value := c.Param("value")

	data, err := getQRCodeImage(value)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate QR code")
		return
	}
	// c.Header("Content-Type", "image/png")
	// c.Data(http.StatusOK, "image/png", data)
	imgBase64 := base64.StdEncoding.EncodeToString(data)

	// createVideo(value)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "data:image/png;base64," + imgBase64,
	})
	return
}

// @Summary post addresses for verification
// @Schemes
// @Description post addresses for verification
// @Tags user
// @Accept json
// @Produce json
// @Param request body AddressesDataValidator true "Addresses Data"
// @Success 200 {string} Success
// @Router /users/verifying [post]
func postUserVerifyingHandler(c *gin.Context) {
	var addressData AddressesDataValidator
	//validating the user
	if err := c.ShouldBind(&addressData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	data, err := verifyAddressData(&addressData)
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
