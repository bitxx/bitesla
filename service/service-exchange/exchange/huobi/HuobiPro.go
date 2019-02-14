package huobi

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/service/service-exchange/exchange"
	"github.com/jason-wj/bitesla/service/service-exchange/proto"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	baseWssUrl           = "wss://api.huobi.br.com/ws"
	baseHttpsUrl         = "https://api.huobi.pro"
	klineUrl             = "/market/history/kline?period=%s&size=%d&symbol=%s"
	accountsInfoUrl      = "/v1/account/accounts"
	accountBalanceUrl    = "/v1/account/accounts/%s/balance"
	placeOrders          = "/v1/order/orders/place"
	oneOrder             = "/v1/order/orders/"
	cancelOrder          = "/v1/order/orders/%s/submitcancel"
	getOrders            = "/v1/order/orders"
	getTicker            = "/market/detail/merged?symbol="
	getDepth             = "/market/depth?symbol=%s&type=step0"
	getSupportCurrencies = "/v1/common/currencys"
	getSupportSymbol     = "/v1/common/symbols"
)

var HBPOINT = exchange.NewCurrency("HBPOINT")
var onceWsConn sync.Once

var inernalKlinePeriodConverter = map[int32]string{
	exchange.KlinePeriod1min:   "1min",
	exchange.KlinePeriod5min:   "5min",
	exchange.KlinePeriod15min:  "15min",
	exchange.KlinePeriod30min:  "30min",
	exchange.KlinePeriod60min:  "60min",
	exchange.KlinePeriod1day:   "1day",
	exchange.KlinePeriod1week:  "1week",
	exchange.KlinePeriod1month: "1mon",
	exchange.KlinePeriod1year:  "1year",
}

var (
	accountType = map[int32]string{
		exchange.PointAccountTypeKey: "point",
		exchange.SpotAccountTypeKey:  "spot",
	}
)

type AccountInfo struct {
	Id    string
	Type  string
	State string
}

type HuoBiPro struct {
	httpClient        *http.Client
	baseUrl           string
	accountInfos      []*AccountInfo
	accessKey         string
	secretKey         string
	ECDSAPrivateKey   string
	ws                *exchange.WsConn
	createWsLock      sync.Mutex
	wsTickerHandleMap map[string]func(*bitesla_srv_exchange.Ticker)
	wsDepthHandleMap  map[string]func(*bitesla_srv_exchange.Depth)
	wsKLineHandleMap  map[string]func(*bitesla_srv_exchange.Kline)
}

type HuoBiProSymbol struct {
	BaseCurrency    string
	QuoteCurrency   string
	PricePrecision  float64
	AmountPrecision float64
	SymbolPartition string
	Symbol          string
}

//NewHuoBi
//needAccountInfo:若为true，则会向火币发起请求获取账户信息，false则不请求。因为对于某些隐私操作是需要用户信息的，
//但还有一些开发性但信息，是不需要获取用户信息
func NewHuoBi(client *http.Client, apikey, secretkey string, needAccountInfo bool) (*HuoBiPro, error) {
	hbpro := &HuoBiPro{}
	hbpro.baseUrl = baseHttpsUrl
	hbpro.httpClient = client
	hbpro.accessKey = apikey
	hbpro.secretKey = secretkey
	hbpro.wsDepthHandleMap = make(map[string]func(*bitesla_srv_exchange.Depth))
	hbpro.wsTickerHandleMap = make(map[string]func(*bitesla_srv_exchange.Ticker))
	hbpro.wsKLineHandleMap = make(map[string]func(*bitesla_srv_exchange.Kline))
	if needAccountInfo {
		params := &url.Values{}
		hbpro.buildPostForm("GET", accountsInfoUrl, params)
		tmpUrl := hbpro.baseUrl + accountsInfoUrl + "?" + params.Encode()
		respmap, err := exchange.HttpGet(hbpro.httpClient, tmpUrl)
		if err != nil {
			return nil, err
		}
		if respmap["status"].(string) != "ok" {
			return nil, errors.New(respmap["err-code"].(string))
		}

		accountInfo := &AccountInfo{}

		data := respmap["data"].([]interface{})
		for _, v := range data {
			iddata := v.(map[string]interface{})
			accountInfo.Id = fmt.Sprintf("%.0f", iddata["id"])
			accountInfo.Type = iddata["type"].(string)
			accountInfo.State = iddata["state"].(string)
			hbpro.accountInfos = append(hbpro.accountInfos, accountInfo)
			logger.Info("火币账户state: ", accountInfo.State, ",账号id:", accountInfo.Id, ",账号type：", accountInfo.Type)
		}
	}

	return hbpro, nil
}

