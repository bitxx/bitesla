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
	GasLimitErr                = 20001
	GasPriceErr                = 20002
	AccountTokenErr            = 20003
	AddressToErr               = 20004
	HashErr                    = 20005
	AddressFromErr             = 20006
	StepErr                    = 20007
	StartErr                   = 20008
	DurationErr                = 20009
	NameErr                    = 20010
	SymbolErr                  = 20011
	ContractConfErr            = 20012
	TransferNotEnd             = 20013
	GnaitCnyEmptyErr           = 20014
	GnaitInfoUrlEmptyErr       = 20015
	CurrencyUrlEmptyErr        = 20016
	GasLimitToken60000Err      = 20017
	GasPriceEmptyErr           = 20018
	GasLimitDefaultEmptyErr    = 20019
	NodeUrlEmptyErr            = 20020
	GasUrlEmptyErr             = 20021
	CustomerTransfersLenErr    = 20022
	GnaitNumberErr             = 20023
	CustomerAddressEmptyErr    = 20024
	CustomerOrderIdEmptyErr    = 20025
	AdminPrivateKeyEmtpyErr    = 20026
	IDGenerateErr              = 20027
	CustomerDuplicateErr       = 20028
	PrivateKeyFormErr          = 20029
	AdminGnaitLow1000Err       = 20030
	AddressFromAndToSameErr    = 20031
	AdminEthLowErr             = 20032
	NotUserAddressErr          = 20033
	ContractAddressNotAllowErr = 20034

	//集群节点相关
	SyncClusterErr = 30001
	HashEmptyErr   = 30002

	//key相关
	KeyServiceIDErr      = 40001
	KeyTypeErr           = 40002
	KeyTypeDescErr       = 40003
	KeyIssuerErr         = 40004
	KeyAppKeyEmptyErr    = 40005
	KeyAppSecretEmptyErr = 40006
	KeyAppTypeErr        = 40007
	KeyJwtEmptyErr       = 40008
	KeyJwtConfErr        = 40009
	KeyNoExist           = 40010

	//IPFS相关
	IpfsIdGetErr            = 50001
	IpfsAddressErr          = 50002
	IpfsUserAddressEmptyErr = 50003
	IpfsUserIdEmptyErr      = 50004
	IpfsDataEmptyErr        = 50005
	IpfsIsPrivateErr        = 50006
	IpfsFileAddErr          = 50007
	IpfsFileGetErr          = 50008
	IpfsFilePublishErr      = 50009
	IpfsFileExistErr        = 50010

	//缓存相关
	CacheInitErr = 60001
	CacheDataErr = 60002
)
