package model

const PULSA_TYPE = "pulsa"
const PAKET_DATA_TYPE = "data"

type PulsaPaketData struct {
	UUIDPrimaryKey
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Code        string  `json:"code"`
	Provider    string  `json:"provider"`
	Price       float64 `json:"price"`
	IsActive    *bool   `json:"is_active" gorm:"default:true"`
	Description string  `json:"description"`
}

type PPDResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Code        string  `json:"code"`
	Provider    string  `json:"provider"`
	Price       float64 `json:"price"`
	IsActive    *bool   `json:"is_active"`
	Description string  `json:"description"`
}

type PrePaidIakBody struct {
	// username IAK
	Username string `json:"username"`

	// phone number
	CustomerID string `json:"customer_id"`

	// Product code for pulsa or paket data
	ProductCode string `json:"product_code"`

	// ID transaction
	RefID string `json:"ref_id"`

	// Api key IAK
	Sign string `json:"sign"`
}

type PrePaidIakResponse struct {
	Data struct {
		RefID       string `json:"ref_id"`
		Status      int    `json:"status"`
		ProductCode string `json:"product_code"`
		CustomerID  string `json:"customer_id"`
		Price       int    `json:"price"`
		Message     string `json:"message"`
		Balance     int    `json:"balance"`
		TrID        int    `json:"tr_id"`
		Rc          string `json:"rc"`
	} `json:"data"`
}
