package image

import "gorm.io/gorm"

type Images struct {
	gorm.Model
	Link   string `gorm:"type:text;not null"`
	RoomID uint
}
