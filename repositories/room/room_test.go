package room

import (
	"group-project2/configs"
	B "group-project2/entities/book"
	I "group-project2/entities/image"
	PM "group-project2/entities/payment-method"
	Rat "group-project2/entities/rating"
	R "group-project2/entities/room"
	U "group-project2/entities/user"
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

	t.Run("fail to insert room", func(t *testing.T) {
		mock := R.Rooms{}
		_, err := repo.Insert(mock)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mock := R.Rooms{
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
		UR.Insert(mockUser)
		res, err := repo.Insert(mock)
		assert.Nil(t, err)
		assert.Equal(t, mock.Name, res.Name)
	})
}

func TestGetAllRooms(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetAllRooms()
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mock := R.Rooms{
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
		UR.Insert(mockUser)
		repo.Insert(mock)
		_, err := repo.GetAllRooms()
		assert.Nil(t, err)
	})
}

func TestGetRoomByID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetRoomByID(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mock := R.Rooms{
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
		UR.Insert(mockUser)
		repo.Insert(mock)
		_, err := repo.GetRoomByID(1)
		assert.Nil(t, err)
	})
}

func TestGetRoomsByUserID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetRoomsByUserID(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mock := R.Rooms{
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
		UR.Insert(mockUser)
		repo.Insert(mock)
		_, err := repo.GetRoomsByUserID(1)
		assert.Nil(t, err)
	})
}

func TestGetRoomsByCity(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		_, err := repo.GetRoomsByCity("aaa")
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mock := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "sby",
			Price:       321,
			UserID:      1,
		}
		UR.Insert(mockUser)
		repo.Insert(mock)
		_, err := repo.GetRoomsByCity("sby")
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		UR.Insert(mockUser)
		mock := R.Rooms{
			Model: gorm.Model{ID: 0},
			Name:  "Deluxe",
		}
		_, err := repo.Update(mock)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mock := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "sby",
			Price:       321,
			UserID:      1,
		}
		mock2 := R.Rooms{
			Model:       gorm.Model{ID: 1},
			Description: "bbb",
		}
		repo.Insert(mock)
		_, err := repo.Update(mock2)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)
	UR := user.New(db)

	t.Run("fail to insert room", func(t *testing.T) {
		err := repo.Delete(1, 2)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mock := R.Rooms{
			Name:        "Deluxe",
			Description: "aaa",
			Guest:       3,
			Bedroom:     4,
			HasWifi:     true,
			HasAc:       true,
			HasKitchen:  true,
			Longitude:   "sad",
			Latitude:    "asd",
			City:        "sby",
			Price:       321,
			UserID:      1,
		}
		UR.Insert(mockUser)
		repo.Insert(mock)
		err := repo.Delete(1, 1)
		assert.Nil(t, err)
	})
}
