package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/vo"
	"github.com/jason-wj/bitesla/service/service-user/client"
	"net/http"
)

const (
	registerEmailUrl = "/user/registerEmail"
	registerPhoneUrl = "/user/registerPhone"
	loginEmailUrl    = "/user/loginEmail"
	loginPhoneUrl    = "/user/loginPhone"
)

func userRouterNoAuth(router *gin.Engine) {
	router.POST(registerEmailUrl, registerEmail)
	router.POST(registerPhoneUrl, registerPhone)
	router.POST(loginEmailUrl, loginEmail)
	router.POST(loginPhoneUrl, loginPhone)
}

func userRouter(router *gin.Engine) {
}

var (
	userClient = client.NewUserClient()
)

// @Summary 邮箱注册
// @Description 邮箱注册
// @Tags 账户操作
// @Accept   json
// @Produce   json
// @Param group body model.EmailRegister true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /user/registerEmail [post]
func registerEmail(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := userClient.RegisterEmail(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("邮箱注册请求错误，错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 手机号注册（暂未实现）
// @Description 手机号注册
// @Tags 账户操作
// @Accept   json
// @Produce   json
// @Param group body model.PhoneRegister true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /user/registerPhone [post]
func registerPhone(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := userClient.RegisterPhone(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 邮箱登陆
// @Description 邮箱登陆
// @Tags 账户操作
// @Accept   json
// @Produce   json
// @Param group body model.EmailLogin true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /user/loginEmail [post]
func loginEmail(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := userClient.LoginEmail(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 手机号登陆（暂未实现）
// @Description 手机号登陆（暂未实现）
// @Tags 账户操作
// @Accept   json
// @Produce   json
// @Param group body model.PhoneLogin true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /user/loginPhone [post]
func loginPhone(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := userClient.LoginPhone(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}
