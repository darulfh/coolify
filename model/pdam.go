package model

type Pdam struct {
	UUIDPrimaryKey
	PartnerId    string  `gorm:"type:varchar(100)" json:"partner_id"`
	CustomerID   string  `gorm:"type:varchar(100)" json:"customer_id"`
	ProviderName string  `gorm:"type:varchar(100)" json:"provider_name"`
	Type         string  `gorm:"type:varchar(100)" json:"product_type"`
	Name         string  `gorm:"type:varchar(100)" json:"name"`
	Address      string  `gorm:"type:text" json:"address"`
	Period       string  `gorm:"type:varchar(100)" json:"period"`
	DiscountId   string  `gorm:"type:varchar(100)" json:"discount_id"`
	Price        float64 `gorm:"type:decimal(12)" json:"price"`
}
