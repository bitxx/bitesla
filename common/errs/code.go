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
	RequestDataFmtErr        = 2001
	RequestHeadCurrUserIdErr = 2002

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
	GetUserErr  = 20006
	UserIdErr   = 20007

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
	ExchangeDescriptionErr           = 30010
	ExchangeIDErr                    = 30011

	//策略管理相关
	StrategyNameErr   = 40001
	StrategyDescErr   = 40002
	StrategyScriptErr = 40003
	StrategyIdErr     = 40004

	//策略执行相关
	TraderNameErr       = 50001
	TraderDescErr       = 50002
	TraderIdErr         = 50003
	TraderStrategyIdErr = 50004
	TraderExchangeIdErr = 50005
)
