package exchange

import "strings"

const (
	OPEN_BUY   = 1 + iota //开多
	OPEN_SELL             //开空
	CLOSE_BUY             //平多
	CLOSE_SELL            //平空
)

//不同类型的账户
//很多平台，都有不同的账户类型操作，这里统一标识，具体的使用，参考各平台用法
const (
	PointAccountTypeKey = 1 + iota
	SpotAccountTypeKey
)

//k线周期
const (
	KlinePeriod1min = 1 + iota
	KlinePeriod3min
	KlinePeriod5min
	KlinePeriod15min
	KlinePeriod30min
	KlinePeriod60min
	KlinePeriod2h
	KlinePeriod4h
	KlinePeriod6h
	KlinePeriod8h
	KlinePeriod12h
	KlinePeriod1day
	KlinePeriod3day
	KlinePeriod1week
	KlinePeriod1month
	KlinePeriod1year
)

var (
	ThisWeekContract = "this_week" //周合约
	NextWeekContract = "next_week" //次周合约
	QuarterContract  = "quarter"   //季度合约
)

//exchanges const
const (
	OkcoinCn   = "okcoin.cn"
	OkcoinCom  = "okcoin.com"
	OKEX       = "okex.com"
	OkexFuture = "okex.com"
	HUOBI      = "huobi.com"
	HuobiPro   = 2358885120275906
	BITSTAMP   = "bitstamp.net"
	KRAKEN     = "kraken.com"
	ZB         = "zb.com"
	BITFINEX   = "bitfinex.com"
	BINANCE    = "binance.com"
	POLONIEX   = "poloniex.com"
	COINEX     = "coinex.com"
	BITHUMB    = "bithumb.com"
	GATEIO     = "gate.io"
	BITTREX    = "bittrex.com"
	GDAX       = "gdax.com"
	WexNz      = "wex.nz"
	BIGONE     = "big.one"
	COIN58     = "58coin.com"
	FCOIN      = "fcoin.com"
	HITBTC     = "hitbtc.com"
	BITMEX     = "bitmex.com"
	CRYPTOPIA  = "cryptopia.co.nz"
)

/*const (
	OkcoinCn   = "okcoin.cn"
	OkcoinCom  = "okcoin.com"
	OKEX       = "okex.com"
	OkexFuture = "okex.com"
	HUOBI      = "huobi.com"
	HuobiPro   = "huobi.pro"
	BITSTAMP   = "bitstamp.net"
	KRAKEN     = "kraken.com"
	ZB         = "zb.com"
	BITFINEX   = "bitfinex.com"
	BINANCE    = "binance.com"
	POLONIEX   = "poloniex.com"
	COINEX     = "coinex.com"
	BITHUMB    = "bithumb.com"
	GATEIO     = "gate.io"
	BITTREX    = "bittrex.com"
	GDAX       = "gdax.com"
	WexNz      = "wex.nz"
	BIGONE     = "big.one"
	COIN58     = "58coin.com"
	FCOIN      = "fcoin.com"
	HITBTC     = "hitbtc.com"
	BITMEX     = "bitmex.com"
	CRYPTOPIA  = "cryptopia.co.nz"
)*/

// ETH_BTC --> ethbtc
type Symbols map[string]string

// huobi.com --> symbols
type ExSymbols map[string]Symbols

var exSymbols ExSymbols

func GetExSymbols(exName string) Symbols {
	ret, ok := exSymbols[exName]
	if !ok {
		return nil
	}
	return ret
}

/*func RegisterExSymbol(exName string, pair string) {
	if exSymbols == nil {
		exSymbols = make(ExSymbols)
	}

	if _, ok := exSymbols[exName]; !ok {
		exSymbols[exName] = make(Symbols)
	}

	exSymbols[exName][pair] = pair.ToSymbol("")
}*/

/*type Currency struct {
	Symbol string
	Desc   string
}*/

// A->B(A兑换为B)
//type CurrencyPair struct {
//	CurrencyA Currency
//	CurrencyB Currency
//}

