package model

type Electricity struct {
	UUIDPrimaryKey
	CustomerId      string  `gorm:"type:varchar(100)" json:"customer_id"`
	ProviderName    string  `gorm:"type:varchar(100)" json:"provider_name"`
	Type            string  `gorm:"type:varchar(100)" json:"product_type"`
	Name            string  `gorm:"type:varchar(100)" json:"name"`
	Period          string  `gorm:"type:varchar(100)" json:"period"`
	Token           string  `gorm:"type:varchar(100)" json:"token"`
	ElectricalPower int     `gorm:"type:int" json:"electrical_power"`
	DiscountId      string  `gorm:"type:varchar(100)" json:"discount_id"`
	Price           float64 `gorm:"type:decimal(12)" json:"price"`
}

type IakPostPaidResponse struct {
	Data struct {
		TrID         int    `json:"tr_id"`
		Code         string `json:"code"`
		Hp           string `json:"hp"`
		TrName       string `json:"tr_name"`
		Period       string `json:"period"`
		Nominal      int    `json:"nominal"`
		Admin        int    `json:"admin"`
		RefID        string `json:"ref_id"`
		ResponseCode string `json:"response_code"`
		Message      string `json:"message"`
		Price        int    `json:"price"`
		SellingPrice int    `json:"selling_price"`
		Desc         struct {
			Tarif         string `json:"tarif"`
			Daya          int    `json:"daya"`
			LembarTagihan string `json:"lembar_tagihan"`
			Tagihan       struct {
				Detail []struct {
					Periode      string `json:"periode"`
					NilaiTagihan string `json:"nilai_tagihan"`
					Admin        string `json:"admin"`
					Denda        string `json:"denda"`
					Total        int    `json:"total"`
				} `json:"detail"`
			} `json:"tagihan"`
		} `json:"desc"`
	} `json:"data"`
	Meta []interface{} `json:"meta"`
}

type IakElectricityTokenInquiry struct {
	Data struct {
		Status       interface{} `json:"status"`
		CustomerID   string      `json:"customer_id"`
		MeterNo      string      `json:"meter_no"`
		SubscriberID string      `json:"subscriber_id"`
		Name         string      `json:"name"`
		SegmentPower string      `json:"segment_power"`
		Message      string      `json:"message"`
		Rc           string      `json:"rc"`
	} `json:"data"`
}
