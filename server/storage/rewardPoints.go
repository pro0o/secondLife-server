package storage

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

func (s *PostgresStore) UpdateUserPoints(userID uuid.UUID, points int) error {
	query := "UPDATE RewardPoints SET points = $1 WHERE user_id = $2"

	_, err := s.DB.Exec(query, points, userID)
	if err != nil {
		return errors.New("failed to update user points: " + err.Error())
	}

	return nil
}

func (s *PostgresStore) GetUserPointsByID(userID uuid.UUID) (int, error) {
	query := "SELECT Points FROM RewardPoints WHERE user_id = $1"
	var points int
	err := s.DB.QueryRow(query, userID).Scan(&points)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return points, nil
}
