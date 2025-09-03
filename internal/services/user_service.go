package services

import (
	"TaasServer/internal/models"
	"TaasServer/internal/repository"
	"database/sql"
)

func UpdateProfilePic(db *sql.DB, id uint, fileName string) error {
	return repository.UpdateProfilePic(db, id, fileName)
}

func GetAllUsers(db *sql.DB) ([]*models.User, error) {
	return repository.GetAllUsers(db)
}

func UpdateUser(db *sql.DB, id uint, username, email string) (*models.User, error) {
	user := &models.User{ID: id, Username: username, Email: email}
	return repository.UpdateUser(db, user)
}

func GetUserById(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserById(db, id)
}
