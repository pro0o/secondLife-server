package validation

import (
	storagev1 "github.com/pro0o/second-life/gen/bowie/v1"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateSignUpRequest(req *storagev1.SignUpRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			is.Email,
		),
		validation.Field(&req.Password,
			validation.Required,
			validation.Length(8, 50),
		),
		validation.Field(&req.Username,
			validation.Required,
			validation.Length(3, 30),
		),
	)
}

func ValidateLoginRequest(req *storagev1.LoginRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			is.Email,
		),
		validation.Field(&req.Password,
			validation.Required,
			validation.Length(1, 50),
		),
	)
}
