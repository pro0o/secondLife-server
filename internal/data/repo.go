package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/pro0o/second-life/gen/bowie/public/model"

	"github.com/google/uuid"
)

var (
	ErrNotFound           = errors.New("resource not found")
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrDuplicateOrgName   = errors.New("organization name already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Repository interface {
	CreateUser(ctx context.Context, email string, profilePicture int32, username string) (*model.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*model.Users, error)

	CreateOrg(ctx context.Context, userID uuid.UUID, orgName string, location string, description string) (*model.Orgs, error)
	GetOrgByName(ctx context.Context, orgName string) (*model.Orgs, error)

	UpdateUserPoints(ctx context.Context, userID uuid.UUID, points int32) (*model.RewardPoints, error)
	GetUserPointsByID(ctx context.Context, userID uuid.UUID) (*model.RewardPoints, error)
}

type repository_Impl struct {
	DB *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) Repository {
	return &repository_Impl{
		DB: db,
	}
}

func isNotFoundError(err error) bool {
	return err != nil && err.Error() == "sql: no rows in result set"
}
