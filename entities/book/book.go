package book

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Status          string `gorm:"type:enum(booked,paid,cancel)"`
	Phone           string `gorm:"type:varchar(14)"`
	CheckIn         time.Time
	CheckOut        time.Time
	UserID          uint
	RoomID          uint
	PaymentMethodID uint
}
