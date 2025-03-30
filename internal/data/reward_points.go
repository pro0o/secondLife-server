package data

import (
	"context"

	"github.com/pro0o/second-life/gen/bowie/public/model"
	"github.com/pro0o/second-life/gen/bowie/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

func (repo *repository_Impl) UpdateUserPoints(ctx context.Context, userID uuid.UUID, points int32) (*model.RewardPoints, error) {
	_, err := repo.GetUserPointsByID(ctx, userID)

	if err != nil && err != ErrNotFound {
		return nil, err
	}

	rewardPoints := &model.RewardPoints{
		UserID: &userID,
		Points: &points,
	}

	if err == ErrNotFound {
		err = table.RewardPoints.
			INSERT(table.RewardPoints.AllColumns).
			MODEL(rewardPoints).
			RETURNING(table.RewardPoints.AllColumns).
			QueryContext(ctx, repo.DB, rewardPoints)
	} else {
		err = table.RewardPoints.
			UPDATE(table.RewardPoints.Points).
			MODEL(rewardPoints).
			WHERE(table.RewardPoints.UserID.EQ(postgres.UUID(userID))).
			RETURNING(table.RewardPoints.AllColumns).
			QueryContext(ctx, repo.DB, rewardPoints)
	}

	if err != nil {
		return nil, err
	}

	return rewardPoints, nil
}

func (repo *repository_Impl) GetUserPointsByID(ctx context.Context, userID uuid.UUID) (*model.RewardPoints, error) {
	rewardPoints := &model.RewardPoints{}

	err := table.RewardPoints.
		SELECT(table.RewardPoints.AllColumns).
		WHERE(table.RewardPoints.UserID.EQ(postgres.UUID(userID))).
		QueryContext(ctx, repo.DB, rewardPoints)

	if err != nil {
		if isNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return rewardPoints, nil
}
