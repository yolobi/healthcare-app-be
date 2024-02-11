package journalrepositoryinterface

import (
	"context"
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/models"
)

type JournalRepository interface {
	FindAll(context.Context, *requests.GlobalQuery, uint64) ([]models.Journal, *responses.Pagination, error)
}
