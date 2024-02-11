package journalhandler

import (
	"healthcare-capt-america/entities/dto/requests"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/journalres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/usecases/journalusecaseinterface"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JournalHandler struct {
	usecase journalusecaseinterface.JournalUsecase
}

func NewJournalHandler(usecase journalusecaseinterface.JournalUsecase) *JournalHandler {
	return &JournalHandler{usecase: usecase}
}

func (handler *JournalHandler) FindAllJournal(ctx *gin.Context) {
	var request requests.GlobalFilter
	var err error
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	request.StartDate, err = time.Parse(enums.DDMMYY, time.Time{}.UTC().Format(enums.DDMMYY))
	if err != nil {
		err = ctx.Error(err)
		return
	}
	request.EndDate, err = time.Parse(enums.DDMMYY, time.Now().UTC().Format(enums.DDMMYY))
	if err != nil {
		err = ctx.Error(err)
		return
	}
	err = ctx.ShouldBindQuery(&request)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	qry := requests.NewQuery(&request)
	qry.AddCondition("journals.updated_at", requests.GreaterEqual, request.StartDate)
	qry.AddCondition("journals.updated_at", requests.Less, request.EndDate.Add(24*time.Hour))
	journals, pagination, err := handler.usecase.FindAll(ctx, qry, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	resp := make([]journalres.JournalResponse, 0)
	for _, journal := range journals {
		respJournal := journalres.JournalResponse{}
		respJournal.Set(&journal)
		resp = append(resp, respJournal)
	}
	pagination.Items = resp
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: responses.NewPagination(pagination, "journals"),
	})
}