func (hbpro *HuoBiPro) LimitBuy(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	orderId, err := hbpro.placeOrder(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, "buy-limit", reqCurrency.AccountType)
	if err != nil {
		return err
	}
	order.CurrencyPair = exchange.CurrencyPair[reqCurrency.CurrencyPair]
	order.OrderID = exchange.ToInt32(orderId)
	order.Amount = exchange.ToFloat64(reqCurrency.Amount)
	order.Price = exchange.ToFloat64(reqCurrency.Price)
	order.TradeSide = bitesla_srv_exchange.TradeSide_BUY
	return nil
}

func (hbpro *HuoBiPro) LimitSell(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	orderId, err := hbpro.placeOrder(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, "sell-limit", reqCurrency.AccountType)
	if err != nil {
		return err
	}
	order.CurrencyPair = exchange.CurrencyPair[reqCurrency.CurrencyPair]
	order.OrderID = exchange.ToInt32(orderId)
	order.Amount = exchange.ToFloat64(reqCurrency.Amount)
	order.Price = exchange.ToFloat64(reqCurrency.Price)
	order.TradeSide = bitesla_srv_exchange.TradeSide_SELL
	return nil
}

func (hbpro *HuoBiPro) MarketBuy(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	orderId, err := hbpro.placeOrder(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, "buy-market", reqCurrency.AccountType)
	if err != nil {
		return err
	}
	order.CurrencyPair = exchange.CurrencyPair[reqCurrency.CurrencyPair]
	order.OrderID = exchange.ToInt32(orderId)
	order.Amount = exchange.ToFloat64(reqCurrency.Amount)
	order.Price = exchange.ToFloat64(reqCurrency.Price)
	order.TradeSide = bitesla_srv_exchange.TradeSide_BUY_MARKET
	return nil
}

func (hbpro *HuoBiPro) MarketSell(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	orderId, err := hbpro.placeOrder(reqCurrency.Amount, reqCurrency.Price, reqCurrency.CurrencyPair, "sell-market", reqCurrency.AccountType)
	if err != nil {
		return err
	}

	order.CurrencyPair = exchange.CurrencyPair[reqCurrency.CurrencyPair]
	order.OrderID = exchange.ToInt32(orderId)
	order.Amount = exchange.ToFloat64(reqCurrency.Amount)
	order.Price = exchange.ToFloat64(reqCurrency.Price)
	order.TradeSide = bitesla_srv_exchange.TradeSide_SELL_MARKET
	return nil
}

func (hbpro *HuoBiPro) CancelOrder(reqCurrency *bitesla_srv_exchange.Currency, b *bitesla_srv_exchange.Boolean) error {
	path := fmt.Sprintf(cancelOrder, reqCurrency.OrderId)
	params := url.Values{}
	hbpro.buildPostForm("POST", path, &params)
	resp, err := exchange.HttpPostForm3(hbpro.httpClient, hbpro.baseUrl+path+"?"+params.Encode(), hbpro.toJson(params),
		map[string]string{"Content-Type": "application/json", "Accept-Language": "zh-cn"})
	if err != nil {
		b.IsBool = false
		return err
	}

	var respmap map[string]interface{}
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		b.IsBool = false
		return err
	}

	if respmap["status"].(string) != "ok" {
		b.IsBool = false
		return errors.New(string(resp))
	}
	b.IsBool = true
	return nil
}

func (hbpro *HuoBiPro) GetOneOrder(reqCurrency *bitesla_srv_exchange.Currency, order *bitesla_srv_exchange.Order) error {
	path := oneOrder + reqCurrency.OrderId
	params := url.Values{}
	hbpro.buildPostForm("GET", path, &params)
	respmap, err := exchange.HttpGet(hbpro.httpClient, hbpro.baseUrl+path+"?"+params.Encode())
	if err != nil {
		return err
	}

	if respmap["status"].(string) != "ok" {
		return errors.New(respmap["err-code"].(string))
	}

	datamap := respmap["data"].(map[string]interface{})
	hbpro.parseOrder(datamap, order)
	order.CurrencyPair = exchange.CurrencyPair[reqCurrency.CurrencyPair]
	return nil
}

