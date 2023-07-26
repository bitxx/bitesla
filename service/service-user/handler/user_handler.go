package handler

import (
	"context"
)

type UserHandler struct {
	repo *userRepository
}

func NewUserHandler() *UserHandler {
	repository := &userRepository{}
	handler := &UserHandler{
		repo: repository,
	}
	return handler
}

func (user *UserHandler) GetUserById(ctx context.Context, req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	return user.repo.GetUserById(req, resp)
}

func (user *UserHandler) LoginEmail(ctx context.Context, req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	err := user.repo.loginEmail(req, resp)
	return err
}

func (user *UserHandler) LoginPhone(ctx context.Context, req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	err := user.repo.loginPhone(req, resp)
	return err
}

func (user *UserHandler) GetCode(ctx context.Context, req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	user.repo.getCode()
	return nil
}

func (user *UserHandler) RegisterEmail(ctx context.Context, req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	err := user.repo.registerEmail(req, resp)
	return err
}

func (user *UserHandler) RegisterPhone(ctx context.Context, req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	err := user.repo.registerPhone(req, resp)
	return err
}
