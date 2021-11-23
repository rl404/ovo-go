package ovo

import (
	"context"
	"net/http"
)

// Transaction is ovo transaction model.
type Transaction struct {
	Type                    string                  `json:"type"`
	ProcessingCode          string                  `json:"processingCode"`
	Amount                  float64                 `json:"amount"`
	Date                    string                  `json:"date"`
	TraceNumber             int                     `json:"traceNumber"`
	HostTime                string                  `json:"hostTime"`
	HostDate                string                  `json:"hostDate"`
	ReferenceNumber         int                     `json:"referenceNumber"`
	ApprovalCode            string                  `json:"approvalCode"`
	ResponseCode            string                  `json:"responseCode"`
	TID                     string                  `json:"tid"`
	MID                     string                  `json:"mid"`
	TransactionRequestData  transactionRequestData  `json:"transactionRequestData"`
	TransactionResponseData transactionResponseData `json:"transactionResponseData"`
}

// Create to create push-to-pay transaction.
func (c *Client) Create(request CreateRequest) (*Transaction, int, error) {
	return c.CreateWithContext(context.Background(), request)
}

// CreateWithContext to create push-to-pay transaction with context.
func (c *Client) CreateWithContext(ctx context.Context, request CreateRequest) (*Transaction, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response Transaction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		c.baseURL,
		c.appID,
		c.key,
		nil,
		c.prepareCreateRequest(request),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response, code, nil
}

// CreateReversal to create reversal push-to-pay transaction.
func (c *Client) CreateReversal(request CreateReversalRequest) (*Transaction, int, error) {
	return c.CreateReversalWithContext(context.Background(), request)
}

// CreateReversalWithContext to create reversal push-to-pay transaction with context.
func (c *Client) CreateReversalWithContext(ctx context.Context, request CreateReversalRequest) (*Transaction, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response Transaction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		c.baseURL,
		c.appID,
		c.key,
		nil,
		c.prepareCreateReversalRequest(request),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response, code, nil
}

// Void to void push-to-pay transaction.
func (c *Client) Void(request VoidRequest) (*Transaction, int, error) {
	return c.VoidWithContext(context.Background(), request)
}

// VoidWithContext to void push-to-pay transaction with context.
func (c *Client) VoidWithContext(ctx context.Context, request VoidRequest) (*Transaction, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response Transaction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		c.baseURL,
		c.appID,
		c.key,
		nil,
		c.prepareVoidRequest(request),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response, code, nil
}

// InquiryPhone to inquiry customer phone number.
func (c *Client) InquiryPhone(request InquiryPhoneRequest) (*Transaction, int, error) {
	return c.InquiryPhoneWithContext(context.Background(), request)
}

// InquiryPhoneWithContext to inquiry customer phone number with context.
func (c *Client) InquiryPhoneWithContext(ctx context.Context, request InquiryPhoneRequest) (*Transaction, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response Transaction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		c.baseURL,
		c.appID,
		c.key,
		nil,
		c.prepareInquiryPhoneRequest(request),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response, code, nil
}


// GetStatus to get transaction data & status.
func (c *Client) GetStatus(request GetStatusRequest) (*Transaction, int, error) {
	return c.GetStatusWithContext(context.Background(), request)
}

// GetStatusWithContext to get transaction data & status with context.
func (c *Client) GetStatusWithContext(ctx context.Context, request GetStatusRequest) (*Transaction, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response Transaction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		c.baseURL,
		c.appID,
		c.key,
		nil,
		c.prepareGetStatusRequest(request),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return &response, code, nil
}
