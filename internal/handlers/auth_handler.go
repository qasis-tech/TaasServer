package handlers

import (
	"TaasServer/internal/models"
	"TaasServer/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Move user variable inside the handler (important for thread safety)
		var user models.User

		// Bind JSON with better error handling
		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Printf("Invalid login request: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request format",
				"details": err.Error(),
			})
			return
		}

		token, err := services.LoginUser(db, user.Username, user.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid credentials",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := services.UserRegistration(db, &user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}
