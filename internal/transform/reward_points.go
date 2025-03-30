package transform

import (
	"github.com/pro0o/second-life/gen/bowie/public/model"
	storagev1 "github.com/pro0o/second-life/gen/bowie/v1"

	"github.com/google/uuid"
)

func RewardPoints_InternalToV1(internal *model.RewardPoints) *storagev1.RewardPoints {
	if internal == nil {
		return nil
	}

	result := &storagev1.RewardPoints{}

	if internal.UserID != nil {
		result.UserId = internal.UserID.String()
	}
	if internal.Points != nil {
		result.Points = *internal.Points
	}

	return result
}

func RewardPoints_InternalFromV1(v1 *storagev1.RewardPoints) *model.RewardPoints {
	if v1 == nil {
		return nil
	}

	userID, _ := uuid.Parse(v1.UserId)
	points := v1.Points

	return &model.RewardPoints{
		UserID: &userID,
		Points: &points,
	}
}

func GetUserPointsByIDRequest_ToInternal(req *storagev1.GetUserPointsByIDRequest) uuid.UUID {
	userID, _ := uuid.Parse(req.UserId)
	return userID
}

func GetUserPointsByIDResponse_FromInternal(points *model.RewardPoints) *storagev1.GetUserPointsByIDResponse {
	return &storagev1.GetUserPointsByIDResponse{
		RewardPoints: RewardPoints_InternalToV1(points),
	}
}
