package book

import (
	B "group-project2/entities/book"
	"time"
)

type Book interface {
	Insert(NewBook B.Books) (B.Books, error)
	SetPaid(BookID uint) (B.Books, error)
	SetCancel(BookID uint) (B.Books, error)
	SetCheckInTime(BookID uint, CheckInTime time.Time) (B.Books, error)
	SetCheckOutTime(BookID uint, CheckOutTime time.Time) (B.Books, error)
}
