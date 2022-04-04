package models

import "gorm.io/gorm"

type Symbol struct {
	gorm.Model
	// ID          uint   `json:"id" gorm:"primaryKey"`
	ListID uint   `json:"list_id" db:"list"`
	Symbol string `json:"symbol"`
}
