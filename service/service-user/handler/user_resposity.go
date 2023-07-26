package handler

import (
	"errors"
	"fmt"
	"github.com/bitxx/bitesla/common/errs"
	"github.com/bitxx/bitesla/common/logger"
	"github.com/bitxx/bitesla/common/util/idgenerate"
	"github.com/bitxx/bitesla/common/util/jwt"
	"github.com/bitxx/bitesla/service/service-user/conf"
	"github.com/bitxx/bitesla/service/service-user/db"
	"strconv"
	"time"
)

const (
	Facebook = "facebook"
	Phone    = "phone"
)

type userRepository struct {
}

func (r *userRepository) registerEmail(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	//生成id
	userId, err := idgenerate.GetId()
	if err != nil {
		return err
	}
	resp.UserId = userId
	err = db.AddUserByEmail(req.Email, req.Password, resp.UserId)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = generateToken(req, resp)
	return err
}

func (r *userRepository) registerPhone(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	return nil
}

func (r *userRepository) loginEmail(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	user, err := db.LoginUserByEmail(req.Email, req.Password)
	if err != nil || user == nil {
		if err != nil {
			return err
		}
		return errors.New(errs.GetMsg(errs.LoginErr))
	}
	resp.UserId = user.Id
	resp.Username = user.Username
	resp.Birthday = user.Birthday
	resp.Email = user.Email
	resp.Phone = user.Phone
	resp.Nickname = user.Nickname
	resp.Sex = int32(user.Sex)
	err = generateToken(req, resp)
	return err
}

func (r *userRepository) GetUserById(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	user, err := db.GetUserById(req.UserId)
	if err != nil {
		return errors.New(errs.GetMsg(errs.GetUserErr))
	}
	resp.UserId = user.Id
	resp.Username = user.Username
	resp.Birthday = user.Birthday
	resp.Email = user.Email
	resp.Phone = user.Phone
	resp.Nickname = user.Nickname
	resp.Sex = int32(user.Sex)
	return err
}

func (r *userRepository) loginPhone(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	fmt.Println("----------loginPhone")
	return nil
}

func (r *userRepository) getCode() {
	fmt.Println("----------getCode")
}

// 生成token
func generateToken(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	issuer := conf.CurrentConfig.ServerConf.JwtIssuer
	secret := conf.CurrentConfig.ServerConf.JwtSecret
	duration, err := strconv.Atoi(conf.CurrentConfig.ServerConf.JwtDuration)
	if err != nil {
		return err
	}
	dur := time.Duration(duration) * time.Hour
	s, err := jwt.GetToken(req.Email, req.Password, issuer, secret, resp.UserId, dur)
	if err != nil {
		return err
	}
	resp.Token = s
	return nil
}
