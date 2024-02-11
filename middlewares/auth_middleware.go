package middlewares

import (
	"errors"
	"healthcare-capt-america/entities/dto/responses"
	"healthcare-capt-america/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func validateToken(c *gin.Context) (*security.JWTClaim, error) {
	tokenAuthorization := c.GetHeader("Authorization")
	if tokenAuthorization == "" {
		return nil, errors.New("no token")
	}
	tokenStr := tokenAuthorization[len("Bearer "):]
	if tokenStr == "" {
		return nil, errors.New("bad token")
	}
	claims := &security.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(j *jwt.Token) (interface{}, error) {
		return security.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("bad token")
	}
	return claims, nil
}

func AuthMiddleware(c *gin.Context) {
	claims, err := validateToken(c)
	if err != nil {
		resp := responses.DefaultResponse{Message: "Bad Token"}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	if claims == nil {
		resp := responses.DefaultResponse{Message: "Unauthorized"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	c.Set("user-id", claims.UserID)
	c.Set("user-email", claims.Email)
	c.Next()
}
