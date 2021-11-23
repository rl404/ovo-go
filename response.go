package ovo

type transactionResponseData struct {
	OVOID            string `json:"ovoid"`
	FullName         string `json:"fullName"`
	StoreName        string `json:"storeName"`
	StoreCode        string `json:"storeCode"`
	StoreAddress1    string `json:"storeAddress1"`
	StoreAddress2    string `json:"storeAddress2"`
	CashUsed         string `json:"cashUsed"`
	CashBalance      string `json:"cashBalance"`
	OVOPointsEarned  string `json:"ovoPointsEarned"`
	OVOPointsUsed    string `json:"ovoPointsUsed"`
	OVOPointsBalance string `json:"ovoPointsBalance"`
	PaymentType      string `json:"paymentType"`
}