func (hbpro *HuoBiPro) GetUnfinishOrders(reqCurrency *bitesla_srv_exchange.Currency, orders *bitesla_srv_exchange.Orders) error {
	err := hbpro.getOrders(queryOrdersParams{
		pair:   exchange.CurrencyPair[reqCurrency.CurrencyPair],
		states: "pre-submitted,submitted,partial-filled",
		size:   reqCurrency.Size,
		//direct:""
	}, orders)
	return err
}

func (hbpro *HuoBiPro) GetOrderHistorys(reqCurrency *bitesla_srv_exchange.Currency, orders *bitesla_srv_exchange.Orders) error {
	err := hbpro.getOrders(queryOrdersParams{
		pair:   exchange.CurrencyPair[reqCurrency.CurrencyPair],
		states: "partial-canceled,filled",
		size:   reqCurrency.Size,
		direct: "next",
	}, orders)
	return err
}

func (hbpro *HuoBiPro) GetTicker(reqCurrency *bitesla_srv_exchange.Currency, ticker *bitesla_srv_exchange.Ticker) error {
	url := hbpro.baseUrl + getTicker + exchange.CurrencyPair[reqCurrency.CurrencyPair]
	respmap, err := exchange.HttpGet(hbpro.httpClient, url)
	if err != nil {
		return err
	}

	if respmap["status"].(string) == "error" {
		return errors.New(respmap["err-msg"].(string))
	}

	tickmap, ok := respmap["tick"].(map[string]interface{})
	if !ok {
		return errors.New("tick assert error")
	}

	ticker.Vol = exchange.ToFloat64(tickmap["amount"])
	ticker.Low = exchange.ToFloat64(tickmap["low"])
	ticker.High = exchange.ToFloat64(tickmap["high"])
	bid, isOk := tickmap["bid"].([]interface{})
	if isOk != true {
		return errors.New("no bid")
	}
	ask, isOk := tickmap["ask"].([]interface{})
	if isOk != true {
		return errors.New("no ask")
	}
	ticker.Buy = exchange.ToFloat64(bid[0])
	ticker.Sell = exchange.ToFloat64(ask[0])
	ticker.Last = exchange.ToFloat64(tickmap["close"])
	ticker.Date = exchange.ToUint64(respmap["ts"])

	return nil
}

func (hbpro *HuoBiPro) GetKlineRecords(reqCurrency *bitesla_srv_exchange.Currency, kLines *bitesla_srv_exchange.Klines) error {
	symbol := exchange.CurrencyPair[reqCurrency.CurrencyPair]
	periodS, isOk := inernalKlinePeriodConverter[reqCurrency.Period]
	if isOk != true {
		periodS = "1min"
	}

	logger.Info("火币交易所请求k线URL：", fmt.Sprintf(hbpro.baseUrl+klineUrl, periodS, reqCurrency.Size, symbol))

	ret, err := exchange.HttpGet(hbpro.httpClient, fmt.Sprintf(hbpro.baseUrl+klineUrl, periodS, reqCurrency.Size, symbol))
	if err != nil {
		return err
	}

	data, ok := ret["data"].([]interface{})
	if !ok {
		return errors.New("response format error")
	}

	for _, e := range data {
		item := e.(map[string]interface{})
		kLines.Klines = append(kLines.Klines, &bitesla_srv_exchange.Kline{
			CurrencyPair: symbol,
			Open:         exchange.ToFloat64(item["open"]),
			Close:        exchange.ToFloat64(item["close"]),
			High:         exchange.ToFloat64(item["high"]),
			Low:          exchange.ToFloat64(item["low"]),
			Vol:          exchange.ToFloat64(item["vol"]),
			Timestamp:    int64(exchange.ToUint64(item["id"]))})
	}

	return nil
}

func (hbpro *HuoBiPro) GetTrades(*bitesla_srv_exchange.Currency, *bitesla_srv_exchange.Trades) error {
	return nil
}

