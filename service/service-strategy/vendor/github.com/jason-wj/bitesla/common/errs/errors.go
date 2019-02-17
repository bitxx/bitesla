package errs

import "errors"

var (
	DBInitError   = errors.New("数据库初始化异常")
	DBInsertError = errors.New("数据插入失败")

	GasLimitError    = errors.New("gasLimit获取失败")
	GasPriceError    = errors.New("gasPrice获取失败")
	DataConvertError = errors.New("数据转换失败")
)
