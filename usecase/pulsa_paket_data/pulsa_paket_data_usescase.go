package pulsa

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/dto"
	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"

	"github.com/google/uuid"
)

type PulsaPaketDataUsecase interface {
	CreatePulsaPaketData(data model.PulsaPaketData) (model.PPDResponse, error)
	GetAllPulsaPaketData(data dto.PulsaDto, isUser *bool) ([]model.PPDResponse, error)
	GetPulsaPaketDataById(id string) (model.PPDResponse, error)
	UpdatePulsaById(id string, data model.PulsaPaketData) (model.PPDResponse, error)
	DeletePulsaById(id string) error
	CreateTransactionPPD(userID string, payload dto.TransactionPPDDto) (*model.Transaction, error)
}

type pulsaPaketDataUsecase struct {
	ppdRepository         repository.PulsaPaketDataRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
	discountRepository    repository.DiscountRepository
	iakRepository         repository.IakApiRepository
}

func NewPulsaPaketDataUsecase(ppdRepository repository.PulsaPaketDataRepository, userRepository repository.UserRepository, transactionRepository repository.TransactionRepository, discountRepository repository.DiscountRepository, iakApi repository.IakApiRepository) *pulsaPaketDataUsecase {
	return &pulsaPaketDataUsecase{
		ppdRepository:         ppdRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
		discountRepository:    discountRepository,
		iakRepository:         iakApi,
	}
}

func (u *pulsaPaketDataUsecase) CreatePulsaPaketData(data model.PulsaPaketData) (model.PPDResponse, error) {
	ppd, err := u.ppdRepository.CreatePulsaPaketData(data)

	if err != nil {
		return model.PPDResponse{}, err
	}

	response := model.PPDResponse{
		ID:          ppd.ID,
		Name:        ppd.Name,
		Code:        ppd.Code,
		Type:        ppd.Type,
		Provider:    ppd.Provider,
		Price:       ppd.Price,
		IsActive:    ppd.IsActive,
		Description: ppd.Description,
	}

	return response, nil
}

func (u *pulsaPaketDataUsecase) GetAllPulsaPaketData(data dto.PulsaDto, isUser *bool) ([]model.PPDResponse, error) {
	var provider string
	if isUser != nil {
		provider = getProviderByPhone(data.PhoneNumber)
		if provider == "" && data.Provider == "" {
			return []model.PPDResponse{}, errors.New("provider is not supported")
		}
	}

	if provider != "" {
		data.Provider = provider
	}

	ppd, err := u.ppdRepository.GetAllPulsaPaketData(data, isUser)

	if err != nil {
		return []model.PPDResponse{}, err
	}
	response := make([]model.PPDResponse, 0)

	for _, v := range ppd {
		response = append(response, model.PPDResponse{
			ID:          v.ID,
			Name:        v.Name,
			Code:        v.Code,
			Type:        v.Type,
			Provider:    v.Provider,
			Price:       v.Price,
			IsActive:    v.IsActive,
			Description: v.Description,
		})
	}

	return response, nil
}

func (u *pulsaPaketDataUsecase) GetPulsaPaketDataById(id string) (model.PPDResponse, error) {
	ppd, err := u.ppdRepository.GetPulsaPaketDataById(id)
	if err != nil {
		return model.PPDResponse{}, err
	}

	response := model.PPDResponse{
		ID:          ppd.ID,
		Name:        ppd.Name,
		Code:        ppd.Code,
		Type:        ppd.Type,
		Provider:    ppd.Provider,
		Price:       ppd.Price,
		IsActive:    ppd.IsActive,
		Description: ppd.Description,
	}

	return response, nil

}

func (u *pulsaPaketDataUsecase) UpdatePulsaById(id string, data model.PulsaPaketData) (model.PPDResponse, error) {
	if err := u.ppdRepository.UpdatePulsaById(id, data); err != nil {
		return model.PPDResponse{}, err
	}

	ppd, err := u.GetPulsaPaketDataById(id)
	if err != nil {
		return model.PPDResponse{}, err
	}
	return ppd, err
}

func (u *pulsaPaketDataUsecase) DeletePulsaById(id string) error {
	if err := u.ppdRepository.DeletePulsaById(id); err != nil {
		return err
	}

	return nil
}

