package model

type Discount struct {
	UUIDPrimaryKey
	DiscountCode  string  `gorm:"type:varchar(100)" json:"discount_code"`
	Image         string  `gorm:"type:varchar(100)" json:"image"`
	Description   string  `gorm:"type:text" json:"description"`
	DiscountPrice float64 `gorm:"type:decimal(12)" json:"discount_price"`
}
