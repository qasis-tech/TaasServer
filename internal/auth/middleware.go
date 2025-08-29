package auth

import (
	"TaasServer/pkg/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization Header",
			})
			ctx.Abort()
			return
		}

		log.Printf("Ping 1 %s", authHeader)

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header must start with 'Bearer '",
			})
			ctx.Abort()
			return
		}

		// Extract the token (remove "Bearer " prefix)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		log.Printf("tokenString ==> %s", tokenString)

		claims, err := utils.ValidateToken(tokenString)
		log.Printf("Validated %v", claims)

		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			ctx.Abort()
			return
		}

		log.Printf("Before Set ==> %d", claims.UserId)

		ctx.Set("userID", claims.UserId)
		ctx.Set("userName", claims.Username)
		ctx.Set("roles", claims.Role)

		ctx.Next()
	}
}
