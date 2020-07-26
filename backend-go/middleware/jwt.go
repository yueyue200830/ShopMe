package middleware

import (
	"backend-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := 0
		token := c.GetHeader("Bearer")
		if token == "" {
			status = 1
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				status = 1
			} else if time.Now().Unix() > claims.ExpiresAt {
				status = 1
			}
		}

		if status > 0 {
			c.JSON(http.StatusUnauthorized, status)
			c.Abort()
			return
		}

		c.Next()
	}
}
