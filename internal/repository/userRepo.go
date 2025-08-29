package repository

import (
	"TaasServer/internal/models"
	"database/sql"
	"fmt"
	"log"
)

func GetAllUsers(db *sql.DB) ([]*models.User, error) {
	query := `SELECT id,username,email FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(db *sql.DB, user *models.User) (*models.User, error) {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3 `
	log.Printf("%s %s %v 11111", user.Username, user.Email, user.ID)
	_, err := db.Exec(query, user.Username, user.Email, user.ID)

	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}
	return user, nil
}

func GetUserByUserName(db *sql.DB, username string) (*models.User, error) {
	query := `SELECT * FROM users WHERE username = $1`
	user := &models.User{}
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	log.Printf("Result ==>", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}

func CreateUserRepo(db *sql.DB, user *models.User) error {
	query := `
        INSERT INTO users (username, password, email)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
	err := db.QueryRow(
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt) // Scan the returned values

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
