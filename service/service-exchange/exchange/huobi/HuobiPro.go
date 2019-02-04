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
	wsTickerHandleMap map[string]func(*bitesla_srv_trader.Ticker)
	wsDepthHandleMap  map[string]func(*bitesla_srv_trader.Depth)
	wsKLineHandleMap  map[string]func(*bitesla_srv_trader.Kline)
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
	hbpro.wsDepthHandleMap = make(map[string]func(*bitesla_srv_trader.Depth))
	hbpro.wsTickerHandleMap = make(map[string]func(*bitesla_srv_trader.Ticker))
	hbpro.wsKLineHandleMap = make(map[string]func(*bitesla_srv_trader.Kline))
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

func (hbpro *HuoBiPro) GetAccount() ([]*bitesla_srv_trader.Account, error) {
	var accounts []*bitesla_srv_trader.Account
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
			return nil, err
		}

		if respmap["status"].(string) != "ok" {
			return nil, errors.New(respmap["err-code"].(string))
		}

		datamap := respmap["data"].(map[string]interface{})
		if datamap["state"].(string) != "working" {
			return nil, errors.New(datamap["state"].(string))
		}

		list := datamap["list"].([]interface{})
		acc := new(bitesla_srv_trader.Account)
		acc.SubAccounts = make(map[string]*bitesla_srv_trader.SubAccount, 6)
		acc.Exchange = hbpro.GetExchangeName()

		subAccMap := make(map[string]*bitesla_srv_trader.SubAccount)

		for _, v := range list {
			balancemap := v.(map[string]interface{})
			currencySymbol := balancemap["currency"].(string)
			currency := exchange.NewCurrency(currencySymbol)
			typeStr := balancemap["type"].(string)
			balance := exchange.ToFloat64(balancemap["balance"])
			if subAccMap[currency] == nil {
				subAccMap[currency] = new(bitesla_srv_trader.SubAccount)
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
		accounts = append(accounts, acc)
	}
	return accounts, nil
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

func (hbpro *HuoBiPro) LimitBuy(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error) {
	orderId, err := hbpro.placeOrder(amount, price, currency, "buy-limit", accountType)
	if err != nil {
		return nil, err
	}
	return &bitesla_srv_trader.Order{
		CurrencyPair: exchange.CurrencyPair[currency],
		OrderID:      exchange.ToInt32(orderId),
		Amount:       exchange.ToFloat64(amount),
		Price:        exchange.ToFloat64(price),
		TradeSide:    bitesla_srv_trader.TradeSide_BUY}, nil
}

func (hbpro *HuoBiPro) LimitSell(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error) {
	orderId, err := hbpro.placeOrder(amount, price, currency, "sell-limit", accountType)
	if err != nil {
		return nil, err
	}
	return &bitesla_srv_trader.Order{
		CurrencyPair: exchange.CurrencyPair[currency],
		OrderID:      exchange.ToInt32(orderId),
		Amount:       exchange.ToFloat64(amount),
		Price:        exchange.ToFloat64(price),
		TradeSide:    bitesla_srv_trader.TradeSide_SELL}, nil
}

func (hbpro *HuoBiPro) MarketBuy(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error) {
	orderId, err := hbpro.placeOrder(amount, price, currency, "buy-market", accountType)
	if err != nil {
		return nil, err
	}
	return &bitesla_srv_trader.Order{
		CurrencyPair: exchange.CurrencyPair[currency],
		OrderID:      exchange.ToInt32(orderId),
		Amount:       exchange.ToFloat64(amount),
		Price:        exchange.ToFloat64(price),
		TradeSide:    bitesla_srv_trader.TradeSide_BUY_MARKET}, nil
}

func (hbpro *HuoBiPro) MarketSell(amount, price, currency string, accountType int32) (*bitesla_srv_trader.Order, error) {
	orderId, err := hbpro.placeOrder(amount, price, currency, "sell-market", accountType)
	if err != nil {
		return nil, err
	}
	return &bitesla_srv_trader.Order{
		CurrencyPair: exchange.CurrencyPair[currency],
		OrderID:      exchange.ToInt32(orderId),
		Amount:       exchange.ToFloat64(amount),
		Price:        exchange.ToFloat64(price),
		TradeSide:    bitesla_srv_trader.TradeSide_SELL_MARKET}, nil
}

func (hbpro *HuoBiPro) parseOrder(ordmap map[string]interface{}) *bitesla_srv_trader.Order {
	ord := &bitesla_srv_trader.Order{
		OrderID:    exchange.ToInt32(ordmap["id"]),
		Amount:     exchange.ToFloat64(ordmap["amount"]),
		Price:      exchange.ToFloat64(ordmap["price"]),
		DealAmount: exchange.ToFloat64(ordmap["field-amount"]),
		Fee:        exchange.ToFloat64(ordmap["field-fees"]),
		OrderTime:  exchange.ToInt32(ordmap["created-at"]),
	}

	state := ordmap["state"].(string)
	switch state {
	case "submitted", "pre-submitted":
		ord.Status = bitesla_srv_trader.TradeStatus_UNFINISH
	case "filled":
		ord.Status = bitesla_srv_trader.TradeStatus_FINISH
	case "partial-filled":
		ord.Status = bitesla_srv_trader.TradeStatus_PART_FINISH
	case "canceled", "partial-canceled":
		ord.Status = bitesla_srv_trader.TradeStatus_CANCEL
	default:
		ord.Status = bitesla_srv_trader.TradeStatus_UNFINISH
	}

	if ord.DealAmount > 0.0 {
		ord.AvgPrice = exchange.ToFloat64(ordmap["field-cash-amount"]) / ord.DealAmount
	}

	typeS := ordmap["type"].(string)
	switch typeS {
	case "buy-limit":
		ord.TradeSide = bitesla_srv_trader.TradeSide_BUY
	case "buy-market":
		ord.TradeSide = bitesla_srv_trader.TradeSide_BUY_MARKET
	case "sell-limit":
		ord.TradeSide = bitesla_srv_trader.TradeSide_SELL
	case "sell-market":
		ord.TradeSide = bitesla_srv_trader.TradeSide_SELL_MARKET
	}
	return ord
}

func (hbpro *HuoBiPro) GetOneOrder(orderId string, currency string) (*bitesla_srv_trader.Order, error) {
	path := oneOrder + orderId
	params := url.Values{}
	hbpro.buildPostForm("GET", path, &params)
	respmap, err := exchange.HttpGet(hbpro.httpClient, hbpro.baseUrl+path+"?"+params.Encode())
	if err != nil {
		return nil, err
	}

	if respmap["status"].(string) != "ok" {
		return nil, errors.New(respmap["err-code"].(string))
	}

	datamap := respmap["data"].(map[string]interface{})
	order := hbpro.parseOrder(datamap)
	order.CurrencyPair = exchange.CurrencyPair[currency]
	return order, nil
}

func (hbpro *HuoBiPro) GetUnfinishOrders(currency string) ([]*bitesla_srv_trader.Order, error) {
	return hbpro.getOrders(queryOrdersParams{
		pair:   exchange.CurrencyPair[currency],
		states: "pre-submitted,submitted,partial-filled",
		size:   100,
		//direct:""
	})
}

func (hbpro *HuoBiPro) CancelOrder(orderId string, currency string) (bool, error) {
	path := fmt.Sprintf(cancelOrder, orderId)
	params := url.Values{}
	hbpro.buildPostForm("POST", path, &params)
	resp, err := exchange.HttpPostForm3(hbpro.httpClient, hbpro.baseUrl+path+"?"+params.Encode(), hbpro.toJson(params),
		map[string]string{"Content-Type": "application/json", "Accept-Language": "zh-cn"})
	if err != nil {
		return false, err
	}

	var respmap map[string]interface{}
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		return false, err
	}

	if respmap["status"].(string) != "ok" {
		return false, errors.New(string(resp))
	}

	return true, nil
}

