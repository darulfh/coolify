package dto

type TransactionPPDDto struct {
	Type        string `json:"type"`
	ProductID   string `json:"product_id"`
	DiscountID  string `json:"discount_id"`
	PhoneNumber string `json:"phone_number"`
}

type TransactionTransferDto struct {
	PhoneNumber string  `json:"phone_number"`
	Amount      float64 `json:"amount"`
	Note        string  `json:"note"`
}
