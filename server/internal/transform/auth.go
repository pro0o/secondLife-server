package transform

import (
	storagev1 "secondLife/gen/bowie/v1"
)

func SignUpRequest_ToInternal(req *storagev1.SignUpRequest) (string, string, string) {
	return req.Email, req.Password, req.Username
}

func SignUpResponse_FromInternal(accessToken, refreshToken string) *storagev1.SignUpResponse {
	return &storagev1.SignUpResponse{
		// TODO: authentication
		// User:         User_InternalToV1(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func LoginRequest_ToInternal(req *storagev1.LoginRequest) (string, string) {
	return req.Email, req.Password
}

func LoginResponse_FromInternal(accessToken, refreshToken string) *storagev1.LoginResponse {
	return &storagev1.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
