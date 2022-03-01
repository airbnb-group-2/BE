package image

import (
	"group-project2/configs"
	B "group-project2/entities/book"
	I "group-project2/entities/image"
	PM "group-project2/entities/payment-method"
	Rat "group-project2/entities/rating"
	R "group-project2/entities/room"
	U "group-project2/entities/user"
	"group-project2/repositories/room"
	"group-project2/repositories/user"
	"group-project2/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

	t.Run("fail to insert room", func(t *testing.T) {
		mock := I.Images{}
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
		mock := I.Images{
			Link:   "http",
			RoomID: 1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)

		res, err := repo.Insert(mock)
		assert.Nil(t, err)
		assert.Equal(t, mock.Link, res.Link)
	})
}

func TestGetImagesByRoomID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetImagesByRoomID(1)
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
		mock := I.Images{
			Link:   "http",
			RoomID: 1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		repo.Insert(mock)

		_, err := repo.GetImagesByRoomID(1)
		assert.Nil(t, err)
	})
}

func TestGetImageByID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetImageByID(1)
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
		mock := I.Images{
			Link:   "http",
			RoomID: 1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		repo.Insert(mock)

		_, err := repo.GetImageByID(1)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		mock2 := I.Images{
			Model: gorm.Model{ID: 1},
			Link:  "https",
		}
		_, err := repo.Update(mock2)
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
		mock := I.Images{
			Link:   "http",
			RoomID: 1,
		}
		mock2 := I.Images{
			Model: gorm.Model{ID: 1},
			Link:  "https",
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		repo.Insert(mock)

		res, err := repo.Update(mock2)
		assert.Nil(t, err)
		assert.Equal(t, mock2.Link, res.Link)
	})
}

func TestDeleteImageByID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		err := repo.DeleteImageByID(1)
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
		mock := I.Images{
			Link:   "http",
			RoomID: 1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		repo.Insert(mock)

		err := repo.DeleteImageByID(1)
		assert.Nil(t, err)
	})
}

func TestDeleteImageByRoomID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)
	RR := room.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		err := repo.DeleteImageByRoomID(1)
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
		mock := I.Images{
			Link:   "http",
			RoomID: 1,
		}
		UR.Insert(mockUser)
		RR.Insert(mockR)
		repo.Insert(mock)

		err := repo.DeleteImageByRoomID(1)
		assert.Nil(t, err)
	})
}
