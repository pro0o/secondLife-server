package storage

import "secondLife/types"

func (s *PostgresStore) CreateUser(user *types.User) error {
	query := `
		INSERT INTO users(email, profile_picture, userName, encrypted_password)
		VALUES($1, $2, $3, $4)
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

func (s *PostgresStore) CheckEmailExists(email string) bool {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"
	var exists bool
	err := s.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (s *PostgresStore) CheckPassword(email, password string) bool {
	query := "SELECT encrypted_password FROM users WHERE email = $1"
	var storedPassword string
	err := s.DB.QueryRow(query, email).Scan(&storedPassword)
	if err != nil {
		return false
	}
	return storedPassword == password
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	query := "SELECT profile_picture, userName FROM users WHERE email = $1"
	var user types.User
	err := s.DB.QueryRow(query, email).Scan(&user.ProfilePicture, &user.UserName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
