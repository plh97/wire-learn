package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}
