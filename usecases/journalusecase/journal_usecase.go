package journalusecase

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/interfaces/repositories/adminrepositoryinterface"
	"healthcare-capt-america/interfaces/repositories/journalrepositoryinterface"
	"healthcare-capt-america/interfaces/usecases/journalusecaseinterface"
)

type journalUsecase struct {
	journalRepo journalrepositoryinterface.JournalRepository
	adminRepo   adminrepositoryinterface.AdminPharmacyRepository
}

func NewJournalUsecase(journalRepo journalrepositoryinterface.JournalRepository, adminRepo adminrepositoryinterface.AdminPharmacyRepository) *journalUsecase {
	return &journalUsecase{journalRepo: journalRepo, adminRepo: adminRepo}
}

func (usecase *journalUsecase) FindAll(ctx context.Context, qry *requests.GlobalQuery, uId uint64) ([]models.Journal, *responses.Pagination, error) {
	return usecase.journalRepo.FindAll(ctx, qry, uId)
}

var _ journalusecaseinterface.JournalUsecase = &journalUsecase{}
