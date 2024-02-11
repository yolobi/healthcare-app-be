package adminpharmacyres

import "healthcare-capt-america/entities/models"

type AddressRes struct {
	Id         uint64 `json:"id"`
	Detail     string `json:"detail"`
	ProvinceID uint64 `json:"province_id"`
	Province   string `json:"province"`
	CityID     uint64 `json:"city_id"`
	City       string `json:"city"`
	Longtitude string `json:"longitude"`
	Latitude   string `json:"latitude"`
	IsDefault  bool   `json:"is_default"`
}

func (a *AddressRes) NewAddressRes(addr models.Address) AddressRes {
	return AddressRes{
		Id:         addr.ID,
		Detail:     addr.Detail,
		ProvinceID: addr.ProvinceID,
		Province:   addr.Province.Name,
		CityID:     addr.CityID,
		City:       addr.City.Name,
		Longtitude: addr.Longtitude.String(),
		Latitude:   addr.Latitude.String(),
		IsDefault:  addr.IsDefault,
	}
}
