package middlewares

import (
	"errors"
	"healthcare-capt-america/apperror"
	"healthcare-capt-america/entities/dto/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var (
			clientError     *apperror.ClientError
			serverError     *apperror.ServerError
			validationError *apperror.ValidationError
		)

		err := c.Errors.Last()
		if err == nil {
			return
		}

		resp := responses.DefaultResponse{
			Message: err.Error(),
		}
		switch {
		case errors.As(err, &clientError):
			c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		case errors.As(err, &serverError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		case errors.As(err, &validationError):
			c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		}
	}
}
