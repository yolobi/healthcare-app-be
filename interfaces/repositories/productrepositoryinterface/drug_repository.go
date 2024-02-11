package productrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/drugsres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/entities/models/transaction"
)

type DrugRepository interface {
	Save(context.Context, *models.Drug) (*models.Drug, error)
	Update(context.Context, *models.Drug) (*models.Drug, error)
	FindByID(context.Context, uint64) (*drugsres.DrugMaster, error)
	FindByIDUpdate(context.Context, uint64) (*models.Drug, error)
	FindAllWithDist(context.Context, transaction.Position) ([]*drugsres.DrugUserRes, error)
	FindAll(context.Context, *requests.GlobalQuery) ([]drugsres.DrugMaster, *responses.Pagination, error)
}
