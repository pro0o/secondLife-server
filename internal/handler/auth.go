package handler

import (
	"context"

	"connectrpc.com/connect"

	storagev1 "github.com/pro0o/second-life/gen/bowie/v1"
	"github.com/pro0o/second-life/internal/transform"
	"github.com/pro0o/second-life/internal/validation"
)

func (h *Handler) SignUp(
	ctx context.Context,
	req *connect.Request[storagev1.SignUpRequest],
) (*connect.Response[storagev1.SignUpResponse], error) {
	err := validation.ValidateSignUpRequest(req.Msg)
	if err != nil {
		h.Log.Error().Err(err).Ctx(ctx).Msg("Invalid signup request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// email, password, username := transform.SignUpRequest_ToInternal(req.Msg)

	// user, err := h.Repo.CreateUser(ctx, email, password, username)
	// if err != nil {
	// 	h.Log.Error().Err(err).Ctx(ctx).Msg("Failed to create user during signup")
	// 	if errors.Is(err, data.ErrDuplicateEmail) {
	// 		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("email already registered"))
	// 	}
	// 	return nil, connect.NewError(connect.CodeInternal, nil)
	// }

	// // Generate tokens
	// accessToken, refreshToken, err := app.tokenService.GenerateTokens(user.ID.String())
	// if err != nil {
	// 	h.Log.Error().Err(err).Ctx(ctx).Msg("Failed to generate auth tokens")
	// 	return nil, connect.NewError(connect.CodeInternal, nil)
	// }

	return connect.NewResponse(transform.SignUpResponse_FromInternal("accessToken", "refreshToken")), nil
}

func (h *Handler) Login(
	ctx context.Context,
	req *connect.Request[storagev1.LoginRequest],
) (*connect.Response[storagev1.LoginResponse], error) {
	err := validation.ValidateLoginRequest(req.Msg)
	if err != nil {
		h.Log.Error().Err(err).Ctx(ctx).Msg("Invalid login request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// email, password := transform.LoginRequest_ToInternal(req.Msg)

	// user, err := h.Repo.AuthenticateUser(ctx, email, password)
	// if err != nil {
	// 	h.Log.Error().Err(err).Ctx(ctx).Msg("Authentication failed")
	// 	if errors.Is(err, data.ErrInvalidCredentials) {
	// 		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid email or password"))
	// 	}
	// 	return nil, connect.NewError(connect.CodeInternal, nil)
	// }

	// // Generate tokens
	// accessToken, refreshToken, err := app.tokenService.GenerateTokens(user.ID.String())
	// if err != nil {
	// 	h.Log.Error().Err(err).Ctx(ctx).Msg("Failed to generate auth tokens")
	// 	return nil, connect.NewError(connect.CodeInternal, nil)
	// }

	return connect.NewResponse(transform.LoginResponse_FromInternal("accessToken", "refreshToken")), nil
}
