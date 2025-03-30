package data

import (
	"context"
	"secondLife/gen/bowie/public/model"
	"secondLife/gen/bowie/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

func (repo *repository_Impl) CreateOrg(ctx context.Context, userID uuid.UUID, orgName string, location string, description string) (*model.Orgs, error) {
	org := &model.Orgs{
		UserID:      &userID,
		OrgName:     &orgName,
		Location:    &location,
		Description: &description,
	}

	err := table.Orgs.
		INSERT(table.Orgs.AllColumns).
		MODEL(org).
		RETURNING(table.Orgs.AllColumns).
		QueryContext(ctx, repo.DB, org)

	if err != nil {
		return nil, err
	}

	return org, nil
}

func (repo *repository_Impl) GetOrgByName(ctx context.Context, orgName string) (*model.Orgs, error) {
	org := &model.Orgs{}

	err := table.Orgs.
		SELECT(table.Orgs.AllColumns).
		WHERE(table.Orgs.OrgName.EQ(postgres.String(orgName))).
		QueryContext(ctx, repo.DB, org)

	if err != nil {
		if isNotFoundError(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return org, nil
}
