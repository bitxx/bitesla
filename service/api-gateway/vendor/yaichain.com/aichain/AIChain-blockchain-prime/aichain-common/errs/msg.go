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
	GasLimitErr:                "gasLimit数据异常",
	GasPriceErr:                "gasPrice数据异常",
	AccountTokenErr:            "交易额输入错误",
	AddressToErr:               "AddressTo不得为空",
	HashErr:                    "交易hash不得为空",
	AddressFromErr:             "AddressFrom不得为空",
	StepErr:                    "step输入异常",
	StartErr:                   "start输入异常",
	DurationErr:                "duration输入异常",
	NameErr:                    "名称不得为空",
	SymbolErr:                  "名称缩写不得为空",
	ContractConfErr:            "合约配置信息错误",
	TransferNotEnd:             "交易未完成，请稍后查询",
	GnaitCnyEmptyErr:           "Gnait价格设置必须大于0",
	GnaitInfoUrlEmptyErr:       "Gnait信息链接不得为空",
	CurrencyUrlEmptyErr:        "法币估值链接不得为空",
	GasLimitToken60000Err:      "用于erc20代币兑换的gaslimit不得低于60000",
	GasPriceEmptyErr:           "gasPrice必须大于0",
	GasLimitDefaultEmptyErr:    "用于eth交易的gaslimit不得低于21000",
	NodeUrlEmptyErr:            "节点链接不得为空",
	GasUrlEmptyErr:             "gas估算链接不得为空",
	CustomerTransfersLenErr:    "待转账列表不得为空",
	GnaitNumberErr:             "gnait必须大于0",
	CustomerAddressEmptyErr:    "用户地址不得为空",
	CustomerOrderIdEmptyErr:    "订单号不得为空",
	AdminPrivateKeyEmtpyErr:    "管理员密钥不得为空",
	IDGenerateErr:              "系统内部发生错误，请重试",
	CustomerDuplicateErr:       "订单号重复",
	PrivateKeyFormErr:          "私钥格式异常，请检查",
	AdminGnaitLow1000Err:       "管理员账户gnait不足1000，请充值",
	AddressFromAndToSameErr:    "发送方和接收方不得为同一账户",
	AdminEthLowErr:             "管理员账户最少要有0.1eth",
	NotUserAddressErr:          "以太坊账户地址不规范",
	ContractAddressNotAllowErr: "用户地址不得输入合约地址",

	//key相关
	KeyServiceIDErr:      "ServiceID错误，请检查",
	KeyTypeErr:           "类型标示异常",
	KeyTypeDescErr:       "类型描述异常",
	KeyIssuerErr:         "签发人不得为空",
	KeyAppKeyEmptyErr:    "appKey不得为空",
	KeyAppSecretEmptyErr: "appSecret不得为空",
	KeyAppTypeErr:        "type必须大于0",
	KeyJwtEmptyErr:       "secret为空",
	KeyJwtConfErr:        "jwt配置异常",
	KeyNoExist:           "信息不存在",

	//ipfs节点相关
	IpfsIdGetErr:            "ipfs节点信息获取失败",
	IpfsAddressErr:          "ipfs的address错误",
	IpfsUserAddressEmptyErr: "用户账户地址不得为空",
	IpfsUserIdEmptyErr:      "user_id不得为空",
	IpfsDataEmptyErr:        "上传数据不得为空",
	IpfsIsPrivateErr:        "无法确定上传文件是否私有，请检查相关字段是否正常",
	IpfsFileAddErr:          "文件添加失败",
	IpfsFileGetErr:          "文件获取失败",
	IpfsFilePublishErr:      "文件未能发送到AI服务器",
	IpfsFileExistErr:        "文件已存在",

	//缓存相关
	CacheInitErr: "请先初始化缓存",
	CacheDataErr: "数据redis缓存异常",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Errors]
}
