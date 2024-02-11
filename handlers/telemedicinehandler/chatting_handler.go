package telemedicinehandler

import (
	"context"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/entities/dto/responses/telemecineres"
	"healthcare-capt-america/enums"
	"healthcare-capt-america/interfaces/handlers/telemedicinehandlerinterface"
	"healthcare-capt-america/interfaces/usecases/telemedicineusecaseinterface"
	"healthcare-capt-america/pkg/security"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

type ChattingHandler struct {
	chatUsecase telemedicineusecaseinterface.ChattingUsecase
}

func (ch *ChattingHandler) GetChatHistoriesRoom(c *gin.Context) {
	rId := c.Param("id")
	roomId, err := strconv.Atoi(rId)
	if err != nil {
		apperror.NewClientError(err, "room id is invalid format")
		return
	}
	chats, err := ch.chatUsecase.GetAllChatHistories(c.Request.Context(), uint64(roomId))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, responses.DefaultResponse{Data: gin.H{"chats": telemecineres.NewChatResponses(chats)}})
}

func (ch *ChattingHandler) WebSocketHandler(ctx *gin.Context) {
	tokenStr := ctx.Query("token")
	claims := &security.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(j *jwt.Token) (interface{}, error) {
		return security.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		ctx.Error(apperror.NewValidationError(err))
		return
	}
	ctx2 := context.WithValue(ctx, enums.UserIdKey, claims.UserID)
	room_id := ctx.Param("room_id")
	rid, err := strconv.Atoi(room_id)
	if err != nil {
		ctx.Error(apperror.NewValidationError(err))
		return
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	log.Println("Conn:", conn.LocalAddr().String())
	err = ch.chatUsecase.HandleWebSocket(ctx2, claims.UserID, uint64(rid), conn)
	if err != nil {
		ctx.Error(err)
		return
	}
}

func NewChattingHandler(cu telemedicineusecaseinterface.ChattingUsecase) *ChattingHandler {
	return &ChattingHandler{
		chatUsecase: cu,
	}
}

var _ telemedicinehandlerinterface.ChattingHandler = &ChattingHandler{}
