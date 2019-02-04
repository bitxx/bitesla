package errs

const (
	Succ    = 0   //0 调用成功，运行结果参考API说明
	Success = 200 // 200 调用成功，运行结果参考API说明

	Errors        = 500
	InvalidParams = 400

	//token
	TokenEmptyErr = 1001
	TokenErr      = 1002
	TokenExpire   = 1003

	//请求相关
	RequestDataFmtErr = 2001

	//数据处理问题
	DataConvertErr = 3001

	// DB问题
	DBInitErr = 10001

	//账户相关
	RegisterErr = 20001
	EmailErr    = 20002
	PwdEmptyErr = 20003
	LoginErr    = 20004
	LoginSucc   = 20005

	//交易所相关
	ExchangeCurrencyPairSymbolFmtErr = 30001
	ExchangeNameErr                  = 30002
	ExchangeApiKeyAndSecret          = 30003
	ExchangePeriodErr                = 30004
	ExchangeSizeErr                  = 30005
	ExchangeSinceErr                 = 30006
	ExchangeCoinErr                  = 30007
	ExchangeAccountTypeErr           = 30008
	ExchangeOrderIDErr               = 30009
)
