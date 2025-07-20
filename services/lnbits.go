package services

import (
	//"fmt"
	"github.com/go-resty/resty/v2"
)

var baseURL = "http://localhost:5000"

type LNBitsWallet struct {
	ID         string `json:"id"`
	AdminKey   string `json:"adminkey"`
	InvoiceKey string `json:"inkey"`
	Name       string `json:"name"`
}

type InvoiceResponse struct {
	PaymentRequest string `json:"payment_request"`
	PaymentHash    string `json:"payment_hash"`
	ExpiresAt      int64  `json:"expiry"`
}

func CreateWallet(username string) (*LNBitsWallet, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"user_name":   username,
			"wallet_name": "default",
		}).
		SetResult(&LNBitsWallet{}).
		Post(baseURL + "/wallet")

	if err != nil {
		return nil, err
	}

	wallet := resp.Result().(*LNBitsWallet)
	return wallet, nil
}

func GenerateInvoice(invoiceKey string, amount int64, memo string) (*InvoiceResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Api-Key", invoiceKey).
		SetBody(map[string]interface{}{
			"out":    false,
			"amount": amount,
			"memo":   memo,
			"expiry": 3600,
		}).
		SetResult(&InvoiceResponse{}).
		Post(baseURL + "/api/v1/payments")

	if err != nil {
		return nil, err
	}

	return resp.Result().(*InvoiceResponse), nil
}
