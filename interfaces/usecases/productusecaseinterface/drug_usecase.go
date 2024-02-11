package productusecaseinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/entities/models"
)

type DrugUsecase interface {
	CreateDrug(context.Context, *models.Drug) (*models.Drug, error)
	EditDrug(context.Context, *models.Drug) (*models.Drug, error)
	UpdateDrug(context.Context, uint64, string) error
	FindAllDrug(context.Context, *requests.GlobalQuery) ([]drugsres.DrugMaster, *responses.Pagination, error)
	FindByID(context.Context, uint64) (*drugsres.DrugMaster, error)
	FindAllWithDist(ctx context.Context) ([]*drugsres.DrugUserRes, error)
}
