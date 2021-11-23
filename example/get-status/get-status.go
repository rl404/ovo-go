package main

import (
	"log"

	"github.com/rl404/ovo-go"
)

func main() {
	appID := "appID"
	key := "key123"
	tid := "123"
	mid := "123"
	merchantID := "123"
	storeCode := "ABC123"
	invoice := "invoice123"

	o := ovo.NewDefault(appID, key, tid, mid, merchantID, storeCode, ovo.Sandbox)

	tx, code, err := o.GetStatus(ovo.GetStatusRequest{
		Amount:          10000,
		Phone:           "081234567890",
		MerchantInvoice: invoice,
		ReferenceNumber: 1,
		BatchNo:         1,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, tx)
}
