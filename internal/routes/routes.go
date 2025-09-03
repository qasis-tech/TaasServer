package routes

import (
	"TaasServer/internal/auth"
	"TaasServer/internal/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("/login", handlers.LoginHandler(db))
	r.POST("/register", handlers.RegisterHandler(db))

	authenticated := r.Group("/")
	authenticated.Use(auth.AuthMiddleware())
	authenticated.GET("/users", handlers.GetAllUsers(db))
	authenticated.GET("/users/:id", handlers.GetUserByIdHandler(db))
	authenticated.PUT("/users/:id", handlers.UpdateUserProfile(db))
	authenticated.PUT("/users/:id/update-profile-pic", handlers.UpdateUserProfilePic(db))
}
