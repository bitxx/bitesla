package client

import (
	"context"
	"encoding/json"
	"github.com/bitxx/bitesla/common/errs"
	"github.com/bitxx/bitesla/common/util/stringutl"
	pb "github.com/bitxx/bitesla/service/service-user/proto"
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

// 邮箱注册
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

	resp, err := client.client.RegisterEmail(context.Background(), &pb.UserReq{
		Email:    userReq.Email,
		Password: userReq.Password,
	})
	if err != nil {
		return nil, errs.RegisterErr, err
	}
	return resp, errs.Success, nil
}

// 手机号注册
func (client *UserClient) RegisterPhone(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	resp, err := client.client.RegisterPhone(context.Background(), &pb.UserReq{
		Phone: userReq.Phone,
	})
	return resp, 0, err
}

// 邮箱登录
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
	resp, err := client.client.LoginEmail(context.Background(), &pb.UserReq{
		Email:    userReq.Email,
		Password: userReq.Password,
	})

	if err != nil {
		return nil, errs.LoginErr, err
	}
	return resp, errs.Success, nil
}

// 手机号登录
func (client *UserClient) LoginPhone(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	resp, err := client.client.LoginPhone(context.Background(), &pb.UserReq{
		Phone: userReq.Phone,
	})
	return resp, 0, err
}

func (client *UserClient) GetUserById(data []byte) (interface{}, int, error) {
	userReq := &pb.UserReq{}
	err := json.Unmarshal(data, userReq)
	if err != nil {
		return nil, errs.RequestDataFmtErr, err
	}

	if userReq.UserId <= 0 {
		return nil, errs.UserIdErr, err
	}

	resp, err := client.client.GetUserById(context.Background(), &pb.UserReq{
		UserId: userReq.UserId,
	})
	return resp, errs.Success, err
}
