package model

type Wifi struct {
	UUIDPrimaryKey
	Name         string  `gorm:"type:varchar(100)" json:"name"`
	CustomerID   string  `gorm:"type:varchar(100)" json:"customer_id"`
	ProviderName string  `gorm:"type:varchar(100)" json:"provider_name"`
	ProductType  string  `gorm:"type:varchar(100)" json:"product_type"`
	Period       string  `gorm:"type:varchar(100)" json:"period"`
	Code         string  `gorm:"type:varchar(100)" json:"code"`
	DiscountId   string  `gorm:"type:varchar(100)" json:"discount_id"`
	WifiBandwith int     `gorm:"type:int" json:"wifi_bandwith"`
	Price        float64 `gorm:"type:decimal(12)" json:"price"`
}
