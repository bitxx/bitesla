package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/vo"
	"github.com/jason-wj/bitesla/service/service-trader/client"
	"net/http"
)

const (
	traderPutUrl    = "/trader/put"
	traderListUrl   = "/trader/list"
	traderDetailUrl = "/trader/detail"
	traderSwitchUrl = "/trader/switch"
)

func traderRouter(router *gin.Engine) {
	router.POST(traderPutUrl, traderPut)
	router.POST(traderListUrl, traderList)
	router.POST(traderDetailUrl, traderDetail)
	router.POST(traderSwitchUrl, traderSwitch)
}

var (
	traderClient = client.NewTraderClient()
)

// @Summary 新增或者更新一个要执行的策略
// @Description 新增或者更新一个要执行的策略
// @Tags 策略执行相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.TraderPut true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/put [post]
func traderPut(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.PutTrader(reqData)
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
// @Tags 策略执行相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.TraderList true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/list [post]
func traderList(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.ListTrader(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 当前所执行的策略状况详情
// @Description 当前所执行的策略状况详情
// @Tags 策略执行相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.TraderDetail true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/detail [post]
func traderDetail(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetTraderDetail(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 执行指定策略
// @Description 执行指定策略
// @Tags 策略执行相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.TraderSwitch true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/switch [post]
func traderSwitch(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.SwitchTrader(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}
