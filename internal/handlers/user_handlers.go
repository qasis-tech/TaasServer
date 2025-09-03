package handlers

import (
	"TaasServer/internal/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUserProfilePic(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
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

		file, err := ctx.FormFile("profile-pic")
		log.Printf("Files ==== %v %v", file, err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Error Uploading the File"})
			return
		}

		if err := os.MkdirAll(os.Getenv("UPLOAD_DIR"), os.ModePerm); err != nil {
			log.Printf("Error creating upload dir: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Creating upload directory failed"})
			return
		}

		fileName := fmt.Sprintf("%d-%s", id, filepath.Base(file.Filename))
		filePath := filepath.Join(os.Getenv("UPLOAD_DIR"), fileName)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			log.Printf("Error saving uploaded file: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Saving uploaded file failed"})
			return
		}

		err1 := services.UpdateProfilePic(db, userId, filePath)
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": "Error updating profile pics",
			})
			return
		}

		ctx.JSON(http.StatusOK, "Profile update successs")

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
