package databases

import (
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/utils"

	"github.com/shopspring/decimal"
)

var roles = []enums.Role{
	enums.Admin,
	enums.PharmacyAdmin,
	enums.User,
	enums.Doctor,
}

var permissions = []enums.Permission{
	enums.AdminDashboard,
}

var categories = []models.Category{
	{Name: "Free Medicines", Icon: "public/seeding_user_photo/medicine/129.jpg"},
	{Name: "Prescription Drugs", Icon: "public/seeding_user_photo/medicine/112.jpg"},
	{Name: "Limited Free Medicines", Icon: "public/seeding_user_photo/medicine/119.jpg"},
	{Name: "Non Medicines", Icon: "public/seeding_user_photo/medicine/136.jpg"},
}

var shipments = []models.Shipment{
	{Name: "Instant", CostPerKM: decimal.NewFromInt(1000)},
	{Name: "JNE", CostPerKM: decimal.NewFromInt(0)},
}

var room_status = []models.RoomStatus{
	{Name: "ongoing"},
	{Name: "end"},
}

func HotHash(pass string) *string {
	res, _ := utils.HashPassword(pass)
	return &res
}
