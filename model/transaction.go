package model

import (
	"time"

	"gorm.io/gorm"
)

const STATUS_UNPAID = "unpaid"
const STATUS_FAIL = "fail"
const STATUS_PROCESSING = "processing"
const STATUS_SUCCESSFUL = "successful"

const ADMIN_FEE = 2500

// ======== Product Type=========

type Transaction struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	UserID        string         `gorm:"index" json:"user_id"`
	Status        string         `gorm:"type:varchar(50)" json:"status"`
	ProductType   string         `gorm:"type:varchar(50)" json:"product_type"`
	ProductDetail interface{}    `gorm:"serializer:json" json:"product_detail"`
	Description   string         `gorm:"type:text" json:"description"`
	DiscountPrice float64        `gorm:"type:decimal(12)" json:"discount_price"`
	AdminFee      float64        `gorm:"type:decimal(12)" json:"admin_fee"`
	Price         float64        `gorm:"type:decimal(12)" json:"price"`
	TotalPrice    float64        `gorm:"type:decimal(12)" json:"total_price"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
type GetProductDetail struct {
	Status     string `json:"status"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	Period     string `json:"period"`
}

type TransactionPPD struct {
	Phone       string `json:"phone_number"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
	DiscountID  string `json:"discount_id"`
}

type TransactionWifi struct {
	Customer_Name string  `json:"customer_name"`
	Code          string  `json:"code"`
	Provider_Name string  `json:"provider_name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	DiscountId    string  `json:"discount_id"`
}

type TransactionTransfer struct {
	Phone  string `json:"phone_number"`
	UserID string `json:"user_id"`
	Note   string `json:"note"`
}

type PayloadMail struct {
	// Insurance
	Class          string `json:"class"`
	NumberOffamily int    `json:"number_of_family"`
	// Electricity
	Amount          string `json:"amount"`
	Token           string `json:"token"`
	ElectricalPower int    `json:"electrical_power"`
	// WIFI
	WifiBandwith int `json:"wifi_bandwith"`
	// PPD
	Phone string `json:"phone"`
	// =========

	CustomerName  string    `json:"name"`
	OrderId       string    `json:"order_id"`
	CustomerId    string    `json:"costumer_id"`
	ProductType   string    `json:"product_type"`
	Status        string    `json:"status"`
	RecipentEmail string    `json:"recipent_email"`
	ProviderName  string    `json:"provider_name"`
	Period        string    `json:"period"`
	Subject       string    `json:"subject"`
	TransactionAt time.Time `json:"transaction_at"`
	Description   string    `gorm:"type:text" json:"description"`
	DiscountPrice float64   `json:"discount_price"`
	AdminFee      float64   `json:"admin_fee"`
	Price         float64   `json:"price"`
	TotalPrice    float64   `json:"total_price"`
}

type TransactionCountInfo struct {
	ID      string  `json:"id"`
	Product string  `json:"product"`
	Price   float64 `json:"count_price"`
	Month   string  `json:"month"`
	Year    int     `json:"year"`
}
