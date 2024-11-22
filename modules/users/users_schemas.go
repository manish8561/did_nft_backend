package users

import (
	"log"
	"time"

	"github.com/manish-antiersoultions/dynamic_nft_backend/utils/helpers"
)

type User struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Username        string `json:"username" gorm:"size:30;index"`
	Address         string `json:"address" gorm:"size:50;index"`
	IpAddress       string `json:"ip_address" gorm:"size:50"`
	IpfsUrl         string `json:"ipfs_url" gorm:"size:255"`
	Signature       string `json:"signature" gorm:"size:255"`
	TransactionHash string `json:"transaction_hash" gorm:"size:100"`
	Status          string `json:"status" gorm:"size:20;default:active"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
 * migrate function with user struct
 */
func MigrateSchema() error {
	err := helpers.DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

/**
 * insert user in the database
 */
func createUser(user *User) {
	result := helpers.DB.Create(user)
	if result.Error != nil {
		log.Println("Error while inserting user record: ", result.Error)
	}
	log.Println("User inserted with ID", user.ID)
}
