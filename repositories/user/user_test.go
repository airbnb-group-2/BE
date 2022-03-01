package user

import (
	"group-project2/configs"
	B "group-project2/entities/book"
	I "group-project2/entities/image"
	PM "group-project2/entities/payment-method"
	Rat "group-project2/entities/rating"
	R "group-project2/entities/room"
	U "group-project2/entities/user"
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

	t.Run("fail to get user", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		mockUser2 := U.Users{
			Name:     "Ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		repo.Insert(mockUser2)
		_, err := repo.Insert(mockUser)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		res, err := repo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})
}

func TestGet(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("fail to get user", func(t *testing.T) {
		_, err := repo.GetUserByID(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to get users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		repo.Insert(mockUser)

		res, err := repo.GetUserByID(1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})
}

func TestUpdate(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("succeed to update users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		mockUser2 := U.Users{
			Model: gorm.Model{ID: 1},
			Name:  "Ucup2",
			Email: "ucup2@ucup.com",
		}
		repo.Insert(mockUser)

		res, err := repo.Update(mockUser2)
		assert.Nil(t, err)
		assert.Equal(t, mockUser2.Name, res.Name)
		assert.Equal(t, mockUser2.Email, res.Email)
	})

	t.Run("fail to update user", func(t *testing.T) {
		mockUser1 := U.Users{
			Name:     "Ucup3",
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}
		mockUser2 := U.Users{
			Model:    gorm.Model{ID: 3},
			Email:    "ucup3@ucup.com",
			Password: "ucup123",
		}

		repo.Insert(mockUser1)

		_, err := repo.Update(mockUser2)
		assert.NotNil(t, err)
	})
}

func TestSetRenter(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("succeed to set renter users", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		repo.Insert(mockUser)

		_, err := repo.SetRenter(1)
		assert.Nil(t, err)
	})

	t.Run("fail to set renter user", func(t *testing.T) {
		_, err := repo.SetRenter(2)
		assert.NotNil(t, err)
	})
}

func TestDeleteByID(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("fail to delete user", func(t *testing.T) {
		err := repo.DeleteByID(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to delete user", func(t *testing.T) {
		mockUser := U.Users{
			Name:     "Ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		}
		repo.Insert(mockUser)

		err := repo.DeleteByID(1)
		assert.Nil(t, err)
	})
}
