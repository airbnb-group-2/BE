package paymentmethod

import (
	"errors"
	P "group-project2/entities/payment-method"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		db: db,
	}
}

func (repo *PaymentMethodRepository) Insert(NewPaymentMethod P.PaymentMethods) (P.PaymentMethods, error) {
	if err := repo.db.Create(&NewPaymentMethod).Error; err != nil {
		log.Warn(err)
		return P.PaymentMethods{}, err
	}
	return NewPaymentMethod, nil
}

func (repo *PaymentMethodRepository) Get() ([]P.PaymentMethods, error) {
	paymentMethods := []P.PaymentMethods{}
	if RowsAffected := repo.db.Find(&paymentMethods).RowsAffected; RowsAffected == 0 {
		return nil, errors.New("belum ada payment method yang terdaftar")
	}
	return paymentMethods, nil
}

func (ur *PaymentMethodRepository) Delete(ID int) error {
	paymentMethod := P.PaymentMethods{}
	if RowsAffected := ur.db.Delete(&paymentMethod, ID).RowsAffected; RowsAffected == 0 {
		return errors.New("tidak ada payment method yang dihapus")
	}
	return nil
}
