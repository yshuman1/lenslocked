package models

import (
	"github.com/jinzhu/gorm"
)

// Gallery is our image container that visitors view
type Gallery struct {
	gorm.Model
	userID uint   `gorm: "not_null; index`
	Title  string `gorm:"not_null"`
}
