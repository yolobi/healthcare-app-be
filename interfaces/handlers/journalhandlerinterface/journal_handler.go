package journalhandlerinterface

import "github.com/gin-gonic/gin"

type JournalHandler interface {
	FindAllJournal(*gin.Context)
}
