package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Qualtity  int
	Total     float64
}
