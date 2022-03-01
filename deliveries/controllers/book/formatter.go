package book

import (
	B "group-project2/entities/book"
	"group-project2/repositories/book"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
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
	ID               uint      `json:"id"`
	CheckInReserved  time.Time `json:"check_in_reserved"`
	CheckOutReserved time.Time `json:"check_out_reserved"`
	Phone            string    `json:"phone"`
	RoomID           uint      `json:"room_id"`
	PaymentMethodID  uint      `json:"payment_method_id"`
	Link             string    `json:"link"`
}

func ToResponseCreateBook(Book B.Books, c *coreapi.ChargeResponse) ResponseCreateBook {
	return ResponseCreateBook{
		ID:               Book.ID,
		CheckInReserved:  Book.CheckInReserved,
		CheckOutReserved: Book.CheckOutReserved,
		Phone:            Book.Phone,
		RoomID:           Book.RoomID,
		PaymentMethodID:  Book.PaymentMethodID,
		Link:             c.RedirectURL,
	}
}

type ResponseSetPaid struct {
	ID               uint       `json:"id"`
	CheckInReserved  time.Time  `json:"check_in_reserved"`
	CheckOutReserved time.Time  `json:"check_out_reserved"`
	CheckInTime      *time.Time `json:"check_in_time"`
	CheckOutTime     *time.Time `json:"check_out_time"`
	Status           string     `json:"status"`
	Phone            string     `json:"phone"`
	UserID           uint       `json:"user_id"`
	RoomID           uint       `json:"room_id"`
	PaymentMethodID  uint       `json:"payment_method_id"`
}

func ToResponseSetPaid(Book B.Books) ResponseSetPaid {
	return ResponseSetPaid{
		ID:               Book.ID,
		CheckInReserved:  Book.CheckInReserved,
		CheckOutReserved: Book.CheckOutReserved,
		CheckInTime:      Book.CheckInTime,
		CheckOutTime:     Book.CheckOutTime,
		Status:           Book.Status,
		Phone:            Book.Phone,
		UserID:           Book.UserID,
		RoomID:           Book.RoomID,
		PaymentMethodID:  Book.PaymentMethodID,
	}
}

type ResponseSetCancel struct {
	ID               uint       `json:"id"`
	CheckInReserved  time.Time  `json:"check_in_reserved"`
	CheckOutReserved time.Time  `json:"check_out_reserved"`
	CheckInTime      *time.Time `json:"check_in_time"`
	CheckOutTime     *time.Time `json:"check_out_time"`
	Status           string     `json:"status"`
	Phone            string     `json:"phone"`
	UserID           uint       `json:"user_id"`
	RoomID           uint       `json:"room_id"`
	PaymentMethodID  uint       `json:"payment_method_id"`
}

func ToResponseSetCancel(Book B.Books) ResponseSetCancel {
	return ResponseSetCancel{
		ID:               Book.ID,
		CheckInReserved:  Book.CheckInReserved,
		CheckOutReserved: Book.CheckOutReserved,
		CheckInTime:      Book.CheckInTime,
		CheckOutTime:     Book.CheckOutTime,
		Status:           Book.Status,
		Phone:            Book.Phone,
		UserID:           Book.UserID,
		RoomID:           Book.RoomID,
		PaymentMethodID:  Book.PaymentMethodID,
	}
}

type RequestCheckInTime struct {
	CheckInTime time.Time `json:"check_in_time"`
}

type ResponseCheckInTime struct {
	ID               uint       `json:"id"`
	CheckInReserved  time.Time  `json:"check_in_reserved"`
	CheckOutReserved time.Time  `json:"check_out_reserved"`
	CheckInTime      *time.Time `json:"check_in_time"`
	CheckOutTime     *time.Time `json:"check_out_time"`
	Status           string     `json:"status"`
	Phone            string     `json:"phone"`
	UserID           uint       `json:"user_id"`
	RoomID           uint       `json:"room_id"`
	PaymentMethodID  uint       `json:"payment_method_id"`
}

func ToResponseCheckInTime(Book B.Books) ResponseCheckInTime {
	return ResponseCheckInTime{
		ID:               Book.ID,
		CheckInReserved:  Book.CheckInReserved,
		CheckOutReserved: Book.CheckOutReserved,
		CheckInTime:      Book.CheckInTime,
		CheckOutTime:     Book.CheckOutTime,
		Status:           Book.Status,
		Phone:            Book.Phone,
		UserID:           Book.UserID,
		RoomID:           Book.RoomID,
		PaymentMethodID:  Book.PaymentMethodID,
	}
}

type RequestCheckOutTime struct {
	CheckOutTime time.Time `json:"check_out_time"`
}

type ResponseCheckOutTime struct {
	ID               uint       `json:"id"`
	CheckInReserved  time.Time  `json:"check_in_reserved"`
	CheckOutReserved time.Time  `json:"check_out_reserved"`
	CheckInTime      *time.Time `json:"check_in_time"`
	CheckOutTime     *time.Time `json:"check_out_time"`
	Status           string     `json:"status"`
	Phone            string     `json:"phone"`
	UserID           uint       `json:"user_id"`
	RoomID           uint       `json:"room_id"`
	PaymentMethodID  uint       `json:"payment_method_id"`
}

func ToResponseCheckOutTime(Book B.Books) ResponseCheckOutTime {
	return ResponseCheckOutTime{
		ID:               Book.ID,
		CheckInReserved:  Book.CheckInReserved,
		CheckOutReserved: Book.CheckOutReserved,
		CheckInTime:      Book.CheckInTime,
		CheckOutTime:     Book.CheckOutTime,
		Status:           Book.Status,
		Phone:            Book.Phone,
		UserID:           Book.UserID,
		RoomID:           Book.RoomID,
		PaymentMethodID:  Book.PaymentMethodID,
	}
}

type ResponseGet struct {
	BookID            uint       `json:"book_id"`
	UserID            uint       `json:"user_id"`
	RoomID            uint       `json:"room_id"`
	RoomName          string     `json:"room_name"`
	PaymentMethodID   uint       `json:"payment_method_id"`
	PaymentMethodName string     `json:"payment_method_name"`
	CheckInReserved   time.Time  `json:"check_in_reserved"`
	CheckOutReserved  time.Time  `json:"check_out_reserved"`
	CheckInTime       *time.Time `json:"check_in_time"`
	CheckOutTime      *time.Time `json:"check_out_time"`
	Status            string     `json:"status"`
	Phone             string     `json:"phone"`
}

func ToResponseGet(Books []book.JoinBooks) []ResponseGet {
	Responses := make([]ResponseGet, len(Books))

	for i := 0; i < len(Books); i++ {
		Responses[i].BookID = Books[i].BookID
		Responses[i].UserID = Books[i].UserID
		Responses[i].RoomID = Books[i].RoomID
		Responses[i].RoomName = Books[i].RoomName
		Responses[i].PaymentMethodID = Books[i].PaymentMethodID
		Responses[i].PaymentMethodName = Books[i].PaymentMethodName
		Responses[i].CheckInReserved = Books[i].CheckInReserved
		Responses[i].CheckOutReserved = Books[i].CheckOutReserved
		Responses[i].CheckInTime = Books[i].CheckInTime
		Responses[i].CheckOutTime = Books[i].CheckOutTime
		Responses[i].Status = Books[i].Status
		Responses[i].Phone = Books[i].Phone
	}

	return Responses
}