func (hbpro *HuoBiPro) GetOrderHistorys(currency string, currentPage, pageSize int32) ([]*bitesla_srv_trader.Order, error) {
	return hbpro.getOrders(queryOrdersParams{
		pair:   currency,
		size:   pageSize,
		states: "partial-canceled,filled",
		direct: "next",
	})
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

func (hbpro *HuoBiPro) getOrders(queryparams queryOrdersParams) ([]*bitesla_srv_trader.Order, error) {
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
		return nil, err
	}

	if respmap["status"].(string) != "ok" {
		return nil, errors.New(respmap["err-code"].(string))
	}

	datamap := respmap["data"].([]interface{})
	var orders []*bitesla_srv_trader.Order
	for _, v := range datamap {
		ordmap := v.(map[string]interface{})
		ord := hbpro.parseOrder(ordmap)
		ord.CurrencyPair = queryparams.pair
		orders = append(orders, ord)
	}

	return orders, nil
}

//提供24小时交易聚合信息
func (hbpro *HuoBiPro) GetTicker(currencyPair string) (*bitesla_srv_trader.Ticker, error) {
	url := hbpro.baseUrl + getTicker + exchange.CurrencyPair[currencyPair]
	respmap, err := exchange.HttpGet(hbpro.httpClient, url)
	if err != nil {
		return nil, err
	}

	if respmap["status"].(string) == "error" {
		return nil, errors.New(respmap["err-msg"].(string))
	}

	tickmap, ok := respmap["tick"].(map[string]interface{})
	if !ok {
		return nil, errors.New("tick assert error")
	}

	ticker := new(bitesla_srv_trader.Ticker)
	ticker.Vol = exchange.ToFloat64(tickmap["amount"])
	ticker.Low = exchange.ToFloat64(tickmap["low"])
	ticker.High = exchange.ToFloat64(tickmap["high"])
	bid, isOk := tickmap["bid"].([]interface{})
	if isOk != true {
		return nil, errors.New("no bid")
	}
	ask, isOk := tickmap["ask"].([]interface{})
	if isOk != true {
		return nil, errors.New("no ask")
	}
	ticker.Buy = exchange.ToFloat64(bid[0])
	ticker.Sell = exchange.ToFloat64(ask[0])
	ticker.Last = exchange.ToFloat64(tickmap["close"])
	ticker.Date = exchange.ToUint64(respmap["ts"])

	return ticker, nil
}

