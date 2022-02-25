package room

import (
	"group-project2/entities/image"
	"group-project2/entities/rating"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	Name        string           `gorm:"type:varchar(255);not null;unique"`
	Description string           `gorm:"type:text"`
	Guest       uint             `gorm:"default:1"`
	Bedroom     uint             `gorm:"default:1"`
	HasWifi     bool             `gorm:"type:boolean"`
	HasAc       bool             `gorm:"type:boolean"`
	HasKitchen  bool             `gorm:"type:boolean"`
	HasFp       bool             `gorm:"type:boolean"`
	Longitude   string           `gorm:"type:varchar(255)"`
	Latitude    string           `gorm:"type:varchar(255)"`
	City        string           `gorm:"type:varchar(100)"`
	Price       uint             `gorm:"default:1"`
	Images      []image.Images   `gorm:"foreignKey:RoomID"`
	Ratings     []rating.Ratings `gorm:"foreignKey:RoomID"`
	UserID      uint
}
