package transform

import (
	"secondLife/gen/bowie/public/model"
	storagev1 "secondLife/gen/bowie/v1"

	"github.com/google/uuid"
)

// Organization transformations
func Org_InternalToV1(internal *model.Orgs) *storagev1.Orgs {
	if internal == nil {
		return nil
	}

	result := &storagev1.Orgs{}

	if internal.UserID != nil {
		result.UserId = internal.UserID.String()
	}
	if internal.OrgName != nil {
		result.OrgName = *internal.OrgName
	}
	if internal.Location != nil {
		result.Location = *internal.Location
	}
	if internal.Description != nil {
		result.Description = *internal.Description
	}

	return result
}

func Org_InternalFromV1(v1 *storagev1.Orgs) *model.Orgs {
	if v1 == nil {
		return nil
	}

	userID, _ := uuid.Parse(v1.UserId)
	orgName := v1.OrgName
	location := v1.Location
	description := v1.Description

	return &model.Orgs{
		UserID:      &userID,
		OrgName:     &orgName,
		Location:    &location,
		Description: &description,
	}
}

func CreateOrgRequest_ToInternal(req *storagev1.CreateOrgRequest) (uuid.UUID, string, string, string) {
	userID, _ := uuid.Parse(req.UserId)
	return userID, req.OrgName, req.Location, req.Description
}

func CreateOrgResponse_FromInternal(org *model.Orgs) *storagev1.CreateOrgResponse {
	return &storagev1.CreateOrgResponse{
		Org: Org_InternalToV1(org),
	}
}

func GetOrgByNameRequest_ToInternal(req *storagev1.GetOrgByNameRequest) string {
	return req.OrgName
}

func GetOrgByNameResponse_FromInternal(org *model.Orgs) *storagev1.GetOrgByNameResponse {
	return &storagev1.GetOrgByNameResponse{
		Org: Org_InternalToV1(org),
	}
}