func (hbpro *HuoBiPro) GetAccount(reqCurrency *bitesla_srv_exchange.Currency, accounts *bitesla_srv_exchange.Accounts) error {
	for _, accountInfo := range hbpro.accountInfos {

		path := fmt.Sprintf(accountBalanceUrl, accountInfo.Id)
		params := &url.Values{}
		params.Set("account-id", accountInfo.Id)
		hbpro.buildPostForm("GET", path, params)

		urlStr := hbpro.baseUrl + path + "?" + params.Encode()
		//println(urlStr)
		logger.Info("火币交易所请求账户URL：", urlStr)
		respmap, err := exchange.HttpGet(hbpro.httpClient, urlStr)

		if err != nil {
			return err
		}

		if respmap["status"].(string) != "ok" {
			return errors.New(respmap["err-code"].(string))
		}

		datamap := respmap["data"].(map[string]interface{})
		if datamap["state"].(string) != "working" {
			return errors.New(datamap["state"].(string))
		}

		list := datamap["list"].([]interface{})

		acc := new(bitesla_srv_exchange.Account)
		acc.SubAccounts = make(map[string]*bitesla_srv_exchange.SubAccount, 6)

		subAccMap := make(map[string]*bitesla_srv_exchange.SubAccount)

		for _, v := range list {
			balancemap := v.(map[string]interface{})
			currencySymbol := balancemap["currency"].(string)
			currency := exchange.NewCurrency(currencySymbol)
			typeStr := balancemap["type"].(string)
			balance := exchange.ToFloat64(balancemap["balance"])
			if subAccMap[currency] == nil {
				subAccMap[currency] = new(bitesla_srv_exchange.SubAccount)
			}
			subAccMap[currency].Currency = currency
			switch typeStr {
			case "trade":
				subAccMap[currency].Amount = balance
			case "frozen":
				subAccMap[currency].ForzenAmount = balance
			}
		}

		for k, v := range subAccMap {
			acc.SubAccounts[k] = v
		}
		accounts.Accounts = append(accounts.Accounts, acc)
	}
	return nil
}

//发送一个新订单到火币以进行撮合
func (hbpro *HuoBiPro) placeOrder(amount, price, pair, orderType string, accType int32) (string, error) {
	params := url.Values{}
	accountID := ""
	for _, accountInfo := range hbpro.accountInfos {
		if accountInfo.Type == accountType[accType] {
			accountID = accountInfo.Id
		}
	}
	if accountID == "" {
		return "", errors.New("该类型账号为空")
	}
	params.Set("account-id", accountID)
	params.Set("amount", amount)
	params.Set("symbol", exchange.CurrencyPair[pair])
	params.Set("type", orderType)

	switch orderType {
	case "buy-limit", "sell-limit":
		params.Set("price", price)
	}

	hbpro.buildPostForm("POST", placeOrders, &params)

	resp, err := exchange.HttpPostForm3(hbpro.httpClient, hbpro.baseUrl+placeOrders+"?"+params.Encode(), hbpro.toJson(params),
		map[string]string{"Content-Type": "application/json", "Accept-Language": "zh-cn"})
	if err != nil {
		return "", err
	}

	respmap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		return "", err
	}

	if respmap["status"].(string) != "ok" {
		return "", errors.New(respmap["err-code"].(string))
	}

	return respmap["data"].(string), nil
}

