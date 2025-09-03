package services

import (
	"TaasServer/internal/models"
	"TaasServer/internal/repository"
	utils "TaasServer/pkg/utils"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(db *sql.DB, username, password string) (string, error) {
	user, err := repository.GetUserByUserName(db, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		cost, costErr := bcrypt.Cost([]byte(user.Password))
		if costErr != nil {
			return "", fmt.Errorf("invalid hash format: %w", costErr)
		}
		log.Printf("BCrypt cost parameter: %d", cost)
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJwtToken(user.Username, user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func UserRegistration(db *sql.DB, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password hashing failed: %w", err)
	}
	user.Password = string(hashedPassword)
	return repository.CreateUserRepo(db, user)
}
