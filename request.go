package ovo

import (
	"strconv"
	"time"
)

// CreateRequest is request model for create push-to-pay transaction.
type CreateRequest struct {
	Amount          float64 `validate:"required,gt=0"`
	Phone           string  `validate:"e164" mod:"no_space,e164"`
	MerchantInvoice string  `validate:"required" mod:"no_space"`
	ReferenceNumber int     `validate:"lt=1000000"`
	BatchNo         int     `validate:"lt=1000000"`
}

// CreateReversalRequest is request model for create reversal push-to-pay transaction.
type CreateReversalRequest struct {
	Amount          float64 `validate:"required,gt=0"`
	Phone           string  `validate:"e164" mod:"no_space,e164"`
	MerchantInvoice string  `validate:"required" mod:"no_space"`
	ReferenceNumber int     `validate:"lt=1000000"`
	BatchNo         int     `validate:"lt=1000000"`
}

// VoidRequest is request model for void push-to-pay transaction.
type VoidRequest struct {
	Amount          float64 `validate:"required,gt=0"`
	Phone           string  `validate:"e164" mod:"no_space,e164"`
	MerchantInvoice string  `validate:"required" mod:"no_space"`
	ReferenceNumber int     `validate:"lt=1000000"`
	BatchNo         int     `validate:"lt=1000000"`
}

// InquiryPhoneRequest is request model for inquiry phone number.
type InquiryPhoneRequest struct {
	Phone string `validate:"e164" mod:"no_space,e164"`
}

// GetStatusRequest is request model for get transaction data & status.
type GetStatusRequest struct {
	Amount          float64 `validate:"required,gt=0"`
	Phone           string  `validate:"e164" mod:"no_space,e164"`
	MerchantInvoice string  `validate:"required" mod:"no_space"`
	ReferenceNumber int     `validate:"lt=1000000"`
	BatchNo         int     `validate:"lt=1000000"`
}

type request struct {
	Type                   string                 `json:"type"`
	ProcessingCode         string                 `json:"processingCode"`
	Amount                 float64                `json:"amount,omitempty"`
	Date                   string                 `json:"date"`
	ReferenceNumber        string                 `json:"referenceNumber,omitempty"`
	TID                    string                 `json:"tid"`
	MID                    string                 `json:"mid"`
	MerchantID             string                 `json:"merchantId"`
	StoreCode              string                 `json:"storeCode"`
	AppSource              string                 `json:"appSource"`
	TransactionRequestData transactionRequestData `json:"transactionRequestData"`
}

type transactionRequestData struct {
	BatchNo         string `json:"batchNo,omitempty"`
	Phone           string `json:"phone"`
	MerchantInvoice string `json:"merchantInvoice,omitempty"`
}

func (c *Client) prepareCreateRequest(req CreateRequest) request {
	return request{
		Type:            "0200",
		ProcessingCode:  "040000",
		Amount:          req.Amount,
		Date:            time.Now().Format("2006-01-02 15:04:05.000"),
		ReferenceNumber: strconv.Itoa(req.ReferenceNumber),
		TID:             c.tid,
		MID:             c.mid,
		MerchantID:      c.merchantID,
		StoreCode:       c.storeCode,
		AppSource:       "POS",
		TransactionRequestData: transactionRequestData{
			BatchNo:         strconv.Itoa(req.BatchNo),
			Phone:           req.Phone,
			MerchantInvoice: req.MerchantInvoice,
		},
	}
}

func (c *Client) prepareCreateReversalRequest(req CreateReversalRequest) request {
	return request{
		Type:            "0400",
		ProcessingCode:  "040000",
		Amount:          req.Amount,
		Date:            time.Now().Format("2006-01-02 15:04:05.000"),
		ReferenceNumber: strconv.Itoa(req.ReferenceNumber),
		TID:             c.tid,
		MID:             c.mid,
		MerchantID:      c.merchantID,
		StoreCode:       c.storeCode,
		AppSource:       "POS",
		TransactionRequestData: transactionRequestData{
			BatchNo:         strconv.Itoa(req.BatchNo),
			Phone:           req.Phone,
			MerchantInvoice: req.MerchantInvoice,
		},
	}
}

func (c *Client) prepareVoidRequest(req VoidRequest) request {
	return request{
		Type:            "0200",
		ProcessingCode:  "020040",
		Amount:          req.Amount,
		Date:            time.Now().Format("2006-01-02 15:04:05.000"),
		ReferenceNumber: strconv.Itoa(req.ReferenceNumber),
		TID:             c.tid,
		MID:             c.mid,
		MerchantID:      c.merchantID,
		StoreCode:       c.storeCode,
		AppSource:       "POS",
		TransactionRequestData: transactionRequestData{
			BatchNo:         strconv.Itoa(req.BatchNo),
			Phone:           req.Phone,
			MerchantInvoice: req.MerchantInvoice,
		},
	}
}

func (c *Client) prepareInquiryPhoneRequest(req InquiryPhoneRequest) request {
	return request{
		Type:           "0100",
		ProcessingCode: "050000",
		Date:           time.Now().Format("2006-01-02 15:04:05.000"),
		TID:            c.tid,
		MID:            c.mid,
		MerchantID:     c.merchantID,
		StoreCode:      c.storeCode,
		AppSource:      "POS",
		TransactionRequestData: transactionRequestData{
			Phone: req.Phone,
		},
	}
}

func (c *Client) prepareGetStatusRequest(req GetStatusRequest) request {
	return request{
		Type:            "0100",
		ProcessingCode:  "040000",
		Amount:          req.Amount,
		Date:            time.Now().Format("2006-01-02 15:04:05.000"),
		ReferenceNumber: strconv.Itoa(req.ReferenceNumber),
		TID:             c.tid,
		MID:             c.mid,
		MerchantID:      c.merchantID,
		StoreCode:       c.storeCode,
		AppSource:       "POS",
		TransactionRequestData: transactionRequestData{
			BatchNo:         strconv.Itoa(req.BatchNo),
			Phone:           req.Phone,
			MerchantInvoice: req.MerchantInvoice,
		},
	}
}