func (hbpro *HuoBiPro) parseOrder(ordmap map[string]interface{}, order *bitesla_srv_exchange.Order) {

	order.OrderID = exchange.ToInt32(ordmap["id"])
	order.Amount = exchange.ToFloat64(ordmap["amount"])
	order.Price = exchange.ToFloat64(ordmap["price"])
	order.DealAmount = exchange.ToFloat64(ordmap["field-amount"])
	order.Fee = exchange.ToFloat64(ordmap["field-fees"])
	order.OrderTime = exchange.ToInt32(ordmap["created-at"])

	state := ordmap["state"].(string)
	switch state {
	case "submitted", "pre-submitted":
		order.Status = bitesla_srv_exchange.TradeStatus_UNFINISH
	case "filled":
		order.Status = bitesla_srv_exchange.TradeStatus_FINISH
	case "partial-filled":
		order.Status = bitesla_srv_exchange.TradeStatus_PART_FINISH
	case "canceled", "partial-canceled":
		order.Status = bitesla_srv_exchange.TradeStatus_CANCEL
	default:
		order.Status = bitesla_srv_exchange.TradeStatus_UNFINISH
	}

	if order.DealAmount > 0.0 {
		order.AvgPrice = exchange.ToFloat64(ordmap["field-cash-amount"]) / order.DealAmount
	}

	typeS := ordmap["type"].(string)
	switch typeS {
	case "buy-limit":
		order.TradeSide = bitesla_srv_exchange.TradeSide_BUY
	case "buy-market":
		order.TradeSide = bitesla_srv_exchange.TradeSide_BUY_MARKET
	case "sell-limit":
		order.TradeSide = bitesla_srv_exchange.TradeSide_SELL
	case "sell-market":
		order.TradeSide = bitesla_srv_exchange.TradeSide_SELL_MARKET
	}
}

type queryOrdersParams struct {
	types,
	startDate,
	endDate,
	states,
	from,
	direct string
	size int32
	pair string
}

func (hbpro *HuoBiPro) getOrders(queryparams queryOrdersParams, orders *bitesla_srv_exchange.Orders) error {
	path := getOrders
	params := url.Values{}
	params.Set("symbol", exchange.CurrencyPair[queryparams.pair])
	params.Set("states", queryparams.states)

	if queryparams.direct != "" {
		params.Set("direct", queryparams.direct)
	}

	if queryparams.size > 0 {
		params.Set("size", fmt.Sprint(queryparams.size))
	}

	hbpro.buildPostForm("GET", path, &params)
	respmap, err := exchange.HttpGet(hbpro.httpClient, fmt.Sprintf("%s%s?%s", hbpro.baseUrl, path, params.Encode()))
	if err != nil {
		return err
	}

	if respmap["status"].(string) != "ok" {
		return errors.New(respmap["err-code"].(string))
	}

	datamap := respmap["data"].([]interface{})
	for _, v := range datamap {
		order := &bitesla_srv_exchange.Order{}
		ordmap := v.(map[string]interface{})
		hbpro.parseOrder(ordmap, order)
		order.CurrencyPair = queryparams.pair
		orders.Orders = append(orders.Orders, order)
	}

	return nil
}

//返回深度数据
func (hbpro *HuoBiPro) GetDepth(reqCurrency *bitesla_srv_exchange.Currency, depth *bitesla_srv_exchange.Depth) error {
	url := hbpro.baseUrl + getDepth
	respmap, err := exchange.HttpGet(hbpro.httpClient, fmt.Sprintf(url, exchange.CurrencyPair[reqCurrency.CurrencyPair]))
	if err != nil {
		return err
	}

	if "ok" != respmap["status"].(string) {
		return errors.New(respmap["err-msg"].(string))
	}

	tick, _ := respmap["tick"].(map[string]interface{})
	hbpro.parseDepthData(tick, depth)
	return nil
}

type ecdsaSignature struct {
	R, S *big.Int
}

func (hbpro *HuoBiPro) buildPostForm(reqMethod, path string, postForm *url.Values) {
	postForm.Set("AccessKeyId", hbpro.accessKey)
	postForm.Set("SignatureMethod", "HmacSHA256")
	postForm.Set("SignatureVersion", "2")
	postForm.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))
	domain := strings.Replace(hbpro.baseUrl, "https://", "", len(hbpro.baseUrl))
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", reqMethod, domain, path, postForm.Encode())
	sign, _ := exchange.GetParamHmacSHA256Base64Sign(hbpro.secretKey, payload)
	postForm.Set("Signature", sign)

	/**
	p, _ := pem.Decode([]byte(hbpro.ECDSAPrivateKey))
	pri, _ := secp256k1_go.PrivKeyFromBytes(secp256k1_go.S256(), p.Bytes)
	signer, _ := pri.Sign([]byte(sign))
	signAsn, _ := asn1.Marshal(signer)
	priSign := base64.StdEncoding.EncodeToString(signAsn)
	postForm.Set("PrivateSignature", priSign)
	*/
}

func (hbpro *HuoBiPro) toJson(params url.Values) string {
	parammap := make(map[string]string)
	for k, v := range params {
		parammap[k] = v[0]
	}
	jsonData, _ := json.Marshal(parammap)
	return string(jsonData)
}