var (
	UNKNOWN = "UNKNOWN"
	CNY     = "CNY"
	USD     = "USD"
	USDT    = "USDT"
	EUR     = "EUR"
	KRW     = "KRW"
	JPY     = "JPY"
	BTC     = "BTC"
	XBT     = "XBT"
	BCC     = "BCC"
	BCH     = "BCH"
	BCX     = "BCX"
	LTC     = "LTC"
	ETH     = "ETH"
	ETC     = "ETC"
	EOS     = "EOS"
	BTS     = "BTS"
	QTUM    = "QTUM"
	SC      = "SC"
	ANS     = "ANS"
	ZEC     = "ZEC"
	DCR     = "DCR"
	XRP     = "XRP"
	BTG     = "BTG"
	BCD     = "BCD"
	NEO     = "NEO"
	HSR     = "HSR"

	//currency pair

	BtcCny  = BTC + "_" + CNY
	LtcCny  = LTC + "_" + CNY
	BccCny  = BCC + "_" + CNY
	EthCny  = ETH + "_" + CNY
	EtcCny  = ETC + "_" + CNY
	EosCny  = EOS + "_" + CNY
	BtsCny  = BTS + "_" + CNY
	QtumCny = QTUM + "_" + CNY
	ScCny   = SC + "_" + CNY
	AnsCny  = ANS + "_" + CNY
	ZecCny  = ZEC + "_" + CNY

	BtcKrw = BTC + "_" + KRW
	EthKrw = ETH + "_" + KRW
	EtcKrw = ETC + "_" + KRW
	LtcKrw = LTC + "_" + KRW
	BchKrw = BCH + "_" + KRW

	BtcUsd = BTC + "_" + USD
	LtcUsd = LTC + "_" + USD
	EthUsd = ETH + "_" + USD
	EtcUsd = ETC + "_" + USD
	BchUsd = BCH + "_" + USD
	BccUsd = BCC + "_" + USD
	XrpUsd = XRP + "_" + USD
	BcdUsd = BCD + "_" + USD
	EosUsd = EOS + "_" + USD
	BtgUsd = BTG + "_" + USD

	BtcUsdt = BTC + "_" + USDT
	LtcUsdt = LTC + "_" + USDT
	BchUsdt = BCH + "_" + USDT
	BccUsdt = BCC + "_" + USDT
	EtcUsdt = ETC + "_" + USDT
	EthUsdt = ETH + "_" + USDT
	BcdUsdt = BCD + "_" + USDT
	NeoUsdt = NEO + "_" + USDT
	EosUsdt = EOS + "_" + USDT
	XrpUsdt = XRP + "_" + USDT
	HsrUsdt = HSR + "_" + USDT

	XrpEur = XRP + "_" + EUR

	BtcJpy = BTC + "_" + JPY
	LtcJpy = LTC + "_" + JPY
	EthJpy = ETH + "_" + JPY
	EtcJpy = ETC + "_" + JPY
	BchJpy = BCH + "_" + JPY

	LtcBtc = LTC + "_" + BTC
	EthBtc = ETH + "_" + BTC
	EtcBtc = ETC + "_" + BTC
	BccBtc = BCC + "_" + BTC
	BchBtc = BCH + "_" + BTC
	DcrBtc = DCR + "_" + BTC
	XrpBtc = XRP + "_" + BTC
	BtgBtc = BTG + "_" + BTC
	BcdBtc = BCD + "_" + BTC
	NeoBtc = NEO + "_" + BTC
	EosBtc = EOS + "_" + BTC
	HsrBtc = HSR + "_" + BTC

	EtcEth = ETC + "_" + ETH
	EosEth = EOS + "_" + ETH
	ZecEth = ZEC + "_" + ETH
	NeoEth = NEO + "_" + ETH
	HsrEth = HSR + "_" + ETH

	//currencyPair map
	CurrencyPair = map[string]string{
		BTC + "_" + CNY:  "btccny",
		LTC + "_" + CNY:  "ltccny",
		BCC + "_" + CNY:  "bcccny",
		ETH + "_" + CNY:  "ethcny",
		ETC + "_" + CNY:  "etccny",
		EOS + "_" + CNY:  "eoscny",
		BTS + "_" + CNY:  "btscny",
		QTUM + "_" + CNY: "qtumcny",
		SC + "_" + CNY:   "sccny",
		ANS + "_" + CNY:  "anscny",
		ZEC + "_" + CNY:  "zeccny",

		BTC + "_" + KRW: "btckrw",
		ETH + "_" + KRW: "ethkrw",
		ETC + "_" + KRW: "etckrw",
		LTC + "_" + KRW: "ltckrw",
		BCH + "_" + KRW: "bchkrw",

		BTC + "_" + USD: "btcusd",
		LTC + "_" + USD: "ltcusd",
		ETH + "_" + USD: "ethusd",
		ETC + "_" + USD: "etcusd",
		BCH + "_" + USD: "bchusd",
		BCC + "_" + USD: "bccusd",
		XRP + "_" + USD: "xrpusd",
		BCD + "_" + USD: "bcdusd",
		EOS + "_" + USD: "eosusd",
		BTG + "_" + USD: "btgusd",

		BTC + "_" + USDT: "btcusdt",
		LTC + "_" + USDT: "ltcusdt",
		BCH + "_" + USDT: "bchusdt",
		BCC + "_" + USDT: "bccusdt",
		ETC + "_" + USDT: "etcusdt",
		ETH + "_" + USDT: "ethusdt",
		BCD + "_" + USDT: "bcdusdt",
		NEO + "_" + USDT: "neousdt",
		EOS + "_" + USDT: "eosusdt",
		XRP + "_" + USDT: "xrpusdt",
		HSR + "_" + USDT: "hsrusdt",

		XRP + "_" + EUR: "xrpeur",

		BTC + "_" + JPY: "btcjpy",
		LTC + "_" + JPY: "ltcjpy",
		ETH + "_" + JPY: "ethjpy",
		ETC + "_" + JPY: "etcjpy",
		BCH + "_" + JPY: "bchjpy",

		LTC + "_" + BTC: "ltcbtc",
		ETH + "_" + BTC: "ethbtc",
		ETC + "_" + BTC: "etcbtc",
		BCC + "_" + BTC: "bccbtc",
		BCH + "_" + BTC: "bchbtc",
		DCR + "_" + BTC: "dcrbtc",
		XRP + "_" + BTC: "xrpbtc",
		BTG + "_" + BTC: "btgbtc",
		BCD + "_" + BTC: "bcdbtc",
		NEO + "_" + BTC: "neobtc",
		EOS + "_" + BTC: "eosbtc",
		HSR + "_" + BTC: "hsrbtc",

		ETC + "_" + ETH: "etceth",
		EOS + "_" + ETH: "eoseth",
		ZEC + "_" + ETH: "zeceth",
		NEO + "_" + ETH: "neoeth",
		HSR + "_" + ETH: "hsreth",
	}
)

func NewCurrency(symbol string) string {
	return strings.ToUpper(symbol)
}

func NewCurrencyPair(currencyA string, currencyB string) string {
	return currencyA + "_" + currencyB
}

func CurrencyPairSymbol(currencyA, currencyB, joinChar string) string {
	return strings.Join([]string{currencyA, currencyB}, joinChar)
}
