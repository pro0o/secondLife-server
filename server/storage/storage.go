package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
)

type PostgresStore struct {
	DB *sql.DB
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
		return
	}
}

func (s *PostgresStore) Init() error {
	return s.createTables()
}

func NewDB() (*PostgresStore, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", getEnv(dbHost), getEnv(dbPort), getEnv(dbUser), getEnv(dbName), getEnv(dbPassword))
	log.Println("the connection string is:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (s *PostgresStore) createTables() error {
	query := `
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
				CREATE TYPE role AS ENUM('user', 'admin');
			END IF;
		END
		$$;		

		CREATE TABLE IF NOT EXISTS users (
		    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		    email TEXT,
		    profile_picture INT, 
		    userName VARCHAR(15)
		);

		CREATE TABLE IF NOT EXISTS orgs (
		    user_id UUID REFERENCES users(id) ON DELETE CASCADE, 
		    orgName VARCHAR(15),
			Location VARCHAR(50),
			description VARCHAR(100)
		);

		CREATE TABLE IF NOT EXISTS RewardPoints (
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			Points INT NOT NULL DEFAULT 0
		);
	`

	_, err := s.DB.Exec(query)
	return err
}
