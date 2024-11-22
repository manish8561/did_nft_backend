package users

import (
	"errors"
	"strings"
)

/**
 * Validating the addresses
 */
type AddressesDataValidator struct {
	Username      string       `form:"username" json:"username" binding:"required,max=20"`
	Addresses []interface{}  `form:"addresses" json:"addresses" binding:"required"`
}

/**
 * User Validator
 */
type UserValidator struct {
	Username string `form:"username" json:"username" binding:"required,max=20,trimData"`
	Network  string `form:"network" json:"network" binding:"required,max=50,trimData"`
}

/**
 * User Decentrialize Validator
 */
type UserDecentrializeValidator struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Address   string `form:"address" json:"address" binding:"required"`
	Data      string `form:"data" json:"data" binding:"required"`
	Signature string `form:"signature" json:"signature" binding:"required"`
}

func (u *UserDecentrializeValidator) UserDecentrializeValidation() error {
	username := strings.TrimSpace(u.Username)
	if username == "" {
		return errors.New("Username is required field!")
	}
	if len(username) > 20 {
		return errors.New("Username is greater than 20 characters!")
	}
	if strings.TrimSpace(u.Address) == "" {
		return errors.New("Address is required field!")
	}
	if strings.TrimSpace(u.Data) == "" {
		return errors.New("Encrypted data is required field!")
	}
	if strings.TrimSpace(u.Signature) == "" {
		return errors.New("Signature is required field!")
	}
	return nil
}
