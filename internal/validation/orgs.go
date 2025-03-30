package validation

import (
	storagev1 "github.com/pro0o/second-life/gen/bowie/v1"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateCreateOrgRequest(req *storagev1.CreateOrgRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.UserId,
			validation.Required,
			is.UUID,
		),
		validation.Field(&req.OrgName,
			validation.Required,
			validation.Length(2, 50),
		),
		validation.Field(&req.Location,
			validation.Required,
			validation.Length(2, 100),
		),
		validation.Field(&req.Description,
			validation.Length(0, 500), // Optional but with max length
		),
	)
}

func ValidateGetOrgByNameRequest(req *storagev1.GetOrgByNameRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.OrgName,
			validation.Required,
			validation.Length(2, 50),
		),
	)
}
