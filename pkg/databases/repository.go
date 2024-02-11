package databases

import (
	"healthcare-capt-america/repositories/adminrepository"
	"healthcare-capt-america/repositories/authrepository"
	"healthcare-capt-america/repositories/cartrepository"
	"healthcare-capt-america/repositories/doctorrepository"
	"healthcare-capt-america/repositories/journalrepository"
	"healthcare-capt-america/repositories/masterrepository"
	"healthcare-capt-america/repositories/operationalrepository"
	"healthcare-capt-america/repositories/orderrepository"
	"healthcare-capt-america/repositories/pharmacyproductrepository"
	"healthcare-capt-america/repositories/pharmacyrepository"
	"healthcare-capt-america/repositories/productrepository"
	"healthcare-capt-america/repositories/stockmutationrepository"
	"healthcare-capt-america/repositories/telemedicinerepository"
	"healthcare-capt-america/repositories/userrepository"

	"gorm.io/gorm"
)

func (r *Repositories) InitRepositories() {
	r.CategoryRepository = masterrepository.NewCategoryRepository(r.db)
	r.CityRepository = masterrepository.NewCityRepository(r.db)
	r.FormRepository = masterrepository.NewFormRepository(r.db)
	r.ManufactureRepository = masterrepository.NewManufactureRepository(r.db)
	r.PaymentRepository = masterrepository.NewPaymentRepository(r.db)
	r.ProvinceRepository = masterrepository.NewProvinceRepository(r.db)
	r.ShipmentRepository = masterrepository.NewShipmentRepository(r.db)
	r.SubDistrictRepository = masterrepository.NewSubdistrictRepository(r.db)
	r.StatusMutationRepository = masterrepository.NewStatusMutationRepository(r.db)
	r.SpecializationRepository = masterrepository.NewSpecializationRepository(r.db)
	r.PharmacyRepository = pharmacyrepository.NewPharmacyRepository(r.db)
	r.DrugRepository = productrepository.NewDrugRepository(r.db)
	r.PharmacyDrugRepository = pharmacyproductrepository.NewPharmacyDrugRepository(r.db)
	r.AuthRepository = authrepository.NewAuthrepository(r.db)
	r.RegisterTokenRepository = authrepository.NewRegisterTokenRepostiory(r.db)
	r.UserRepository = userrepository.NewUserRepository(r.db)
	r.ForgotPasswordRepository = authrepository.NewForgotPasswordRepository(r.db)
	r.PharmacyDrugRepository = pharmacyproductrepository.NewPharmacyDrugRepository(r.db)
	r.DrugRepository = productrepository.NewDrugRepository(r.db)
	r.CartRepository = cartrepository.NewCartRepository(r.db)
	r.OperationalRepository = operationalrepository.NewOperationalRepository(r.db)
	r.JournalRepository = journalrepository.NewJournalRepository(r.db)
	r.StockMutationRepository = stockmutationrepository.NewStockMutationRepository(r.db)
	r.AdminPharmacyRepository = adminrepository.NewAdminPharmacyRepository(r.db)
	r.DoctorRepository = doctorrepository.NewDoctorRepository(r.db)
	r.OrderRepository = orderrepository.NewOrderRepository(r.db)
	r.AddrressRepository = masterrepository.NewADdressRepository(r.db)
	r.ChatRepository = telemedicinerepository.NewChatRepository(r.db)
	r.RoomRepository = telemedicinerepository.NewRoomRepository(r.db)
}

func (r *Repositories) GetDB() *gorm.DB {
	return r.db
}