//返回深度数据
func (hbpro *HuoBiPro) GetDepth(size int32, currency string) (*bitesla_srv_trader.Depth, error) {
	url := hbpro.baseUrl + getDepth
	respmap, err := exchange.HttpGet(hbpro.httpClient, fmt.Sprintf(url, exchange.CurrencyPair[currency]))
	if err != nil {
		return nil, err
	}

	if "ok" != respmap["status"].(string) {
		return nil, errors.New(respmap["err-msg"].(string))
	}

	tick, _ := respmap["tick"].(map[string]interface{})

	return hbpro.parseDepthData(tick), nil
}

//倒序
func (hbpro *HuoBiPro) GetKlineRecords(currencyPair string, period, size, since int32) ([]*bitesla_srv_trader.Kline, error) {
	symbol := exchange.CurrencyPair[currencyPair]
	periodS, isOk := inernalKlinePeriodConverter[period]
	if isOk != true {
		periodS = "1min"
	}

	logger.Info("火币交易所请求k线URL：", fmt.Sprintf(hbpro.baseUrl+klineUrl, periodS, size, symbol))

	ret, err := exchange.HttpGet(hbpro.httpClient, fmt.Sprintf(hbpro.baseUrl+klineUrl, periodS, size, symbol))
	if err != nil {
		return nil, err
	}

	data, ok := ret["data"].([]interface{})
	if !ok {
		return nil, errors.New("response format error")
	}

	var klines []*bitesla_srv_trader.Kline
	for _, e := range data {
		item := e.(map[string]interface{})
		klines = append(klines, &bitesla_srv_trader.Kline{
			CurrencyPair: currencyPair,
			Open:         exchange.ToFloat64(item["open"]),
			Close:        exchange.ToFloat64(item["close"]),
			High:         exchange.ToFloat64(item["high"]),
			Low:          exchange.ToFloat64(item["low"]),
			Vol:          exchange.ToFloat64(item["vol"]),
			Timestamp:    int64(exchange.ToUint64(item["id"]))})
	}

	return klines, nil
}

