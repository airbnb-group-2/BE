package book

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	CheckInReserved  time.Time
	CheckOutReserved time.Time
	CheckInTime      time.Time
	CheckOutTime     time.Time
	Status           string `gorm:"type:enum(booked,paid,cancel)"`
	Phone            string `gorm:"type:varchar(14)"`
	UserID           uint   `gorm:"not null"`
	RoomID           uint   `gorm:"not null"`
	PaymentMethodID  uint   `gorm:"not null"`
}
