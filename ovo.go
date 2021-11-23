package ovo

import (
	"net/http"
	"time"
)

// Client is ovo client.
type Client struct {
	appID      string
	key        string
	tid        string
	mid        string
	merchantID string
	storeCode  string
	baseURL    string
	requester  Requester
	logger     Logger
}

// Option is config for ovo client.
type Option struct {
	AppID      string
	Key        string
	TID        string
	MID        string
	MerchantID string
	StoreCode  string
	BaseURL    string
	Requester  Requester
	Logger     Logger
}

// New to create new ovo client with config.
func New(option Option) *Client {
	if option.Logger == nil {
		option.Logger = defaultLogger(LogError)
	}

	if option.Requester == nil {
		option.Requester = defaultRequester(&http.Client{
			Timeout: 70 * time.Second,
		}, option.Logger)
	}

	return &Client{
		appID:      option.AppID,
		key:        option.Key,
		tid:        option.TID,
		mid:        option.MID,
		merchantID: option.MerchantID,
		storeCode:  option.StoreCode,
		baseURL:    option.BaseURL,
		requester:  option.Requester,
		logger:     option.Logger,
	}
}

// NewDefault to create new ovo client with default config.
func NewDefault(appID, key, tid, mid, merchantID, storeCode string, env EnvironmentType) *Client {
	return New(Option{
		AppID:      appID,
		Key:        key,
		TID:        tid,
		MID:        mid,
		MerchantID: merchantID,
		StoreCode:  storeCode,
		BaseURL:    envURL[env],
		Requester: defaultRequester(&http.Client{
			Timeout: 70 * time.Second,
		}, defaultLogger(envLog[env])),
		Logger: defaultLogger(envLog[env]),
	})
}