func (hbpro *HuoBiPro) createWsConn() {

	onceWsConn.Do(func() {
		hbpro.ws = exchange.NewWsConn(baseWssUrl)
		hbpro.ws.Heartbeat(func() interface{} {
			return map[string]interface{}{
				"ping": time.Now().Unix()}
		}, 5*time.Second)
		hbpro.ws.ReConnect()
		hbpro.ws.ReceiveMessage(func(msg []byte) {
			gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
			data, _ := ioutil.ReadAll(gzipreader)
			datamap := make(map[string]interface{})
			err := json.Unmarshal(data, &datamap)
			if err != nil {
				log.Println("json unmarshal error for ", string(data))
				return
			}

			if datamap["ping"] != nil {
				//log.Println(datamap)
				hbpro.ws.UpdateActivedTime()
				hbpro.ws.SendWriteJSON(map[string]interface{}{
					"pong": datamap["ping"]}) // 回应心跳
				return
			}

			if datamap["pong"] != nil { //
				hbpro.ws.UpdateActivedTime()
				return
			}

			if datamap["id"] != nil { //忽略订阅成功的回执消息
				log.Println(string(data))
				return
			}

			ch, isok := datamap["ch"].(string)
			if !isok {
				log.Println("error:", string(data))
				return
			}

			tick := datamap["tick"].(map[string]interface{})
			pair := hbpro.getPairFromChannel(ch)
			if hbpro.wsTickerHandleMap[ch] != nil {
				tick := hbpro.parseTickerData(tick)
				tick.Pair = pair
				tick.Date = exchange.ToUint64(datamap["ts"])
				(hbpro.wsTickerHandleMap[ch])(tick)
				return
			}

			if hbpro.wsDepthHandleMap[ch] != nil {
				depth := &bitesla_srv_exchange.Depth{}
				hbpro.parseDepthData(tick, depth)
				depth.CurrencyPair = pair
				(hbpro.wsDepthHandleMap[ch])(depth)
				return
			}

			if hbpro.wsKLineHandleMap[ch] != nil {
				kline := hbpro.parseWsKLineData(tick)
				kline.CurrencyPair = pair
				(hbpro.wsKLineHandleMap[ch])(kline)
				return
			}

			//log.Println(string(data))
		})
	})

}

func (hbpro *HuoBiPro) getPairFromChannel(ch string) string {
	s := strings.Split(ch, ".")
	var currA, currB string
	if strings.HasSuffix(s[1], "usdt") {
		currB = "usdt"
	} else if strings.HasSuffix(s[1], "husd") {
		currB = "husd"
	} else if strings.HasSuffix(s[1], "btc") {
		currB = "btc"
	} else if strings.HasSuffix(s[1], "eth") {
		currB = "eth"
	} else if strings.HasSuffix(s[1], "ht") {
		currB = "ht"
	}

	currA = strings.TrimSuffix(s[1], currB)

	a := exchange.NewCurrency(currA)
	b := exchange.NewCurrency(currB)
	pair := exchange.NewCurrencyPair(a, b)
	return pair
}

func (hbpro *HuoBiPro) parseTickerData(tick map[string]interface{}) *bitesla_srv_exchange.Ticker {
	t := new(bitesla_srv_exchange.Ticker)

	t.Last = exchange.ToFloat64(tick["close"])
	t.Low = exchange.ToFloat64(tick["low"])
	t.Vol = exchange.ToFloat64(tick["vol"])
	t.High = exchange.ToFloat64(tick["high"])
	return t
}

