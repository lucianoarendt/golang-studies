package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	// ID          uint   `json:"id" gorm:"primaryKey"`
	UserID  uint     `json:"user_id"`
	Name    string   `json:"name"`
	Symbols []Symbol `json:"symbols" gorm:"foreignKey:Project"`
}
