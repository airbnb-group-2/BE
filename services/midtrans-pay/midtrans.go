package midtranspay

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitConnection() coreapi.Client {
	if err := godotenv.Load("local.env"); err != nil {
		log.Info(err)
	}
	MIDTRANS_KEY := os.Getenv("MIDTRANS_KEY")

	client := coreapi.Client{}
	client.New(MIDTRANS_KEY, midtrans.Sandbox)
	return client
}

func CreateTransaction(core coreapi.Client, OrderID uint) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "Pembayaran BCA Klik Pay",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("be-6-%d", OrderID),
			GrossAmt: 50000,
		},
	}

	apiRes, err := core.ChargeTransaction(req)
	if err != nil {
		log.Warn("payment error:", err)
	}
	// fmt.Println("API Result:", apiRes)
	return apiRes
}

func Notification(c coreapi.Client, OrderID uint) (string, error) {
	trxStatusResp, e := c.CheckTransaction(fmt.Sprintf("be-6-%d", OrderID))
	res := ""
	if e != nil {
		log.Warn(e)
		return "", errors.New("internal server error")
	} else {
		if trxStatusResp != nil {
			if trxStatusResp.TransactionStatus == "capture" {
				if trxStatusResp.FraudStatus == "challenge" {
					res = "status challenge"
				} else if trxStatusResp.FraudStatus == "accept" {
					res = "status accept"
				}
			} else if trxStatusResp.TransactionStatus == "settlement" {
				res = "status settlement"
			} else if trxStatusResp.TransactionStatus == "deny" {
				res = "status deny"
			} else if trxStatusResp.TransactionStatus == "cancel" || trxStatusResp.TransactionStatus == "expire" {
				res = "status cancel"
			} else if trxStatusResp.TransactionStatus == "pending" {
				res = "status pending"
			}
		}
	}
	// fmt.Println("Trx Status:", trxStatusResp.TransactionStatus)
	return res, nil
}
