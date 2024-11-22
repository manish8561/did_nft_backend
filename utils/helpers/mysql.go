package helpers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize initializes the database connection
func MySQL_Initialize() error {

	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	dsn := MYSQL_USER + `:` + MYSQL_PASSWORD + `@tcp(` + MYSQL_HOST + `:` + MYSQL_PORT + `)/` + MYSQL_DATABASE + `?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}
	log.Println("Mysql is connected successfully.")

	DB = db

	return nil
}
