package midtranspay

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var (
	MIDTRANS_KEY = os.Getenv("MIDTRANS_KEY")
)

func InitConnection() coreapi.Client {
	client := coreapi.Client{}
	client.New(MIDTRANS_KEY, midtrans.Sandbox)
	return client
}

func CreateTransaction(core coreapi.Client) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "Pembayaran BCA Klik Pay",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "be6-001",
			GrossAmt: 50000,
		},
	}

	apiRes, err := core.ChargeTransaction(req)
	if err != nil {
		log.Warn("payment error:", err)
	}
	fmt.Println("API Result:", apiRes)
	return apiRes
}
