package handler

import (
	"context"

	account "disk-server/account/proto"
)

type User struct{
	account.UnimplementedUserServiceServer
}

func (u *User) SignUp(ctx context.Context,  req *account.UserSignUpRequest) (*account.UserSignUpResponse, error) {
	return &account.UserSignUpResponse{}, nil
}

func (u *User) SignIn(ctx context.Context, req *account.UserSignInRequest) (*account.UserSignInResponse, error) {
	return &account.UserSignInResponse{}, nil
}