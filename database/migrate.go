package database

import (
	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(

		&model.User{},
		&model.VaNumber{},
		&model.Bank{},
		&model.Insurance{},
		&model.Electricity{},
		&model.Pdam{},
		&model.Discount{},
		&model.Transaction{},
		&model.PulsaPaketData{},
		&model.Wifi{},
	)

	if err != nil {
		panic(err)
	}
}

func Drop(db *gorm.DB) {
	err := db.Migrator().DropTable(

		&model.User{},
		&model.VaNumber{},
		&model.Bank{},
		&model.Insurance{},
		&model.Electricity{},
		&model.Pdam{},
		&model.Discount{},
		&model.Transaction{},
		&model.PulsaPaketData{},
		&model.Wifi{},
	)
	if err != nil {
		panic(err)
	}
}