//非个人，整个交易所的交易记录
func (hbpro *HuoBiPro) GetTrades(currencyPair string, since int32) ([]*bitesla_srv_trader.Trade, error) {
	//panic("not implement")
	return nil, nil
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
				depth := hbpro.parseDepthData(tick)
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

func (hbpro *HuoBiPro) parseTickerData(tick map[string]interface{}) *bitesla_srv_trader.Ticker {
	t := new(bitesla_srv_trader.Ticker)

	t.Last = exchange.ToFloat64(tick["close"])
	t.Low = exchange.ToFloat64(tick["low"])
	t.Vol = exchange.ToFloat64(tick["vol"])
	t.High = exchange.ToFloat64(tick["high"])
	return t
}

func (hbpro *HuoBiPro) parseDepthData(tick map[string]interface{}) *bitesla_srv_trader.Depth {
	bids, _ := tick["bids"].([]interface{})
	asks, _ := tick["asks"].([]interface{})

	depth := new(bitesla_srv_trader.Depth)
	for _, r := range asks {
		var dr bitesla_srv_trader.DepthRecord
		rr := r.([]interface{})
		dr.Price = exchange.ToFloat64(rr[0])
		dr.Amount = exchange.ToFloat64(rr[1])
		depth.AskList.List = append(depth.AskList.List, &dr)
	}

	for _, r := range bids {
		var dr bitesla_srv_trader.DepthRecord
		rr := r.([]interface{})
		dr.Price = exchange.ToFloat64(rr[0])
		dr.Amount = exchange.ToFloat64(rr[1])
		depth.BidList.List = append(depth.BidList.List, &dr)
	}

	sort.Sort(sort.Reverse(depth.AskList))

	return depth
}

func (hbpro *HuoBiPro) parseWsKLineData(tick map[string]interface{}) *bitesla_srv_trader.Kline {
	return &bitesla_srv_trader.Kline{
		Open:      exchange.ToFloat64(tick["open"]),
		Close:     exchange.ToFloat64(tick["close"]),
		High:      exchange.ToFloat64(tick["high"]),
		Low:       exchange.ToFloat64(tick["low"]),
		Vol:       exchange.ToFloat64(tick["vol"]),
		Timestamp: int64(exchange.ToUint64(tick["id"]))}
}

//获取交易所名称
func (hbpro *HuoBiPro) GetExchangeName() string {
	return exchange.HuobiPro
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

func (hbpro *HuoBiPro) GetTickerWithWs(pair string, handle func(ticker *bitesla_srv_trader.Ticker)) error {
	hbpro.createWsConn()
	sub := fmt.Sprintf("market.%s.detail", exchange.CurrencyPair[pair])
	hbpro.wsTickerHandleMap[sub] = handle
	return hbpro.ws.Subscribe(map[string]interface{}{
		"id":  1,
		"sub": sub})
}

func (hbpro *HuoBiPro) GetDepthWithWs(pair string, handle func(dep *bitesla_srv_trader.Depth)) error {
	hbpro.createWsConn()
	sub := fmt.Sprintf("market.%s.depth.step0", exchange.CurrencyPair[pair])
	hbpro.wsDepthHandleMap[sub] = handle
	return hbpro.ws.Subscribe(map[string]interface{}{
		"id":  2,
		"sub": sub})
}

func (hbpro *HuoBiPro) GetKLineWithWs(pair string, period int32, handle func(kline *bitesla_srv_trader.Kline)) error {
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
