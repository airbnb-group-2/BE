package paymentmethod

import (
	"group-project2/entities/book"

	"gorm.io/gorm"
)

type PaymentMethods struct {
	gorm.Model
	Name  string       `gorm:"type:varchar(100)"`
	Books []book.Books `gorm:"foreignKey:PaymentMethodID"`
}
