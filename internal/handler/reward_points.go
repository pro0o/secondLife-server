package handler

import (
	"context"
	"errors"

	storagev1 "secondLife/gen/bowie/v1"
	"secondLife/internal/data"
	"secondLife/internal/transform"
	"secondLife/internal/validation"

	"connectrpc.com/connect"
)

func (h *Handler) UpdateUserPoints(
	ctx context.Context,
	req *connect.Request[storagev1.UpdateUserPointsRequest],
) (*connect.Response[storagev1.UpdateUserPointsResponse], error) {
	err := validation.ValidateUpdateUserPointsRequest(req.Msg)
	if err != nil {
		h.Log.Error().Err(err).Ctx(ctx).Msg("invalid update user points request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	userID, points := transform.UpdateUserPointsRequest_ToInternal(req.Msg)

	rewardPoints, err := h.Repo.UpdateUserPoints(ctx, userID, points)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
		h.Log.Error().Err(err).Ctx(ctx).Msg("failed to update user points")
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(transform.UpdateUserPointsResponse_FromInternal(rewardPoints)), nil
}

func (h *Handler) GetUserPointsByID(
	ctx context.Context,
	req *connect.Request[storagev1.GetUserPointsByIDRequest],
) (*connect.Response[storagev1.GetUserPointsByIDResponse], error) {
	err := validation.ValidateGetUserPointsByIDRequest(req.Msg)
	if err != nil {
		h.Log.Error().Err(err).Ctx(ctx).Msg("invalid get user points request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	userID := transform.GetUserPointsByIDRequest_ToInternal(req.Msg)

	rewardPoints, err := h.Repo.GetUserPointsByID(ctx, userID)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user points not found"))
		}
		h.Log.Error().Err(err).Ctx(ctx).Msg("failed to get user points")
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(transform.GetUserPointsByIDResponse_FromInternal(rewardPoints)), nil
}
