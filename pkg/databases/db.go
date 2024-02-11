package databases

import (
	"healthcare-capt-america/interfaces/repositories/adminrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/authrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/cartrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/doctorrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/journalrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/operationalrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/orderrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyproductrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/productrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/stockmutationrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/telemedicinerepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/userrepositoryinterface"

	"gorm.io/gorm"
)

type Repositories struct {
	adminrepositoryinterface.AdminPharmacyRepository
	doctorrepositoryinterface.DoctorRepository
	authrepositoryinterface.ForgotPasswordRepository
	authrepositoryinterface.RegisterTokenRepository
	authrepositoryinterface.AuthRepository
	userrepositoryinterface.UserRepository
	masterrepositoryinterface.CategoryRepository
	masterrepositoryinterface.CityRepository
	masterrepositoryinterface.FormRepository
	masterrepositoryinterface.ManufactureRepository
	masterrepositoryinterface.PaymentRepository
	masterrepositoryinterface.ProvinceRepository
	masterrepositoryinterface.ShipmentRepository
	masterrepositoryinterface.SubDistrictRepository
	masterrepositoryinterface.StatusMutationRepository
	masterrepositoryinterface.SpecializationRepository
	pharmacyrepositoryinterface.PharmacyRepository
	pharmacyproductrepositoryinterface.PharmacyDrugRepository
	productrepositoryinterface.DrugRepository
	cartrepositoryinterface.CartRepository
	operationalrepositoryinterface.OperationalRepository
	journalrepositoryinterface.JournalRepository
	stockmutationrepositoryinterface.StockMutationRepository
	orderrepositoryinterface.OrderRepository
	masterrepositoryinterface.AddrressRepository
	telemedicinerepositoryinterface.ChatRepository
	telemedicinerepositoryinterface.RoomRepository
	db *gorm.DB
}