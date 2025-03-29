package transform

import (
	"secondLife/gen/bowie/public/model"
	storagev1 "secondLife/gen/bowie/v1"

	"github.com/google/uuid"
)

func User_InternalToV1(internal *model.Users) *storagev1.Users {
	if internal == nil {
		return nil
	}

	result := &storagev1.Users{
		Id: internal.ID.String(),
	}

	if internal.Email != nil {
		result.Email = *internal.Email
	}
	if internal.ProfilePicture != nil {
		result.ProfilePicture = *internal.ProfilePicture
	}
	if internal.Username != nil {
		result.Username = *internal.Username
	}

	return result
}

func User_InternalFromV1(v1 *storagev1.Users) *model.Users {
	if v1 == nil {
		return nil
	}

	id, _ := uuid.Parse(v1.Id)
	email := v1.Email
	profilePicture := v1.ProfilePicture
	username := v1.Username

	return &model.Users{
		ID:             id,
		Email:          &email,
		ProfilePicture: &profilePicture,
		Username:       &username,
	}
}

func CreateUserRequest_ToInternal(req *storagev1.CreateUserRequest) (string, int32, string) {
	return req.Email, req.ProfilePicture, req.Username
}

func CreateUserResponse_FromInternal(user *model.Users) *storagev1.CreateUserResponse {
	return &storagev1.CreateUserResponse{
		User: User_InternalToV1(user),
	}
}

func GetUserByEmailRequest_ToInternal(req *storagev1.GetUserByEmailRequest) string {
	return req.Email
}

func GetUserByEmailResponse_FromInternal(user *model.Users) *storagev1.GetUserByEmailResponse {
	return &storagev1.GetUserByEmailResponse{
		User: User_InternalToV1(user),
	}
}

func UpdateUserPointsRequest_ToInternal(req *storagev1.UpdateUserPointsRequest) (uuid.UUID, int32) {
	userID, _ := uuid.Parse(req.UserId)
	return userID, req.Points
}

func UpdateUserPointsResponse_FromInternal(points *model.RewardPoints) *storagev1.UpdateUserPointsResponse {
	return &storagev1.UpdateUserPointsResponse{
		RewardPoints: RewardPoints_InternalToV1(points),
	}
}
