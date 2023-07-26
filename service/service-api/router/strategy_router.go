package router

import (
	"github.com/bitxx/bitesla/common/errs"
	"github.com/bitxx/bitesla/common/logger"
	"github.com/bitxx/bitesla/service/service-strategy/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	strategyPutUrl    = "/strategy/put"
	strategyListUrl   = "/strategy/list"
	strategyDetailUrl = "/strategy/detail"
)

func strategyRouter(router *gin.Engine) {
	router.POST(strategyPutUrl, strategyPut)
	router.POST(strategyListUrl, strategyList)
	router.POST(strategyDetailUrl, strategyDetail)
}

var (
	strategyClient = client.NewStrategyClient()
)

// @Summary 新增或者更新一个策略
// @Description 新增或者更新一个策略
// @Tags 策略管理相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.StrategyInfo true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /strategy/put [post]
func strategyPut(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := strategyClient.PutStrategy(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 当前用户拥有的策略，分页展示
// @Description 当前用户拥有的策略，分页展示
// @Tags 策略管理相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.StrategyList true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /strategy/list [post]
func strategyList(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := strategyClient.ListStrategy(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 当前策略详情
// @Description 当前策略详情
// @Tags 策略管理相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.StrategyDetail true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /strategy/detail [post]
func strategyDetail(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := strategyClient.GetStrategyDetail(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}
