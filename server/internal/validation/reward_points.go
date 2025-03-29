package validation

import (
	storagev1 "secondLife/gen/bowie/v1"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateUpdateUserPointsRequest(req *storagev1.UpdateUserPointsRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.UserId,
			validation.Required,
			is.UUID,
		),
		validation.Field(&req.Points,
			validation.Required,
		),
	)
}

func ValidateGetUserPointsByIDRequest(req *storagev1.GetUserPointsByIDRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.UserId,
			validation.Required,
			is.UUID,
		),
	)
}
