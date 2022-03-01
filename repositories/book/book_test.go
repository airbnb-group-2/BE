package book

import (
	"group-project2/configs"
	B "group-project2/entities/book"
	I "group-project2/entities/image"
	PM "group-project2/entities/payment-method"
	Rat "group-project2/entities/rating"
	R "group-project2/entities/room"
	U "group-project2/entities/user"
	paymentmethod "group-project2/repositories/payment-method"
	"group-project2/repositories/room"
	"group-project2/repositories/user"
	"group-project2/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	config = configs.GetConfig(true)
	db     = utils.InitDB(config)
)

func TestInsert(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		mock := B.Books{}
		_, err := repo.Insert(mock)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		res, err := repo.Insert(mock)
		assert.Nil(t, err)
		assert.Equal(t, mock.CheckInReserved, res.CheckInReserved)
	})
}

func TestGetAllBooksByUserID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetAllBooksByUserID(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		_, err := repo.GetAllBooksByUserID(1)
		assert.Nil(t, err)
	})
}

func TestGetBookHistoryByUserID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetBookHistoryByUserID(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		_, err := repo.GetBookHistoryByUserID(1)
		assert.Nil(t, err)
	})
}

func TestIsAvailable(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		isAvailable := repo.IsAvailable(1, time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC))
		assert.False(t, isAvailable)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		isAvailable := repo.IsAvailable(1, time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC), time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC))
		assert.True(t, isAvailable)
	})
}

func TestSetPaid(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.SetPaid(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		_, err := repo.SetPaid(1)
		assert.Nil(t, err)
	})
}

func TestSetCancel(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.SetCancel(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		_, err := repo.SetCancel(1)
		assert.Nil(t, err)
	})
}

func TestSetCheckInTime(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.SetCheckInTime(1, time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC))
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		a := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
		_, err := repo.SetCheckInTime(1, a)
		assert.Nil(t, err)
	})
}

func TestSetCheckOutTime(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)
	PR := paymentmethod.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		a := time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC)
		_, err := repo.SetCheckOutTime(1, a)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockR := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "asd",
			Price:       321,
			UserID:      1,
		}
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		mock := B.Books{
			CheckInReserved:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			CheckOutReserved: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			Phone:            "123",
			UserID:           1,
			RoomID:           1,
			PaymentMethodID:  1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		PR.Insert(mockPM)
		repo.Insert(mock)
		a := time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC)
		_, err := repo.SetCheckOutTime(1, a)
		assert.Nil(t, err)
	})
}
