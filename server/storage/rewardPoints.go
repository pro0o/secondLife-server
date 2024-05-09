package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (s *PostgresStore) UpdateUserPoints(userID uuid.UUID, points int) error {
	existsQuery := "SELECT COUNT(*) FROM RewardPoints WHERE user_id = $1"
	var count int
	err := s.DB.QueryRow(existsQuery, userID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			insertQuery := "INSERT INTO RewardPoints (user_id, points) VALUES ($1, $2)"
			_, err := s.DB.Exec(insertQuery, userID, points)
			if err != nil {
				return fmt.Errorf("failed to insert user points: %w", err)
			}
			log.Printf("Inserted new user with ID %s and points %d", userID, points)
			return nil
		} else {
			return fmt.Errorf("failed to check user entry: %w", err)
		}
	}

	if count == 0 {
		insertQuery := "INSERT INTO RewardPoints (user_id, points) VALUES ($1, $2)"
		_, err := s.DB.Exec(insertQuery, userID, points)
		if err != nil {
			return fmt.Errorf("failed to insert user points: %w", err)
		}
		log.Printf("Inserted new user with ID %s and points %d", userID, points)
		return nil
	}

	updateQuery := "UPDATE RewardPoints SET points = $1 WHERE user_id = $2"
	_, err = s.DB.Exec(updateQuery, points, userID)
	if err != nil {
		return fmt.Errorf("failed to update user points: %w", err)
	}

	log.Printf("Updated points for user with ID %s to %d", userID, points)
	return nil
}

func (s *PostgresStore) GetUserPointsByID(userID uuid.UUID) (int, error) {
	query := "SELECT points FROM RewardPoints WHERE user_id = $1"
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
