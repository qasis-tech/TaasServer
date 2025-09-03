package main

import (
	"TaasServer/internal/routes"
	"TaasServer/pkg/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	routes.InitRoutes(r, db)

	defer db.Close()
	r.Run(":" + port)

}
