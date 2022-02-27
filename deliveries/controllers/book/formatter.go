package book

import (
	B "group-project2/entities/book"
	"time"
)

type RequestCreateBook struct {
	CheckInReserved  time.Time `json:"check_in_reserved" form:"check_in_reserved"`
	CheckOutReserved time.Time `json:"check_out_reserved" form:"check_out_reserved"`
	Phone            string    `json:"phone" form:"phone"`
	RoomID           uint      `json:"room_id" form:"room_id"`
	PaymentMethodID  uint      `json:"payment_method_id" form:"payment_method_id"`
}

func (Req RequestCreateBook) ToEntityBook(UserID uint) B.Books {
	return B.Books{
		CheckInReserved:  Req.CheckInReserved,
		CheckOutReserved: Req.CheckOutReserved,
		Phone:            Req.Phone,
		UserID:           UserID,
		RoomID:           Req.RoomID,
		PaymentMethodID:  Req.PaymentMethodID,
	}
}

type ResponseCreateBook struct {
	CheckInReserved  time.Time `json:"check_in_reserved"`
	CheckOutReserved time.Time `json:"check_out_reserved"`
	Phone            string    `json:"phone"`
	RoomID           uint      `json:"room_id"`
	PaymentMethodID  uint      `json:"payment_method_id"`
}

func ToResponseCreateBook(Book B.Books) ResponseCreateBook {
	return ResponseCreateBook{
		CheckInReserved:  Book.CheckInReserved,
		CheckOutReserved: Book.CheckOutReserved,
		Phone:            Book.Phone,
		RoomID:           Book.RoomID,
		PaymentMethodID:  Book.PaymentMethodID,
	}
}
