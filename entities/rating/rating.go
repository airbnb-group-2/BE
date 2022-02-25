package rating

import "gorm.io/gorm"

type Ratings struct {
	gorm.Model
	Star   uint `gorm:"not null"`
	UserID uint
	RoomID uint
}
