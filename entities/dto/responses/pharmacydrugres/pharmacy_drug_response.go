package pharmacydrugres

import (
	"healthcare-capt-america/entities/models"
	"time"

	"github.com/shopspring/decimal"
)

type PharmacyDrugResponse struct {
	DrugId          uint64           `json:"id"`
	DrugName        string           `json:"name"`
	DrugContent     string           `json:"content"`
	ManufactureName string           `json:"manufacture"`
	DrugGenericName string           `json:"generic_name"`
	Description     string           `json:"description"`
	Form            string           `json:"drug_form"`
	Category        string           `json:"category"`
	UnitInPack      string           `json:"unit_in_pack"`
	Weight          uint64           `json:"weight"`
	Height          uint64           `json:"height"`
	Width           uint64           `json:"width"`
	Image           string           `json:"image"`
	SellingUnit     decimal.Decimal  `json:"selling_unit"`
	Status          string           `json:"status"`
	Stock           int              `json:"stock"`
	Pharmacy        PharmacyResponse `json:"pharmacy"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type AddressRes struct {
	Detail   string `json:"detail"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type PharmacyResponse struct {
	ID              uint64     `json:"id"`
	Name            string     `json:"name"`
	PharmaciestName string     `json:"pharmaciest_name"`
	LicenseNumber   string     `json:"pharmaciest_license_number"`
	PhoneNumber     string     `json:"pharmaciest_phone_number"`
	Address         AddressRes `json:"address"`
}

func (pdr *PharmacyDrugResponse) Set(pharmacydrug *models.PharmacyDrug) {
	pdr.DrugId = pharmacydrug.DrugId
	pdr.DrugName = pharmacydrug.Drug.Name
	pdr.ManufactureName = pharmacydrug.Drug.Manufacture.Name
	pdr.Form = pharmacydrug.Drug.Form.Name
	pdr.Category = pharmacydrug.Drug.Category.Name
	pdr.DrugGenericName = pharmacydrug.Drug.GenericName
	pdr.DrugContent = pharmacydrug.Drug.Content
	pdr.Description = pharmacydrug.Drug.Description
	pdr.Weight = pharmacydrug.Drug.Weight.BigInt().Uint64()
	pdr.Height = pharmacydrug.Drug.Height.BigInt().Uint64()
	pdr.Width = pharmacydrug.Drug.Width.BigInt().Uint64()
	pdr.Pharmacy = PharmacyResponse{
		ID:              pharmacydrug.Pharmacy.ID,
		Name:            pharmacydrug.Pharmacy.Name,
		PharmaciestName: pharmacydrug.Pharmacy.PharmaciestName,
		LicenseNumber:   pharmacydrug.Pharmacy.LicenseNumber,
		PhoneNumber:     pharmacydrug.Pharmacy.PhoneNumber,
		Address: AddressRes{
			Detail:   pharmacydrug.Pharmacy.Address.Detail,
			Province: pharmacydrug.Pharmacy.Address.Province.Name,
			City:     pharmacydrug.Pharmacy.Address.City.Name,
		},
	}
	pdr.Weight = pharmacydrug.Drug.Weight.BigInt().Uint64()
	pdr.Height = pharmacydrug.Drug.Height.BigInt().Uint64()
	pdr.Width = pharmacydrug.Drug.Height.BigInt().Uint64()
	pdr.Image = pharmacydrug.Drug.Image
	pdr.UnitInPack = pharmacydrug.Drug.UnitInPack
	pdr.Stock = pharmacydrug.Stock
	pdr.SellingUnit = pharmacydrug.SellingUnit
	pdr.Status = pharmacydrug.Status
	pdr.CreatedAt = pharmacydrug.CreatedAt
	pdr.UpdatedAt = pharmacydrug.UpdatedAt
}
