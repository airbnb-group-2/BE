package user

import (
	"group-project2/entities/room"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string       `gorm:"type:varchar(255);not null"`
	Email    string       `gorm:"type:varchar(255);not null;unique"`
	Password string       `gorm:"type:varchar(255);not null"`
	IsRenter bool         `gorm:"type:boolean;default:false"`
	Rooms    []room.Rooms `gorm:"foreignKey:UserID"`
}
