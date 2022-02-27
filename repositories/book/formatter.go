package book

import "time"

type JoinBooks struct {
	BookID            uint
	UserID            uint
	RoomID            uint
	RoomName          string
	PaymentMethodID   uint
	PaymentMethodName string
	CheckInReserved   time.Time
	CheckOutReserved  time.Time
	CheckInTime       time.Time
	CheckOutTime      time.Time
	Status            string
	Phone             string
}
