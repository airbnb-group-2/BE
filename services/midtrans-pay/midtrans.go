package midtranspay

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func CreateTransaction(core coreapi.Client) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "Pembayaran BCA Klik Pay",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "be6-003",
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

func Notification(w http.ResponseWriter, r *http.Request, c coreapi.Client) {
	var NotificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&NotificationPayload)
	if err != nil {
		log.Warn("terjadi error ketika decoding json", err)
		return
	}

	orderID, exists := NotificationPayload["order_id"].(string)
	if !exists {
		log.Warn("tidak ditemukan order_id")
		return
	}

	trxStatusResp, e := c.CheckTransaction(orderID)
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if trxStatusResp != nil {
			if trxStatusResp.TransactionStatus == "capture" {
				if trxStatusResp.FraudStatus == "challenge" {
					fmt.Println("status challenge")
				} else if trxStatusResp.FraudStatus == "accept" {
					fmt.Println("status accept")
				}
			} else if trxStatusResp.TransactionStatus == "settlement" {
				fmt.Println("status settlement")
			} else if trxStatusResp.TransactionStatus == "deny" {
				fmt.Println("status deny")
			} else if trxStatusResp.TransactionStatus == "cancel" || trxStatusResp.TransactionStatus == "expire" {
				fmt.Println("status cancel")
			} else if trxStatusResp.TransactionStatus == "pending" {
				fmt.Println("status pending")
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}
