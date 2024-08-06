package model

type Insurance struct {
	UUIDPrimaryKey
	CustomerID     string  `gorm:"type:varchar(100)" json:"customer_id"`
	ProviderName   string  `gorm:"type:varchar(100)" json:"provider_name"`
	Type           string  `gorm:"type:varchar(100)" json:"product_type"`
	Name           string  `gorm:"type:varchar(100)" json:"name"`
	Period         string  `gorm:"type:varchar(100)" json:"period"`
	Class          string  `gorm:"type:varchar(100)" json:"class"`
	NumberOffamily int     `gorm:"type:int" json:"number_of_family"`
	DiscountId     string  `gorm:"type:varchar(100)" json:"discount_id"`
	Price          float64 `gorm:"type:decimal(12)" json:"price"`
}

type IakInquiryBody struct {
	Commands string `json:"commands"`
	Hp       string `json:"hp"`
	Code     string `json:"code"`
	RefID    string `json:"ref_id"`
	Month    int    `json:"month"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
}

type BpjsIAKResponse struct {
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
			KodeCabang     string `json:"kode_cabang"`
			NamaCabang     string `json:"nama_cabang"`
			SisaPembayaran string `json:"sisa_pembayaran"`
			JumlahPeserta  string `json:"jumlah_peserta"`
		} `json:"desc"`
	} `json:"data"`
	Meta []interface{} `json:"meta"`
}

type IakPayBody struct {
	Commands string `json:"commands"`
	TrID     int    `json:"tr_id"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
	RefID    string `json:"ref_id"`
}
