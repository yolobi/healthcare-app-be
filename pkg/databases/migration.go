package databases

import "healthcare-capt-america/entities/models"

func Automigration(r *Repositories) {
	r.db.AutoMigrate(
		&models.Manufacture{},
		&models.SubDistrict{},
		&models.City{},
		&models.Province{},
		&models.Shipment{},
		&models.Address{},
		&models.Form{},
		&models.Category{},
		&models.Drug{},
		&models.User{},
		&models.Operational{},
		&models.Payment{},
		&models.Cart{},
		&models.AdminPharmacy{},
		&models.Pharmacy{},
		&models.AdminPharmacyJob{},
		&models.Order{},
		&models.OrderDetail{},
		&models.PharmacyDrug{},
		&models.Specialization{},
		&models.Doctor{},
		&models.RegisterAccountToken{},
		&models.ForgotPasswordToken{},
		&models.Journal{},
		&models.StatusMutation{},
		&models.StockMutation{},
		&models.RoomStatus{},
		&models.Chat{},
		&models.Room{},
	)
}
