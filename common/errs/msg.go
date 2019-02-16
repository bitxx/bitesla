package errs

var MsgFlags = map[int]string{
	Succ:          "ok",
	Success:       "ok",
	Errors:        "服务器错误",
	InvalidParams: "请求参数错误",

	//token
	TokenEmptyErr: "token头部验证不得为空",
	TokenErr:      "token头部验证错误",
	TokenExpire:   "Token过期",

	//请求数据相关
	RequestDataFmtErr:        "请求数据格式异常",
	RequestHeadCurrUserIdErr: "请求数据中不得拥有currentLoginUserID字段",

	//数据转换问题
	DataConvertErr: "数据转换错误",

	// DB问题
	DBInitErr: "数据库初始化异常",

	//账户相关
	RegisterErr: "注册失败",
	EmailErr:    "邮箱格式异常",
	PwdEmptyErr: "密码不得为空",
	LoginErr:    "登录失败",
	LoginSucc:   "登录成功",
	GetUserErr:  "用户信息获取失败",
	UserIdErr:   "用户id错误",

	//交易所相关
	ExchangeCurrencyPairSymbolFmtErr: "货币格式错误",
	ExchangeNameErr:                  "交易所名称错误",
	ExchangeApiKeyAndSecret:          "appkey或appsecret输入异常",
	ExchangePeriodErr:                "Period不得小于0",
	ExchangeSizeErr:                  "Size必须大于0",
	ExchangeSinceErr:                 "Since必须大于0",
	ExchangeCoinErr:                  "币种异常，两币种需要大写，且用下划线'_'分割",
	ExchangeAccountTypeErr:           "账户类型异常",
	ExchangeOrderIDErr:               "订单号错误",
	ExchangeDescriptionErr:           "交易所描述不得为空",
	ExchangeIDErr:                    "exchangeId必须大于0 ",

	//策略管理相关
	StrategyNameErr:     "策略名称错误",
	StrategyDescErr:     "策略描述错误",
	StrategyScriptErr:   "策略脚本错误",
	StrategyIdErr:       "策略ID错误",
	StrategyLanguageErr: "策略开发语言选择错误",

	//策略执行相关
	TraderNameErr:       "策略执行名称不得为空",
	TraderDescErr:       "策略执行的描述错误",
	TraderIdErr:         "策略执行的id错误",
	TraderStrategyIdErr: "策略执行时，所运行的策略ID错误",
	TraderExchangeIdErr: "策略执行的交易所ID错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Errors]
}
