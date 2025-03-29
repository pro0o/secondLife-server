package main

import (
	"context"
	"errors"

	storagev1 "secondLife/gen/bowie/v1"
	"secondLife/internal/data"
	"secondLife/internal/transform"
	"secondLife/internal/validation"

	"connectrpc.com/connect"
)

func (app *application) CreateUser(
	ctx context.Context,
	req *connect.Request[storagev1.CreateUserRequest],
) (*connect.Response[storagev1.CreateUserResponse], error) {
	err := validation.ValidateCreateUserRequest(req.Msg)
	if err != nil {
		app.log.Error().Err(err).Ctx(ctx).Msg("invalid create user request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	email, profilePicture, username := transform.CreateUserRequest_ToInternal(req.Msg)

	user, err := app.repo.CreateUser(ctx, email, int32(profilePicture), username)
	if err != nil {
		app.log.Error().Err(err).Ctx(ctx).Msg("failed to create user")
		if errors.Is(err, data.ErrDuplicateEmail) {
			return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("email already exists"))
		}
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(transform.CreateUserResponse_FromInternal(user)), nil
}

func (app *application) GetUserByEmail(
	ctx context.Context,
	req *connect.Request[storagev1.GetUserByEmailRequest],
) (*connect.Response[storagev1.GetUserByEmailResponse], error) {
	err := validation.ValidateGetUserByEmailRequest(req.Msg)
	if err != nil {
		app.log.Error().Err(err).Ctx(ctx).Msg("invalid get user by email request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	email := transform.GetUserByEmailRequest_ToInternal(req.Msg)

	user, err := app.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
		}
		app.log.Error().Err(err).Ctx(ctx).Msg("failed to get user by email")
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(transform.GetUserByEmailResponse_FromInternal(user)), nil
}
