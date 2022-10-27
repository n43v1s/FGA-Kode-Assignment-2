package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName string `gorm:"not null;type:varchar(255)"`
	OrderedAt    string `gorm:"not null;type:timestamp"`
	ItemID       int    `gorm:"not null;type:int"`
}
