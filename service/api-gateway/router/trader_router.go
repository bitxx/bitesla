package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/vo"
	"github.com/jason-wj/bitesla/service/service-exchange/client"
	"net/http"
)

const (
	orderPlaceUrl        = "/trader/orderPlace"
	cancelOrderUrl       = "/trader/cancelOrder"
	getOneOrderUrl       = "/trader/getOneOrder"
	getUnfinishOrdersUrl = "/trader/getUnfinishOrders"
	getOrderHistorysUrl  = "/trader/getOrderHistorys"
	getAccountUrl        = "/trader/getAccount"
	getTickerUrl         = "/trader/getTicker"
	getDepthUrl          = "/trader/getDepth"
	getKlineRecordsUrl   = "/trader/getKlineRecords"
	getTradesUrl         = "/trader/getTrades"
	getExchangeNameUrl   = "/trader/getExchangeName"
)

func traderRouter(router *gin.Engine) {
	router.POST(orderPlaceUrl, orderPlace)
	router.POST(cancelOrderUrl, cancelOrder)
	router.POST(getOneOrderUrl, getOneOrder)
	router.POST(getUnfinishOrdersUrl, getUnfinishOrders)
	router.POST(getOrderHistorysUrl, getOrderHistorys)
	router.POST(getAccountUrl, getAccount)
	router.POST(getTickerUrl, getTicker)
	router.POST(getDepthUrl, getDepth)
	router.POST(getKlineRecordsUrl, getKlineRecords)
	router.POST(getTradesUrl, getTrades)
	router.POST(getExchangeNameUrl, getExchangeName)
}

var (
	traderClient = client.NewTraderClient()
)

// @Summary 发送一个新的订单到某交易所进行撮合
// @Description 发送一个新的订单到某交易所进行撮合
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderPlace true "每个参数均不得为空,//	OrderType: 0 表示limitBuy 1 表示limitSell 2 表示marketBuy 3 表示marketSell ; AccountType: 1 表示point 2 表示splot"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/orderPlace [post]
func orderPlace(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.OrderPlace(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 此接口发送一个撤销订单的请求
// @Description 此接口发送一个撤销订单的请求
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderBase true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/cancelOrder [post]
func cancelOrder(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.CancelOrder(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 返回订单的最新状况和详情
// @Description 返回订单的最新状况和详情
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderBase true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getOneOrder [post]
func getOneOrder(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetOneOrder(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 未完成的交易订单
// @Description 未完成的交易订单
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderBase2 true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getUnfinishOrders [post]
func getUnfinishOrders(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetUnfinishOrders(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 获取某交易对交易历史
// @Description 获取某交易对交易历史
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderHistory true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getOrderHistorys [post]
func getOrderHistorys(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetOrderHistorys(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 获取当前账户信息
// @Description 获取当前账户信息
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.TokenAuth true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getAccount [post]
func getAccount(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetAccount(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 获取ticker信息，同时提供24小时交易聚合信息
// @Description 获取ticker信息，同时提供24小时交易聚合信息
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderBase2 true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getTicker [post]
func getTicker(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetTicker(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 返回指定交易对的当前市场深度数据
// @Description 返回指定交易对的当前市场深度数据
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderBase3 true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getDepth [post]
func getDepth(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetDepth(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 获取k线
// @Description 获取k线
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.Kline true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getKlineRecords [post]
func getKlineRecords(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetKlineRecords(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 非个人，整个交易所的交易记录
// @Description 非个人，整个交易所的交易记录
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.OrderBase4 true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getTrades [post]
func getTrades(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetTrades(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 获取交易所名称
// @Description 获取交易所名称
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Param group body model.TokenAuth true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /trader/getExchangeName [post]
func getExchangeName(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := traderClient.GetExchangeName(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}