func (hbpro *HuoBiPro) parseDepthData(tick map[string]interface{}, depth *bitesla_srv_exchange.Depth) {
	bids, _ := tick["bids"].([]interface{})
	asks, _ := tick["asks"].([]interface{})

	//depth := &bitesla_srv_exchange.Depth{}
	depth.AskList = &bitesla_srv_exchange.DepthRecords{}
	depth.AskList.List = []*bitesla_srv_exchange.DepthRecord{}
	depth.BidList = &bitesla_srv_exchange.DepthRecords{}
	depth.BidList.List = []*bitesla_srv_exchange.DepthRecord{}
	for _, r := range asks {
		var dr bitesla_srv_exchange.DepthRecord
		rr := r.([]interface{})
		dr.Price = exchange.ToFloat64(rr[0])
		dr.Amount = exchange.ToFloat64(rr[1])
		depth.AskList.List = append(depth.AskList.List, &dr)
	}

	for _, r := range bids {
		var dr bitesla_srv_exchange.DepthRecord
		rr := r.([]interface{})
		dr.Price = exchange.ToFloat64(rr[0])
		dr.Amount = exchange.ToFloat64(rr[1])
		depth.BidList.List = append(depth.BidList.List, &dr)
	}

	sort.Sort(sort.Reverse(depth.AskList))
}

func (hbpro *HuoBiPro) parseWsKLineData(tick map[string]interface{}) *bitesla_srv_exchange.Kline {
	return &bitesla_srv_exchange.Kline{
		Open:      exchange.ToFloat64(tick["open"]),
		Close:     exchange.ToFloat64(tick["close"]),
		High:      exchange.ToFloat64(tick["high"]),
		Low:       exchange.ToFloat64(tick["low"]),
		Vol:       exchange.ToFloat64(tick["vol"]),
		Timestamp: int64(exchange.ToUint64(tick["id"]))}
}

//返回当前交易所支持的代币
func (hbpro *HuoBiPro) GetCurrenciesList() ([]string, error) {
	url := hbpro.baseUrl + getSupportCurrencies

	ret, err := exchange.HttpGet(hbpro.httpClient, url)
	if err != nil {
		return nil, err
	}

	data, ok := ret["data"].([]interface{})
	if !ok {
		return nil, errors.New("response format error")
	}
	fmt.Println(data)
	return nil, nil
}

func (hbpro *HuoBiPro) GetCurrenciesPrecision() ([]HuoBiProSymbol, error) {
	url := hbpro.baseUrl + getSupportSymbol

	ret, err := exchange.HttpGet(hbpro.httpClient, url)
	if err != nil {
		return nil, err
	}

	data, ok := ret["data"].([]interface{})
	if !ok {
		return nil, errors.New("response format error")
	}
	var Symbols []HuoBiProSymbol
	for _, v := range data {
		_sym := v.(map[string]interface{})
		var sym HuoBiProSymbol
		sym.BaseCurrency = _sym["base-currency"].(string)
		sym.QuoteCurrency = _sym["quote-currency"].(string)
		sym.PricePrecision = _sym["price-precision"].(float64)
		sym.AmountPrecision = _sym["amount-precision"].(float64)
		sym.SymbolPartition = _sym["symbol-partition"].(string)
		sym.Symbol = _sym["symbol"].(string)
		Symbols = append(Symbols, sym)
	}
	return Symbols, nil
}

func (hbpro *HuoBiPro) GetTickerWithWs(pair string, handle func(ticker *bitesla_srv_exchange.Ticker)) error {
	hbpro.createWsConn()
	sub := fmt.Sprintf("market.%s.detail", exchange.CurrencyPair[pair])
	hbpro.wsTickerHandleMap[sub] = handle
	return hbpro.ws.Subscribe(map[string]interface{}{
		"id":  1,
		"sub": sub})
}

func (hbpro *HuoBiPro) GetDepthWithWs(pair string, handle func(dep *bitesla_srv_exchange.Depth)) error {
	hbpro.createWsConn()
	sub := fmt.Sprintf("market.%s.depth.step0", exchange.CurrencyPair[pair])
	hbpro.wsDepthHandleMap[sub] = handle
	return hbpro.ws.Subscribe(map[string]interface{}{
		"id":  2,
		"sub": sub})
}

func (hbpro *HuoBiPro) GetKLineWithWs(pair string, period int32, handle func(kline *bitesla_srv_exchange.Kline)) error {
	hbpro.createWsConn()
	periodS, isOk := inernalKlinePeriodConverter[period]
	if isOk != true {
		periodS = "1min"
	}

	sub := fmt.Sprintf("market.%s.kline.%s", exchange.CurrencyPair[pair], periodS)
	hbpro.wsKLineHandleMap[sub] = handle
	return hbpro.ws.Subscribe(map[string]interface{}{
		"id":  3,
		"sub": sub})
}
