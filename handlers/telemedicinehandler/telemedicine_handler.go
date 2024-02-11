package telemedicinehandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/requests/telemedicinereq"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/telemecineres"
	"healthcare-capt-america/entities/models"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/telemedicinehandlerinterface"
	"healthcare-capt-america/interfaces/usecases/telemedicineusecaseinterface"
	"net/http"
	"strconv"
)

type telemedicineHandler struct {
	usecase telemedicineusecaseinterface.TelemedicineUsecase
}

func (handler *telemedicineHandler) CreateRoom(c *gin.Context) {
	uId := c.Value(enums.UserIdKey.Key).(uint64)
	var request *telemedicinereq.CreateRoom
	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewValidationError(err))
		return
	}
	room := &models.Room{
		UserID:       uId,
		DoctorID:     request.DoctorID,
		RoomStatusID: enums.Ongoing,
	}
	result, err := handler.usecase.CreateRoom(c.Request.Context(), room)
	if err != nil {
		c.Error(err)
		return
	}
	res := telemecineres.RoomResponse{}
	res.Set(result)
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"room": res}})
}

func (handler *telemedicineHandler) GetAllChatByRoomId(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	room_id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	chats, err := handler.usecase.FindAllByRoomId(ctx, room_id, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	response := responses.DefaultResponse{Data: gin.H{"chats": telemecineres.NewChatResponses(chats)}}
	ctx.JSON(http.StatusCreated, response)
}

func (handler *telemedicineHandler) GetAllTelemedicines(ctx *gin.Context) {
	var err error
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	telemedicines, err := handler.usecase.FindAll(ctx.Request.Context(), uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	resp := make([]telemecineres.RoomResponse, 0)
	for _, telemedicine := range telemedicines {
		respTelemedicine := telemecineres.RoomResponse{}
		respTelemedicine.Set(telemedicine)
		resp = append(resp, respTelemedicine)
	}
	ctx.JSON(http.StatusOK, responses.DefaultResponse{
		Data: gin.H{"rooms": resp},
	})
}

func (handler *telemedicineHandler) Update(ctx *gin.Context) {
	uId := ctx.Value(enums.UserIdKey.Key).(uint64)
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	var request telemedicinereq.UpdateRoomStatus
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	room := request.ToRoom()
	room.ID = id
	edited, err := handler.usecase.Update(ctx.Request.Context(), &room, uId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	response := responses.DefaultResponse{Data: gin.H{"product": edited}}
	ctx.JSON(http.StatusCreated, response)
}

func NewTelemedicineHandler(telemedicineUsecase telemedicineusecaseinterface.TelemedicineUsecase) *telemedicineHandler {
	return &telemedicineHandler{
		usecase: telemedicineUsecase,
	}
}

var _ telemedicinehandlerinterface.TelemedicineHandler = &telemedicineHandler{}
