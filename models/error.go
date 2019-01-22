package models

import "github.com/jinzhu/gorm"

// Error -> error model
type Error struct {
	gorm.Model
	Code        string `json:"code" gorm:"not null;unique"`
	Description string `json:"description,omitempty"`
}
