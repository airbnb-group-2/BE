package paymentmethod

import P "group-project2/entities/payment-method"

type PaymentMethod interface {
	Insert(NewPaymentMethod P.PaymentMethods) (P.PaymentMethods, error)
	Get() ([]P.PaymentMethods, error)
	Delete(ID uint) error
}