func (uc *pulsaPaketDataUsecase) CreateTransactionPPD(userID string, payload dto.TransactionPPDDto) (*model.Transaction, error) {
	ppd, err := uc.ppdRepository.GetPulsaPaketDataById(payload.ProductID)
	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error ppd id not found")
	}

	discount, _ := uc.discountRepository.GetDiscountByIdRepository(payload.DiscountID)

	td := model.TransactionPPD{
		Phone:       payload.PhoneNumber,
		Name:        ppd.Name,
		Code:        ppd.Code,
		Provider:    ppd.Provider,
		Description: ppd.Description,
		DiscountID:  discount.ID,
	}

	transaction := &model.Transaction{
		ID:            uuid.New().String(),
		UserID:        userID,
		Status:        model.STATUS_SUCCESSFUL,
		ProductType:   ppd.Type,
		Description:   fmt.Sprintf("Pembelian Pulsa Paket Data  %s ", ppd.Type),
		AdminFee:      0,
		Price:         ppd.Price,
		TotalPrice:    ppd.Price + 0 - discount.DiscountPrice,
		ProductDetail: td,
		DiscountPrice: discount.DiscountPrice,
	}

	user, err := uc.userRepository.GetUserByIDRepository(userID)

	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error user id not found")
	}

	if user.Amount < transaction.TotalPrice {
		return &model.Transaction{}, errors.New("your balance is not enough")
	}

	iakReq := model.PrePaidIakBody{
		Username:    config.AppConfig.UsernameIak,
		CustomerID:  payload.PhoneNumber,
		ProductCode: ppd.Code,
		RefID:       uuid.New().String(),
	}

	fmt.Printf("%+v\n", iakReq)

	iakResp, err := uc.iakRepository.IakTopUpPayRepository(&iakReq)

	fmt.Printf("%+v\n", iakResp)

	transaction.ID = iakResp.Data.RefID

	if err != nil {
		return &model.Transaction{}, fmt.Errorf("failed transaction %s", err)
	}

	ts, err := uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)

	if err != nil {
		return &model.Transaction{}, fmt.Errorf("failed transaction %s", err)
	}

	user.Amount -= transaction.TotalPrice
	_, err = uc.userRepository.UpdateUserAmountByIDRepository(userID, user)

	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error update amount")
	}

	// mailsend := model.PayloadMail{
	// 	OrderId:       transaction.ID,
	// 	CustomerName:  user.Name,
	// 	Status:        ts.Status,
	// 	ProductType:   "PPD",
	// 	RecipentEmail: user.Email,
	// 	Phone:         td.Phone,
	// 	TransactionAt: ts.UpdatedAt,
	// 	AdminFee:      0,
	// 	Description:   transaction.Description,
	// 	Price:         transaction.TotalPrice,
	// 	TotalPrice:    transaction.TotalPrice,
	// }

	// mail.SendingMail(mailsend)

	return ts, nil
}

func getProviderByPhone(phone string) string {
	if len(phone) < 4 {
		return ""
	}
	codeProvider := phone[0:4]
	switch codeProvider {
	case "0852":
		return "AS"
	case "0853":
		return "AS"
	case "0823":
		return "AS"
	case "0851":
		return "AS"
	case "0811":
		return "Halo"
	case "0812":
		return "Telkomsel"
	case "0813":
		return "Telkomsel"
	case "0821":
		return "Telkomsel"
	case "0822":
		return "Telkomsel"
	case "0814":
		return "Indosat"
	case "0815":
		return "Indosat"
	case "0816":
		return "Indosat"
	case "0855":
		return "Indosat"
	case "0856":
		return "Indosat"
	case "0857":
		return "Indosat"
	case "0858":
		return "Indosat"
	case "0817":
		return "XL"
	case "0818":
		return "XL"
	case "0859":
		return "XL"
	case "0877":
		return "XL"
	case "0878":
		return "XL"
	case "0838":
		return "Axis"
	case "0831":
		return "Axis"
	case "0832":
		return "Axis"
	case "0833":
		return "Axis"
	case "0895":
		return "Three"
	case "0896":
		return "Three"
	case "0897":
		return "Three"
	case "0898":
		return "Three"
	case "0899":
		return "Three"
	case "0881":
		return "Smatfren"
	case "0882":
		return "Smatfren"
	case "0883":
		return "Smatfren"
	case "0884":
		return "Smatfren"
	case "0885":
		return "Smatfren"
	case "0886":
		return "Smatfren"
	case "0887":
		return "Smatfren"
	case "0888":
		return "Smatfren"
	case "0889":
		return "Smatfren"
	}
	return ""
}
