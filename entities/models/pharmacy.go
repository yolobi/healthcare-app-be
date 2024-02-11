package models

import (
	"time"

	"gorm.io/gorm"
)

type Pharmacy struct {
	ID              uint64         `gorm:"column:id;primaryKey;autoIncrement"`
	AddressID       uint64         `gorm:"column:address_id;not null" csv:"address_id"`
	Name            string         `gorm:"column:name;not null;unique" csv:"name"`
	PharmaciestName string         `gorm:"column:pharmaciest_name;not null" csv:"pharmaciest_name"`
	LicenseNumber   string         `gorm:"column:license_number;not null;unique" csv:"license_number"`
	PhoneNumber     string         `gorm:"column:phone_number;not null;unique" csv:"phone_number"`
	AdminPharmacyId uint64         `csv:"admin_pharmacy_id"`
	AdminPharmacies AdminPharmacy  `gorm:"foreignKey:AdminPharmacyId"`
	Operationals    []Operational  `gorm:"foreignKey:PharmacyID"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Address         Address        `gorm:"foreignKey:AddressID"`
}

func (pharmacy *Pharmacy) EditRequest(editedPharmacyReq *Pharmacy) *Pharmacy {
	pharmacy.Name = editedPharmacyReq.Name
	pharmacy.PharmaciestName = editedPharmacyReq.PharmaciestName
	pharmacy.LicenseNumber = editedPharmacyReq.LicenseNumber
	pharmacy.PhoneNumber = editedPharmacyReq.PhoneNumber
	return pharmacy
}
