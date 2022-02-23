package room

import (
	"group-project2/entities/image"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(255);not null;unique"`
	Description string         `gorm:"type:text"`
	Guest       uint           `gorm:"default:1"`
	Bedroom     uint           `gorm:"default:1"`
	HasWifi     bool           `gorm:"type:boolean"`
	HasAc       bool           `gorm:"type:boolean"`
	HasKitchen  bool           `gorm:"type:boolean"`
	HasFp       bool           `gorm:"type:boolean"`
	Longitude   string         `gorm:"type:varchar(255)"`
	Latitude    string         `gorm:"type:varchar(255)"`
	City        string         `gorm:"type:varchar(100)"`
	Price       uint           `gorm:"default:1"`
	IsAvailable bool           `gorm:"type:boolean"`
	Images      []image.Images `gorm:"foreignKey:RoomID"`
	UserID      uint
}
