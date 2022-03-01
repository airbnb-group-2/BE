package paymentmethod

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
)

var (
	config = configs.GetConfig(true)
	db     = utils.InitDB(config)
)

func TestInsert(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("fail to insert PM", func(t *testing.T) {
		mockPM := PM.PaymentMethods{}
		repo.Insert(mockPM)
		_, err := repo.Insert(mockPM)
		assert.NotNil(t, err)
	})

	t.Run("succeed to insert PM", func(t *testing.T) {
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}
		res, err := repo.Insert(mockPM)
		assert.Nil(t, err)
		assert.Equal(t, mockPM.Name, res.Name)
	})
}

func TestGet(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("fail to get PM", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})

	t.Run("succeed to get PM", func(t *testing.T) {
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}

		repo.Insert(mockPM)

		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, res[0].Name, mockPM.Name)
	})
}

func TestDelete(t *testing.T) {
	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

	repo := New(db)

	t.Run("fail to delete PM", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("succeed to insert PM", func(t *testing.T) {
		mockPM := PM.PaymentMethods{
			Name: "gopay",
		}

		repo.Insert(mockPM)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
