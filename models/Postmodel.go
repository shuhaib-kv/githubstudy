package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID   uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique"`
	Password     string `gorm:"not null"`
	Block_status bool
}
