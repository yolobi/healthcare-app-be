package userres

import (
	"healthcare-capt-america/entities/dto/responses/adminpharmacyres"
	"healthcare-capt-america/entities/models"
)

type UserDetailResponse struct {
	Id          uint64                         `json:"id"`
	Name        string                         `json:"name"`
	Email       string                         `json:"email"`
	PhoneNumber string                         `json:"phone_number"`
	Photo       string                         `json:"photo"`
	IsVerify    bool                           `json:"is_verify"`
	Addresses   []*adminpharmacyres.AddressRes `json:"addresses"`
}

func NewUserDetail(user *models.User) UserDetailResponse {
	phone := ""
	if user.PhoneNumber != nil {
		phone = *user.PhoneNumber
	}
	addresses := make([]*adminpharmacyres.AddressRes, 0)
	for _, address := range user.Addresses {
		add := adminpharmacyres.AddressRes{}
		ar := add.NewAddressRes(*address)
		addresses = append(addresses, &ar)
	}
	return UserDetailResponse{
		Id:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: phone,
		Photo:       user.Photo,
		IsVerify:    user.IsVerify,
		Addresses:   addresses,
	}
}

func getIconDir(photo string) string {
	if photo != "" {
		photo = photo[len("/varmasea"):]
		photo = photo + "?authuser=2"
		return photo
	}
	return ""
}
