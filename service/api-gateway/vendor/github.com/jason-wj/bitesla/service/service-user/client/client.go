package client

import (
	"context"
	"encoding/json"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/common/util/stringutl"
	pb "github.com/jason-wj/bitesla/service/service-user/proto"
	"github.com/micro/go-micro/client"
)

type UserClient struct {
	client pb.UserService
}

func NewUserClient() *UserClient {
	c := pb.NewUserService("", client.DefaultClient)
	return &UserClient{
		client: c,
	}
}

//邮箱注册
func (client *UserClient) RegisterEmail(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	if !stringutl.VerifyEmailFormat(userReq.Email) {
		return nil, errs.EmailErr, nil
	}
	if userReq.Password == "" {
		return nil, errs.PwdEmptyErr, nil
	}

	resp, err := client.client.RegisterEmail(ctx, &pb.UserReq{
		Email:    userReq.Email,
		Password: userReq.Password,
	})
	if err != nil {
		return nil, errs.RegisterErr, err
	}
	return resp, errs.Success, nil
}

//手机号注册
func (client *UserClient) RegisterPhone(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	resp, err := client.client.RegisterPhone(ctx, &pb.UserReq{
		Phone: userReq.Phone,
	})
	return resp, 0, err
}

//邮箱登录
func (client *UserClient) LoginEmail(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	if userReq.Email == "" {
		return nil, errs.EmailErr, nil
	}
	if userReq.Password == "" {
		return nil, errs.PwdEmptyErr, nil
	}

	resp, err := client.client.LoginEmail(ctx, &pb.UserReq{
		Email:    userReq.Email,
		Password: userReq.Password,
	})

	if err != nil {
		return nil, errs.LoginErr, err
	}
	return resp, errs.Success, nil
}

//手机号登录
func (client *UserClient) LoginPhone(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	resp, err := client.client.LoginPhone(ctx, &pb.UserReq{
		Phone: userReq.Phone,
	})
	return resp, 0, err
}
