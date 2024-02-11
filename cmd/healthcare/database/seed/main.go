package main

import (
	"fmt"
	"healthcare-capt-america/entities/dto/seeding"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/pkg/configs"
	"healthcare-capt-america/pkg/databases"
	"healthcare-capt-america/services"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func structFromCsv[T any](fileName string, entities []*T) ([]*T, error) {
	file := fmt.Sprintf("%s.csv", fileName)
	in, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	if err := gocsv.UnmarshalFile(in, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

var addresses = make([]*seeding.Address, 0)
var admin_pharmacies = make([]*models.AdminPharmacy, 0)
var cities = make([]*seeding.City, 0)
var doctors = make([]*seeding.Doctor, 0)
var forms = make([]*models.Form, 0)
var manufactures = make([]*models.Manufacture, 0)
var pharmacies = make([]*seeding.Pharmacy, 0)
var pharmacy_products = make([]*seeding.PharmacyDrug, 0)
var products = make([]*seeding.Drug, 0)
var provinces = make([]*seeding.Province, 0)
var specializations = make([]*models.Specialization, 0)
var statusMutations = make([]*models.StatusMutation, 0)
var users = make([]*seeding.User, 0)

func main() {
	configPath, err := configs.ParseFlags()
	if err != nil {
		log.Fatal(err)
		return
	}
	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	repo, err := databases.NewRepositories(config)
	if err != nil {
		return
	}
	repo.InitRepositories()
	services.InitAuthority(repo.GetDB())
	db := repo.GetDB()

	addresses, _ = structFromCsv[seeding.Address]("./files/csv/address", addresses)
	cities, _ = structFromCsv[seeding.City]("./files/csv/city", cities)
	doctors, _ = structFromCsv[seeding.Doctor]("./files/csv/doctor", doctors)
	forms, _ = structFromCsv[models.Form]("./files/csv/form", forms)
	manufactures, _ = structFromCsv[models.Manufacture]("./files/csv/manufacture", manufactures)
	pharmacies, _ = structFromCsv[seeding.Pharmacy]("./files/csv/pharmacy", pharmacies)
	pharmacy_products, _ = structFromCsv[seeding.PharmacyDrug]("./files/csv/pharmacy_product", pharmacy_products)
	products, _ = structFromCsv[seeding.Drug]("./files/csv/product", products)
	provinces, _ = structFromCsv[seeding.Province]("./files/csv/province", provinces)
	specializations, _ = structFromCsv[models.Specialization]("./files/csv/specialization", specializations)
	statusMutations, _ = structFromCsv[models.StatusMutation]("./files/csv/status_mutation", statusMutations)
	users, _ = structFromCsv[seeding.User]("./files/csv/user", users)

	for i := 51; i <= 60; i++ {
		userId := uint64(i)
		admin_pharmacy := models.AdminPharmacy{UserId: userId}
		admin_pharmacies = append(admin_pharmacies, &admin_pharmacy)
	}

	repo.CategorySeeding()
	repo.RoleSeeding()
	repo.PermissionSeeding()
	repo.ShipmentSeeding()
	repo.RoomStatusSeeding()

	dataAddresses := seeding.ModelAddress(addresses)
	dataCities := seeding.ModelCity(cities)
	dataDoctors := seeding.ModelDoctor(doctors)
	dataPharmacies := seeding.ModelPharmacy(pharmacies)
	dataPharmacyProducts := seeding.ModelPharmacyDrug(pharmacy_products)
	dataProducts := seeding.ModelProduct(products)
	dataProvinces := seeding.ModelProvince(provinces)
	dataUsers := seeding.ModelUser(users)

	db.Create(specializations)
	db.Create(dataUsers)
	db.Exec(enums.SeedingAuthorityAdminRoles)
	db.Exec(enums.SeedingAuthoritySuperAdminRoles)
	db.Exec(enums.SeedingAuthorityUserRoles)
	db.Exec(enums.SeedingAuthorityDoctorRoles)
	db.Create(dataDoctors)
	db.Create(admin_pharmacies)
	db.Create(forms)
	db.Create(manufactures)
	db.Create(statusMutations)
	db.Create(dataProvinces)
	db.Create(dataCities)
	db.Create(dataAddresses)
	db.Create(dataProducts)
	db.Create(dataPharmacies)
	db.Create(dataPharmacyProducts)
}
