package data

import (
	"context"

	"github.com/pro0o/second-life/gen/bowie/public/model"
	"github.com/pro0o/second-life/gen/bowie/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

// User operations implementation
func (repo *repository_Impl) CreateUser(ctx context.Context, email string, profilePicture int32, username string) (*model.Users, error) {
	id := uuid.New()
	user := &model.Users{
		ID:             id,
		Email:          &email,
		ProfilePicture: &profilePicture,
		Username:       &username,
	}

	err := table.Users.
		INSERT(table.Users.AllColumns).
		MODEL(user).
		RETURNING(table.Users.AllColumns).
		QueryContext(ctx, repo.DB, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repository_Impl) GetUserByEmail(ctx context.Context, email string) (*model.Users, error) {
	user := &model.Users{}

	err := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.Email.EQ(postgres.String(email))).
		QueryContext(ctx, repo.DB, user)

	if err != nil {
		if isNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil
}
