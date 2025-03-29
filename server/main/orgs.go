package main

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	storagev1 "secondLife/gen/bowie/v1"
	"secondLife/internal/data"
	"secondLife/internal/transform"
	"secondLife/internal/validation"
)

func (app *application) CreateOrg(
	ctx context.Context,
	req *connect.Request[storagev1.CreateOrgRequest],
) (*connect.Response[storagev1.CreateOrgResponse], error) {
	err := validation.ValidateCreateOrgRequest(req.Msg)
	if err != nil {
		app.log.Error().Err(err).Msg("Invalid create org request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	userID, orgName, location, description := transform.CreateOrgRequest_ToInternal(req.Msg)

	org, err := app.repo.CreateOrg(ctx, userID, orgName, location, description)
	if err != nil {
		app.log.Error().Err(err).Msg("Failed to create organization")
		if errors.Is(err, data.ErrDuplicateOrgName) {
			return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("Organization name already exists"))
		}
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(transform.CreateOrgResponse_FromInternal(org)), nil
}

func (app *application) GetOrgByName(
	ctx context.Context,
	req *connect.Request[storagev1.GetOrgByNameRequest],
) (*connect.Response[storagev1.GetOrgByNameResponse], error) {
	err := validation.ValidateGetOrgByNameRequest(req.Msg)
	if err != nil {
		app.log.Error().Err(err).Msg("Invalid get org by name request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	orgName := transform.GetOrgByNameRequest_ToInternal(req.Msg)

	org, err := app.repo.GetOrgByName(ctx, orgName)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("Organization not found"))
		}
		app.log.Error().Err(err).Msg("Failed to get organization by name")
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(transform.GetOrgByNameResponse_FromInternal(org)), nil
}
