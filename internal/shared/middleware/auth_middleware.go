package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/homepage408/syncra/internal/shared/constant"
)

const (
	UserIDContextKey = "user_id"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header is required",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header",
			})
			return
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "access token is required",
			})
			return
		}

		// ====================================================
		// TODO:
		// claims, err := authService.ValidateAccessToken(token)
		// if err != nil {
		//     c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//         "message": "invalid access token",
		//     })
		//     return
		// }
		//
		// c.Set(UserIDContextKey, claims.UserID)
		// c.Set("session_id", claims.SessionID)
		// ====================================================

		c.Set(constant.TokenContextKey, token)
		// ctx := context.WithValue(c.Request.Context(), "Authorization", authHeader)
		// c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
