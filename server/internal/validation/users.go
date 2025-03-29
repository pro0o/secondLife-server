package validation

import (
	storagev1 "secondLife/gen/bowie/v1"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateCreateUserRequest(req *storagev1.CreateUserRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			is.Email,
		),
		validation.Field(&req.Username,
			validation.Required,
			validation.Length(3, 30),
		),
		// Profile picture is optional, but if provided, validate URL format
		validation.Field(&req.ProfilePicture,
			validation.When(req.ProfilePicture != 0, is.Digit), // Validate it's a digit if provided
		),
	)
}

func ValidateGetUserByEmailRequest(req *storagev1.GetUserByEmailRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			is.Email,
		),
	)
}
