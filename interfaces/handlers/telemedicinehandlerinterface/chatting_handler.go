package telemedicinehandlerinterface

import "github.com/gin-gonic/gin"

type ChattingHandler interface {
	WebSocketHandler(ctx *gin.Context)
	GetChatHistoriesRoom(c *gin.Context)
}
