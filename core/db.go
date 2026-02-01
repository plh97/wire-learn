package core

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/liom-source/wire-learn/models"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// Initialize database connection here
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.UserModel{})
	return db
}
