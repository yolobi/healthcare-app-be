package masterusecase

import (
	"context"
	"fmt"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
	"healthcare-capt-america/interfaces/repositories/masterrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/pharmacyrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/masterusecaseinterface"
	"healthcare-capt-america/services"

	"github.com/shopspring/decimal"
)

type shipmentUsecase struct {
	shipRepo masterrepositoryinterface.ShipmentRepository
	addrRepo masterrepositoryinterface.AddrressRepository
	pharRepo pharmacyrepositoryinterface.PharmacyRepository
}

func (su *shipmentUsecase) FindAll(ctx context.Context) ([]*models.Shipment, error) {
	shipments, err := su.shipRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return shipments, nil
}

func (su *shipmentUsecase) CalculateDistance(ctx context.Context, address_id uint64, pharmacy_id uint64, weight int64) ([]*transaction.ShipmentFee, error) {
	address, err := su.addrRepo.FindById(ctx, address_id)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find address with id %d", address_id))
	}

	phar, err := su.pharRepo.FindById(ctx, pharmacy_id)
	if err != nil {
		return nil, err
	}
	if phar == nil {
		return nil, apperror.NewClientError(fmt.Errorf("can't find pharmacy with id %d", pharmacy_id))
	}

	dist, err := services.DistanceMatrix(transaction.NewPosition(phar.Address.Longtitude, phar.Address.Latitude), transaction.NewPosition(address.Longtitude, address.Latitude))
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	rajaongkir, err := services.RajaongkirCost(address.CityID, phar.Address.CityID, weight)
	if err != nil {
		return nil, err
	}

	ships, err := su.shipRepo.FindOfficial(ctx)
	if err != nil {
		return nil, err
	}
	resp := make([]*transaction.ShipmentFee, 0)
	for _, ship := range ships {
		resp = append(resp, &transaction.ShipmentFee{
			Name: ship.Name,
			Fee:  dist.Mul(ship.CostPerKM).Div(decimal.NewFromInt(1000)).String(),
		})
	}
	resp = append(resp, rajaongkir...)
	return resp, nil
}

func NewShipmentUsecase(sr masterrepositoryinterface.ShipmentRepository, ar masterrepositoryinterface.AddrressRepository, pr pharmacyrepositoryinterface.PharmacyRepository) *shipmentUsecase {
	return &shipmentUsecase{
		shipRepo: sr,
		addrRepo: ar,
		pharRepo: pr,
	}
}

var _ masterusecaseinterface.ShipmentUsecase = &shipmentUsecase{}
