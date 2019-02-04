package handler

import (
	"errors"
	"fmt"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/util/jwt"
	"github.com/jason-wj/bitesla/service/service-user/conf"
	"github.com/jason-wj/bitesla/service/service-user/db"
	"github.com/jason-wj/bitesla/service/service-user/proto"
	"strconv"
	"time"
)

const (
	Facebook = "facebook"
	Phone    = "phone"
)

type userRepository struct {
	DB *db.ConnectPool
}

func (r *userRepository) registerEmail(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	err := db.AddUserByEmail(req.Email, req.Password)
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
	count := db.LoginUserByEmail(req.Email, req.Password)
	if count != 1 {
		return errors.New(errs.GetMsg(errs.LoginErr))
	}
	err := generateToken(req, resp)
	return err
}

func (r *userRepository) loginPhone(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	fmt.Println("----------loginPhone")
	return nil
}

func (r *userRepository) getCode() {
	fmt.Println("----------getCode")
}

//生成token
func generateToken(req *bitesla_srv_user.UserReq, resp *bitesla_srv_user.UserResp) error {
	issuer := conf.CurrentConfig.ServerConf.JwtIssuer
	secret := conf.CurrentConfig.ServerConf.JwtSecret
	duration, err := strconv.Atoi(conf.CurrentConfig.ServerConf.JwtDuration)
	if err != nil {
		return err
	}
	dur := time.Duration(duration) * time.Hour
	s, err := jwt.GetToken(req.Email, req.Password, issuer, secret, dur)
	if err != nil {
		return err
	}
	resp.Token = s
	return nil
}
