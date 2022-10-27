package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Item_Code   string  `gorm:"not null;unique;type:varchar(16)"`
	Description string  `gorm:"not null;type:varchar(255)"`
	Quantity    int     `gorm:"not null;type:int"`
	Order       []Order `gorm:"foreignKey:ItemID"`
}
