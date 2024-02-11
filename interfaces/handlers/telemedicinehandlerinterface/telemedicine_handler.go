package telemedicinehandlerinterface

import "github.com/gin-gonic/gin"

type TelemedicineHandler interface {
	GetAllTelemedicines(*gin.Context)
	GetAllChatByRoomId(*gin.Context)
	Update(*gin.Context)
	CreateRoom(c *gin.Context)
}
