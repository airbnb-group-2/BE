package paymentmethod

import PM "group-project2/entities/payment-method"

type RequestCreatePM struct {
	Name string `json:"name" form:"name"`
}

func (Req RequestCreatePM) ToEntityPM() PM.PaymentMethods {
	return PM.PaymentMethods{
		Name: Req.Name,
	}
}

type ResponseCreatePM struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToResponseCreatePM(PM PM.PaymentMethods) ResponseCreatePM {
	return ResponseCreatePM{
		ID:   PM.ID,
		Name: PM.Name,
	}
}

type ResponseGet struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToResponseGet(PMs []PM.PaymentMethods) []ResponseGet {
	Responses := make([]ResponseGet, len(PMs))

	for i := 0; i < len(PMs); i++ {
		Responses[i].ID = PMs[i].ID
		Responses[i].Name = PMs[i].Name
	}

	return Responses
}
