package clients

import (
	"log"
	"time"

	"github.com/manish-antiersoultions/dynamic_nft_backend/utils/helpers"
)

// maximum limit for using access api
const RATE_LIMIT = 20

type Client struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	Email             string `json:"email" gorm:"size:100;unique"`
	Is_Email_Verified bool   `json:"is_email_verified" gorm:"default:false"`
	API_KEY           string `json:"api_key" gorm:"size:255;unique"`
	Is_Purchased      bool   `json:"is_purchased" gorm:"default:false"`
	Rate_Limits       uint   `json:"rate_limits" gorm:"default:0"`
	Status            string `json:"status" gorm:"size:20;default:pending"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
 * migrate function with user struct
 */
func MigrateSchema() error {
	err := helpers.DB.AutoMigrate(&Client{})
	if err != nil {
		return err
	}
	return nil
}

/**
 * insert row in the database
 */
func createClient(row *Client) (uint, error) {
	result := helpers.DB.Create(row)
	if result.Error != nil {
		log.Println("Error while inserting row: ", result.Error)
		return 0, result.Error
	}
	return row.ID, nil
}
