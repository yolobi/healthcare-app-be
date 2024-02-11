package apps

import (
	"healthcare-capt-america/handlers/adminhandler"
	"healthcare-capt-america/handlers/authhandler"
	"healthcare-capt-america/handlers/carthandler"
	"healthcare-capt-america/handlers/checkouthandler"
	"healthcare-capt-america/handlers/doctorhandler"
	"healthcare-capt-america/handlers/journalhandler"
	"healthcare-capt-america/handlers/masterhandler"
	"healthcare-capt-america/handlers/operationalhandler"
	"healthcare-capt-america/handlers/orderhandler"
	"healthcare-capt-america/handlers/paymenthandler"
	"healthcare-capt-america/handlers/pharmacyhandler"
	"healthcare-capt-america/handlers/pharmacyproducthandler"
	"healthcare-capt-america/handlers/productcategoryhandler"
	"healthcare-capt-america/handlers/producthandler"
	"healthcare-capt-america/handlers/provincehandler"
	"healthcare-capt-america/handlers/stockmutationhandler"
	"healthcare-capt-america/handlers/telemedicinehandler"
	"healthcare-capt-america/handlers/userhandler"
	"healthcare-capt-america/interfaces/handlers/adminhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/authhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/carthandlerinterface"
	"healthcare-capt-america/interfaces/handlers/checkouthandlerinterface"
	"healthcare-capt-america/interfaces/handlers/doctorhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/journalhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/masterhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/operationalhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/orderhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/paymenthandlerinterface"
	"healthcare-capt-america/interfaces/handlers/pharmacyhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/pharmacyproducthandlerinterface"
	"healthcare-capt-america/interfaces/handlers/productcategoryhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/producthandlerinterface"
	"healthcare-capt-america/interfaces/handlers/provincehandlerinterface"
	"healthcare-capt-america/interfaces/handlers/stockmutationhandlerinterface"
	"healthcare-capt-america/interfaces/handlers/telemedicinehandlerinterface"
	"healthcare-capt-america/interfaces/handlers/userhandlerinterface"
	"healthcare-capt-america/pkg/databases"
)

type Handlers struct {
	AdminPharmacy          adminhandlerinterface.AdminPharmacyHandler
	PharmacyHandler        pharmacyhandlerinterface.PharmacyHandler
	ProductHandler         producthandlerinterface.DrugHandler
	ProductCategoryHandler productcategoryhandlerinterface.CategoryHandler
	AuthHandler            authhandlerinterface.AuthHandler
	ForgotPasswordHandler  authhandlerinterface.ForgotPasswordHandler
	PharmacyProductHandler pharmacyproducthandlerinterface.PharmacyDrugHandler
	CartHandler            carthandlerinterface.CartHandler
	OperationalHandler     operationalhandlerinterface.OperationalHandler
	JournalHandler         journalhandlerinterface.JournalHandler
	StockMutation          stockmutationhandlerinterface.StockMutationHandler
	OrderHandler           orderhandlerinterface.OrderHandler
	PaymentHandler         paymenthandlerinterface.PaymentHandler
	DoctorHandler          doctorhandlerinterface.DoctorHandler
	ProvinceHandler        provincehandlerinterface.ProvinceHandler
	UserHandler            userhandlerinterface.UserHandler
	ManufactureHandler     masterhandlerinterface.ManufactureHandler
	FormHandler            masterhandlerinterface.FormHandler
	CheckoutHandler        checkouthandlerinterface.CheckoutHandler
	TelemedicineHandler    telemedicinehandlerinterface.TelemedicineHandler
	ShipmentHandler        masterhandler.ShipmentHandler
	ChatHandler            telemedicinehandler.ChattingHandler
	SpecializationHandler  masterhandlerinterface.SpecializationHandler
}

func NewHandlers(repo *databases.Repositories) *Handlers {
	usecases := NewUsecases(repo)
	return &Handlers{
		AdminPharmacy:          adminhandler.NewAdminPharmacyHandler(usecases.AdminPharmacyUsecase),
		PharmacyHandler:        pharmacyhandler.NewPharmacyHandler(usecases.PharmacyUsecase),
		ProductHandler:         producthandler.NewDrugHandler(usecases.ProductUsecase),
		ProductCategoryHandler: productcategoryhandler.NewCategoryHandler(usecases.ProductCategoryUsecase),
		AuthHandler:            authhandler.NewAuthHandler(usecases.AuthUsecase),
		ForgotPasswordHandler:  authhandler.NewForgotPasswordHandler(usecases.ForgotPasswordUsecase),
		PharmacyProductHandler: pharmacyproducthandler.NewPharmacyDrugHandler(usecases.PharmacyProductUsecase),
		CartHandler:            carthandler.NewCartHandler(usecases.CartUsecase),
		OperationalHandler:     operationalhandler.NewOperationalHandler(usecases.OperationalUsecase),
		JournalHandler:         journalhandler.NewJournalHandler(usecases.JournalUsecase),
		StockMutation:          stockmutationhandler.NewStockMutationHandler(usecases.StockMutationUsecase),
		OrderHandler:           orderhandler.NewOrderHandler(usecases.OrderUsecase),
		PaymentHandler:         paymenthandler.NewPaymentHandler(usecases.PaymentUsecase),
		DoctorHandler:          doctorhandler.NewDoctorHandler(usecases.DoctorUsecase),
		ProvinceHandler:        provincehandler.NewProvinceHandler(usecases.ProvinceUsecase),
		UserHandler:            userhandler.NewUserHandler(usecases.UserUsecase),
		ManufactureHandler:     masterhandler.NewManufactureHandler(usecases.ManufactureUsecase),
		FormHandler:            masterhandler.NewFormHandler(usecases.FormUsecase),
		CheckoutHandler:        checkouthandler.NewCheckoutHandler(usecases.CheckoutUsecase),
		TelemedicineHandler:    telemedicinehandler.NewTelemedicineHandler(usecases.TelemedicineUsecase),
		ShipmentHandler:        *masterhandler.NewShipmentHandler(usecases.ShipmentUsecase),
		ChatHandler:            *telemedicinehandler.NewChattingHandler(&usecases.ChatUsecase),
		SpecializationHandler:  masterhandler.NewSpecializationHandler(usecases.SpecializationUsecase),
	}
}
