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
	RequestDataFmtErr: "请求数据格式异常",

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
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Errors]
}
