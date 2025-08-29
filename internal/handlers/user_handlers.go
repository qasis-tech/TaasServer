package handlers

import (
	"TaasServer/internal/services"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUserProfilePic(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.FormFile()
	}
}

func GetAllUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userList, err := services.GetAllUsers(db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, userList)
	}
}

func UpdateUserProfile(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
			return
		}

		var updateUser struct {
			Username string `json:"username"`
			EmailID  string `json:"email"`
		}

		if err := ctx.ShouldBindJSON(&updateUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		userId := ctx.GetUint("userID")

		log.Printf("User ID from payload , from DB %d ===> %d", id, userId)

		if userId != uint(id) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "user id didn't match",
			})
			return
		}

		resultUpadatedUser, err := services.UpdateUser(db, userId, updateUser.Username, updateUser.EmailID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error updating profile",
			})
			return
		}
		ctx.JSON(http.StatusOK, resultUpadatedUser)
	}
}

func GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
			return
		}

		user, err := services.GetUserById(db, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "User Not Found"})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
