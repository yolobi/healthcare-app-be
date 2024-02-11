package apps

import (
	"healthcare-capt-america/interfaces/usecases/adminusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/authusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/cartusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/checkoutusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/doctorusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/journalusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/operationalusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/orderusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/paymentusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/pharmacyproductusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/pharmacyusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/productcategoryusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/productusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/provinceusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/stockmutationusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/telemedicineusecaseinterface"
	"healthcare-capt-america/interfaces/usecases/userusecaseinterface"
	"healthcare-capt-america/pkg/databases"
	"healthcare-capt-america/usecases/adminusecase"
	"healthcare-capt-america/usecases/authusecase"
	"healthcare-capt-america/usecases/cartusecase"
	"healthcare-capt-america/usecases/checkoutusecase"
	"healthcare-capt-america/usecases/doctorusecase"
	"healthcare-capt-america/usecases/journalusecase"
	"healthcare-capt-america/usecases/masterusecase"
	"healthcare-capt-america/usecases/operationalusecase"
	"healthcare-capt-america/usecases/orderusecase"
	"healthcare-capt-america/usecases/paymentusecase"
	"healthcare-capt-america/usecases/pharmacyproductusecase"
	"healthcare-capt-america/usecases/pharmacyusecase"
	"healthcare-capt-america/usecases/productcategoryusecase"
	"healthcare-capt-america/usecases/telemedicineusecase"

	"healthcare-capt-america/usecases/productusecase"

	"healthcare-capt-america/usecases/provinceusecase"
	"healthcare-capt-america/usecases/stockmutationusecase"
	"healthcare-capt-america/usecases/userusecase"
)

type Usecases struct {
	AdminPharmacyUsecase   adminusecaseinterface.AdminPharmacyUsecase
	PharmacyUsecase        pharmacyusecaseinterface.PharmacyUsecase
	ProductUsecase         productusecaseinterface.DrugUsecase
	ProductCategoryUsecase productcategoryusecaseinterface.CategoryUsecase
	AuthUsecase            authusecaseinterface.AuthUsecase
	ForgotPasswordUsecase  authusecaseinterface.ForgotPasswordUsecase
	PharmacyProductUsecase pharmacyproductusecaseinterface.PharmacyDrugUsecase
	CartUsecase            cartusecaseinterface.CartUsecase
	OperationalUsecase     operationalusecaseinterface.OperationalUsecase
	JournalUsecase         journalusecaseinterface.JournalUsecase
	StockMutationUsecase   stockmutationusecaseinterface.StockMutationUsecase
	OrderUsecase           orderusecaseinterface.OrderUsecase
	PaymentUsecase         paymentusecaseinterface.PaymentUsecase
	DoctorUsecase          doctorusecaseinterface.DoctorUsecase
	ProvinceUsecase        provinceusecaseinterface.ProvinceUsecase
	UserUsecase            userusecaseinterface.UserUsecase
	ManufactureUsecase     masterusecaseinterface.ManufactureUsecase
	FormUsecase            masterusecaseinterface.FormUsecase
	CheckoutUsecase        checkoutusecaseinterface.CheckoutUsecase
	TelemedicineUsecase    telemedicineusecaseinterface.TelemedicineUsecase
	ShipmentUsecase        masterusecaseinterface.ShipmentUsecase
	ChatUsecase            telemedicineusecase.ChattingUsecase
	SpecializationUsecase  masterusecaseinterface.SpecializationUsecase
}

func NewUsecases(repo *databases.Repositories) *Usecases {
	return &Usecases{
		AdminPharmacyUsecase:   adminusecase.NewAdminPharmacyUsecase(repo.AdminPharmacyRepository, repo.UserRepository),
		PharmacyUsecase:        pharmacyusecase.NewPharmacyUsecase(repo.PharmacyRepository, repo.AdminPharmacyRepository, repo.OperationalRepository, repo.AddrressRepository),
		AuthUsecase:            authusecase.NewAuthUsecase(repo.AuthRepository, repo.RegisterTokenRepository, repo.UserRepository, repo.DoctorRepository),
		ProductUsecase:         productusecase.NewDrugUsecase(repo.DrugRepository, repo.CategoryRepository, repo.FormRepository, repo.ManufactureRepository, repo.AddrressRepository),
		ProductCategoryUsecase: productcategoryusecase.NewCategoryUsecase(repo.CategoryRepository),
		ForgotPasswordUsecase:  authusecase.NewForgotPasswordUsecase(repo.ForgotPasswordRepository, repo.UserRepository),
		PharmacyProductUsecase: pharmacyproductusecase.NewPharmacyDrugUsecase(repo.PharmacyDrugRepository, repo.PharmacyRepository, repo.DrugRepository, repo.AdminPharmacyRepository),
		CartUsecase:            cartusecase.NewCartUsecase(repo.UserRepository, repo.PharmacyDrugRepository, repo.CartRepository, repo.AddrressRepository),
		JournalUsecase:         journalusecase.NewJournalUsecase(repo.JournalRepository, repo.AdminPharmacyRepository),
		OperationalUsecase:     operationalusecase.NewOperationalUsecase(repo.PharmacyRepository, repo.OperationalRepository),
		StockMutationUsecase:   stockmutationusecase.NewStockMutationUsecase(repo.StockMutationRepository, repo.PharmacyDrugRepository, repo.PharmacyRepository, repo.AdminPharmacyRepository),
		OrderUsecase:           orderusecase.NewOrderUsecase(repo.OrderRepository, repo.PaymentRepository),
		PaymentUsecase:         paymentusecase.NewPaymentUsecase(repo.PaymentRepository, repo.OrderRepository),
		DoctorUsecase:          doctorusecase.NewDoctorUsecase(repo.DoctorRepository, repo.UserRepository),
		ProvinceUsecase:        provinceusecase.NewProvinceUsecase(repo.ProvinceRepository),
		ManufactureUsecase:     masterusecase.NewManufactureUsecase(repo.ManufactureRepository),
		FormUsecase:            masterusecase.NewFormUsecase(repo.FormRepository),
		CheckoutUsecase:        checkoutusecase.NewCheckoutUsecase(repo.CartRepository, repo.UserRepository, repo.DrugRepository, repo.OrderRepository, repo.PharmacyRepository, repo.PharmacyDrugRepository, repo.AddrressRepository),
		UserUsecase:            userusecase.NewUserUsecase(repo.UserRepository, repo.AddrressRepository),
		TelemedicineUsecase:    telemedicineusecase.NewTelemedicineUsecase(repo.RoomRepository, repo.ChatRepository, repo.DoctorRepository),
		ShipmentUsecase:        masterusecase.NewShipmentUsecase(repo.ShipmentRepository, repo.AddrressRepository, repo.PharmacyRepository),
		ChatUsecase:            *telemedicineusecase.NewChattingUsecase(repo.RoomRepository, repo.ChatRepository, repo.DoctorRepository),
		SpecializationUsecase:  masterusecase.NewSpecializationUsecase(repo.SpecializationRepository),
	}
}
