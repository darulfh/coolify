package routes

import (
	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/controller"
	"github.com/darulfh/skuy_pay_be/repository"
	insurance "github.com/darulfh/skuy_pay_be/usecase/Insurance"
	"github.com/darulfh/skuy_pay_be/usecase/auth"
	"github.com/darulfh/skuy_pay_be/usecase/balance"
	"github.com/darulfh/skuy_pay_be/usecase/bank"
	"github.com/darulfh/skuy_pay_be/usecase/discount"
	"github.com/darulfh/skuy_pay_be/usecase/electricity"
	"github.com/darulfh/skuy_pay_be/usecase/pdam"
	pulsa "github.com/darulfh/skuy_pay_be/usecase/pulsa_paket_data"
	"github.com/darulfh/skuy_pay_be/usecase/transaction"
	"github.com/darulfh/skuy_pay_be/usecase/users"
	"github.com/darulfh/skuy_pay_be/usecase/wifi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {

	// Auth user

	authRepository := repository.NewAuthRepository(db)
	authUseCase := auth.NewAuthUsecase(authRepository)
	authController := controller.NewAuthController(authUseCase)

	// IAK
	iakRepository := repository.NewIakApiRepository()

	// Transaction
	transactionRepository := repository.NewTransactionRepository(db)
	transactionUseCase := transaction.NewTransactionUsecase(transactionRepository)
	transactionController := controller.NewTransactionController(transactionUseCase)

	// Users
	userRepository := repository.NewUserRepository(db)
	userUseCase := users.NewUserUsecase(userRepository, transactionRepository)
	userController := controller.NewUserController(userUseCase)

	// Discount
	discountRepository := repository.NewDiscountRepository(db)
	discountUseCase := discount.NewDiscountUseCase(discountRepository)
	discountController := controller.NewDiscountController(discountUseCase)

	// Pulsa Paket Data
	ppdRepository := repository.NewPulsaPaketDataRepository(db)
	ppdUsecase := pulsa.NewPulsaPaketDataUsecase(ppdRepository, userRepository, transactionRepository, discountRepository, iakRepository)
	ppdController := controller.NewPulsaPaketDataController(ppdUsecase)

	// Balance
	virtualAgregatorOyApi := repository.NewVirtualAgregatorOyApiRepository()
	balanceRepository := repository.NewBalanceRepository(db)
	balanceUseCase := balance.NewBalanceUsecase(balanceRepository, userRepository, transactionRepository, virtualAgregatorOyApi)
	balanceController := controller.NewBalanceController(balanceUseCase)

	// Payment
	bankRepository := repository.NewBankRepository(db)
	bankUseCase := bank.NewbankUseCase(bankRepository)
	bankController := controller.NewBankController(bankUseCase)

	// PDAM
	billerRepository := repository.NewBillerOyApiOyApiRepository()
	pdamRepository := repository.NewPdamRepository(db)
	pdamUseCase := pdam.NewPdamUseCase(pdamRepository, userRepository, discountRepository, transactionRepository, billerRepository)
	pdamController := controller.NewPdamController(pdamUseCase)

	// Wifi
	wifiRepository := repository.NewWifiRepository(db)
	wifiUsecase := wifi.NewWifiUseCase(wifiRepository, userRepository, discountRepository, transactionRepository, billerRepository)
	wifiController := controller.NewWifiController(wifiUsecase)

	// INSURANCE
	insuranceRepository := repository.NewInsuranceRepository(db)
	insuranceUseCase := insurance.NewInsuranceUseCase(insuranceRepository, userRepository, discountRepository, transactionRepository, billerRepository, iakRepository)
	insuranceController := controller.NewInsuranceController(insuranceUseCase)

	// ELECTRICITY
	electricityRepository := repository.NewElectricityRepository(db)
	electricityUseCase := electricity.NewElectricityUseCase(electricityRepository, userRepository, discountRepository, transactionRepository, billerRepository, iakRepository)
	electricityController := controller.NewElectricityController(electricityUseCase)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, `
	// 		<h1>Welcome to PPOB APP</h1>
	// 	`)
	// })

	// url
	url := e.Group("/api/v1")

	//Patner CallBack
	url.POST("/patner/callback", balanceController.CreateBalanceController)

	// AUTH
	url.POST("/login", authController.LoginController)
	url.POST("/register", authController.RegisterController)
	url.POST("/admin/register", authController.RegisterAdminController)

	// User
	url.GET("/user/email", userController.GetUserByEmailController)
	url.POST("/user/pin/:id", userController.CheckPinForgotPasswordController)
	url.POST("/user/password/forgot/:id", userController.UpdateForgotPasswordController)

	all := e.Group("/api/v1")
	user := e.Group("/api/v1")
	admin := e.Group("/api/v1/admin")

	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(config.AppConfig.SecretJWT),
	}

	// ====== ADMIN ROLE =======
	admin.Use(middleware.JWTWithConfig(jwtConfig))

	// users
	admin.GET("/users", userController.GetAllUsersController)
	admin.GET("/users/query", userController.GetUsersByQueryController)
	admin.GET("/user/:id", userController.AdminGetUserByIdController)

	// Transaction routes
	admin.GET("/transactions/product/", transactionController.GetTransactionByProductController)
	admin.GET("/transactions/search/", transactionController.GetTransactionByQueryController)
	admin.GET("/transactions/status/search/", transactionController.GetTransactionByStatusQueryController)
	admin.GET("/transactions", transactionController.GetAllTransactionsController)
	admin.GET("/transactions/price/product/count", transactionController.GetTransactionsPriceCountController)
	admin.GET("/transactions/price/month/count", transactionController.GetTransactionsPriceByMonthController)
	admin.GET("/transactions/user/:id", transactionController.AdminGetTransactionByUserIdController)

	// Bank
	admin.POST("/bank", bankController.CreateBankController)
	admin.PUT("/bank/:id", bankController.UpdateBankController)
	admin.DELETE("/bank/:id", bankController.DeleteBankByIdController)

	//Discount
	admin.POST("/discount", discountController.CreateDiscountController)
	admin.PUT("/discount/:id", discountController.UpdateDiscountController)
	admin.DELETE("/discount/:id", discountController.DeleteDiscountByIdController)

	//PDAM
	admin.POST("/pdam", pdamController.CreatePdamController)
	admin.PUT("/pdam/:id", pdamController.UpdatePdamController)
	admin.DELETE("/pdam/:id", pdamController.DeletePdamByIdController)

	//Insurance
	admin.POST("/insurance", insuranceController.CreateInsuranceController)
	admin.PUT("/insurance/:id", insuranceController.UpdateInsuranceController)
	admin.DELETE("/insurance/:id", insuranceController.DeleteInsuranceByIdController)

	// Electricity
	admin.POST("/electricity", electricityController.CreateElectricityController)
	admin.PUT("/electricity/:id", electricityController.UpdateElectricityController)
	admin.DELETE("/electricity/:id", electricityController.DeleteElectricityByIdController)

	// pulsa paket data
	admin.POST("/ppd", ppdController.CreatePulsaPaketData)
	admin.GET("/ppd", ppdController.GetPPDByAdmin)
	admin.GET("/ppd/:id", ppdController.GetPPDByID)
	admin.PUT("/ppd/:id", ppdController.UpdatePPDById)
	admin.DELETE("/ppd/:id", ppdController.DeletePPDById)

	//wifi
	admin.POST("/wifi", wifiController.CreateWifiController)
	admin.PUT("/wifi/:id", wifiController.UpdateWifiController)
	admin.DELETE("/wifi/:id", wifiController.DeleteWifiByIdController)

	// ====== USER ROLE =======
	user.Use(middleware.JWTWithConfig(jwtConfig))
	user.GET("/profile", userController.GetUserByIdController)
	user.PUT("/user", userController.UpdateUserByIDController)
	user.PUT("/user/image", userController.UpdateUserImageByIDController)
	user.PUT("/user/password", userController.UpdatePasswordController)
	user.PUT("/user/pin", userController.CreatePINController)
	user.PUT("/user/pin/update", userController.UpdatePINController)
	user.POST("/user/pin", userController.CheckPINController)
	user.POST("/user/transfer/amount", userController.TransferAmountController)
	user.DELETE("/user", userController.DeleteUserByIDController)

	// pulsa paket data
	user.GET("/user/ppd", ppdController.GetPPDByUser)
	user.POST("/user/ppd", ppdController.CreateTransactionPPDUser)

	// transaction
	user.GET("/user/transactions/", transactionController.GetTransactionByUserIdController)

	// ====== ALL ROLE =======
	all.Use(middleware.JWTWithConfig(jwtConfig))

	// transaction
	all.GET("/transaction/:id", transactionController.GetTransactionByIdController)
	// Bank
	all.GET("/banks", bankController.GetAllBankController)
	all.GET("/bank/:id", bankController.GetBankByIdController)

	// Amount routes
	all.POST("/amount", balanceController.GenerateVaController)
	all.GET("/amount/status/:id", balanceController.GetPayBalanceStatusController)

	// Discount
	all.GET("/discounts", discountController.GetAllDiscountController)
	all.GET("/discount/:id", discountController.GetDiscountByIdController)
	all.POST("/discount/code", discountController.GetDiscountByCodeController)

	//Insurance
	all.GET("/insurances", insuranceController.GetAllInsuranceController)
	all.GET("/insurance/:id", insuranceController.GetInsuranceByIdController)
	all.POST("/insurance/inquiry", insuranceController.BillInquiryInsuranceController)
	all.POST("/insurance/pay", insuranceController.PayBillInquiryInsuranceController)
	all.POST("/bpjs/pay", insuranceController.BpjsPayIakUseCase)
	all.POST("/bpjs/inquiry", insuranceController.BpjsInquiryIakUseCase)

	// Electricity
	all.GET("/electricitys", electricityController.GetAllElectricityController)
	all.GET("/electricity/:id", electricityController.GetElectricityByIdController)
	all.POST("/electricity/postpaid/pay", electricityController.PayBillInquiryElectricityController)
	// Electricity PostPaid (Tagihan)
	all.POST("/electricity/postpaid/inquiry", electricityController.BillInquiryPostPaidElectricityController)
	// Electricity PrePaid (Token)
	all.POST("/electricity/prepaid/inquiry", electricityController.BillInquiryPrePaidElectricityController)

	/// IAK Tagihan Listrik
	all.POST("/electricitybill/pay", electricityController.ElectricityBillPayIakUseCase)
	all.POST("/electricitybill/inquiry", electricityController.ElectricityBillInquiryIakUseCase)

	/// IAK Token Listrik
	all.POST("/electricitytoken/pay", electricityController.ElectricityTokenPayIakUseCase)
	all.POST("/electricitytoken/inquiry", electricityController.ElectricityTokenInquiryIakUseCase)

	// PDAM
	all.GET("/pdams", pdamController.GetAllPdamController)
	all.GET("/pdam/:id", pdamController.GetPdamByIdController)
	all.POST("/pdam/inquiry", pdamController.BillInquiryPdamController)
	all.POST("/pdam/pay", pdamController.PayBillInquiryPdamController)

	// WIFI
	all.GET("/wifis", wifiController.GetAllWifiController)
	all.GET("/wifi/:id", wifiController.GetWifiByIdController)
	all.POST("/wifi/inquiry", wifiController.BillInquiryWifiController)
	all.POST("/wifi/pay", wifiController.PayBillWifiController)
}
