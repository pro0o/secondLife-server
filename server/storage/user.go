package storage

import (
	"database/sql"
	"errors"
	"secondLife/types"

	"github.com/google/uuid"
)

func (s *PostgresStore) CreateUser(user *types.User) error {
	query := `
		INSERT INTO users(email, profile_picture, userName)
		VALUES($1, $2, $3)
	`
	_, err := s.DB.Exec(
		query,
		user.Email,
		user.ProfilePicture,
		user.UserName)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetUserUUID(email string) (uuid.UUID, error) {
	var userID uuid.UUID
	err := s.DB.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, errors.New("user not found")
		}
		return uuid.Nil, err
	}
	return userID, nil
}

func (s *PostgresStore) CheckEmailExists(email string) bool {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"
	var exists bool
	err := s.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	query := "SELECT id, profile_picture, userName FROM users WHERE email = $1"
	var user types.User
	err := s.DB.QueryRow(query, email).Scan(&user.UserID, &user.ProfilePicture, &user.UserName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
