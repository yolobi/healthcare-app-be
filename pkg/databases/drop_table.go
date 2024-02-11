package databases

import (
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
)

func (r *Repositories) DropTable() {
	r.db.Migrator().DropTable(
		enums.Authority_role_permissions_table,
		enums.Authority_permissions_table,
		enums.Authority_user_roles_table,
		enums.Authority_roles_table,
		&models.ForgotPasswordToken{},
		&models.RegisterAccountToken{},
		&models.Doctor{},
		&models.Specialization{},
		&models.PharmacyDrug{},
		&models.OrderDetail{},
		&models.Order{},
		&models.Cart{},
		&models.AdminPharmacyJob{},
		&models.AdminPharmacy{},
		&models.Payment{},
		&models.Operational{},
		&models.User{},
		&models.Pharmacy{},
		&models.Drug{},
		&models.Category{},
		&models.Form{},
		&models.Address{},
		&models.Shipment{},
		&models.Province{},
		&models.City{},
		&models.SubDistrict{},
		&models.Manufacture{},
		&models.Journal{},
		&models.StockMutation{},
		&models.StatusMutation{},
		&models.Room{},
		&models.RoomStatus{},
		&models.Chat{},
	)
}
