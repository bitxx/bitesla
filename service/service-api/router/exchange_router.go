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
	exchangeOrderPlaceUrl        = "/exchange/orderPlace"
	exchangeCancelOrderUrl       = "/exchange/cancelOrder"
	exchangeGetOneOrderUrl       = "/exchange/getOneOrder"
	exchangeGetUnfinishOrdersUrl = "/exchange/getUnfinishOrders"
	exchangeGetOrderHistorysUrl  = "/exchange/getOrderHistorys"
	exchangeGetAccountUrl        = "/exchange/getAccount"
	exchangeGetTickerUrl         = "/exchange/getTicker"
	exchangeGetDepthUrl          = "/exchange/getDepth"
	exchangeGetKlineRecordsUrl   = "/exchange/getKlineRecords"
	exchangeGetTradesUrl         = "/exchange/getTrades"
	exchangeGetExchangeDetailUrl = "/exchange/getExchangeDetail"
	exchangePutUrl               = "/exchange/put"
	exchangeListUrl              = "/exchange/list"
)

func exchangeRouter(router *gin.Engine) {
	router.POST(exchangeOrderPlaceUrl, exchangeOrderPlace)
	router.POST(exchangeCancelOrderUrl, exchangeCancelOrder)
	router.POST(exchangeGetOneOrderUrl, exchangeGetOneOrder)
	router.POST(exchangeGetUnfinishOrdersUrl, exchangeGetUnfinishOrders)
	router.POST(exchangeGetOrderHistorysUrl, exchangeGetOrderHistorys)
	router.POST(exchangeGetAccountUrl, exchangeGetAccount)
	router.POST(exchangeGetTickerUrl, exchangeGetTicker)
	router.POST(exchangeGetDepthUrl, exchangeGetDepth)
	router.POST(exchangeGetKlineRecordsUrl, exchangeGetKlineRecords)
	router.POST(exchangeGetTradesUrl, exchangeGetTrades)
	router.POST(exchangeGetExchangeDetailUrl, exchangeGetExchangeDetail)
	router.POST(exchangePutUrl, exchangePut)
	router.POST(exchangeListUrl, exchangeList)
}

var (
	exchangeClient = client.NewTraderClient()
)

// @Summary 发送一个新的订单到某交易所进行撮合
// @Description 发送一个新的订单到某交易所进行撮合
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.OrderPlace true "每个参数均不得为空,//	OrderType: 0 表示limitBuy 1 表示limitSell 2 表示marketBuy 3 表示marketSell ; AccountType: 1 表示point 2 表示splot"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/orderPlace [post]
func exchangeOrderPlace(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.OrderPlace(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.CancelOrder true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/cancelOrder [post]
func exchangeCancelOrder(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.CancelOrder(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.OneOrder true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getOneOrder [post]
func exchangeGetOneOrder(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetOneOrder(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.UnfinishOrders true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getUnfinishOrders [post]
func exchangeGetUnfinishOrders(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetUnfinishOrders(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.OrderHistory true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getOrderHistorys [post]
func exchangeGetOrderHistorys(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetOrderHistorys(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.Account true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getAccount [post]
func exchangeGetAccount(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetAccount(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.Ticker true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getTicker [post]
func exchangeGetTicker(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetTicker(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.Depth true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getDepth [post]
func exchangeGetDepth(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetDepth(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.Kline true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getKlineRecords [post]
func exchangeGetKlineRecords(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetKlineRecords(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
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
// @Security token
// @Param group body model.Trades true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getTrades [post]
func exchangeGetTrades(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetTrades(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 获取交易所详情
// @Description 获取交易所详情
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.ExchangeDetail true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/getExchangeDetail [post]
func exchangeGetExchangeDetail(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.GetExchangeDetail(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}

// @Summary 新增或者更新一个交易所
// @Description 新增或者更新一个交易所
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.ExchangeInfo true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/put [post]
func exchangePut(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.PutExchange(reqData)
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
// @Tags 交易所相关
// @Accept   json
// @Produce   json
// @Security token
// @Param group body model.ExchangeList true "每个参数均不得为空"
// @Success 200 {string} string "返回成功与否"
// @Router /exchange/list [post]
func exchangeList(c *gin.Context) {
	res := result.NewResult()
	defer c.JSON(http.StatusOK, res)
	reqData, _ := c.GetRawData()
	data, code, err := exchangeClient.ListExchange(reqData)
	res.Code = code
	res.Msg = errs.GetMsg(code)
	if err != nil {
		res.Msg = err.Error()
		logger.Error("错误信息：", err.Error())
		return
	}
	res.Data = data
}
