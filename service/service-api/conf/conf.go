package conf

import (
	"fmt"
	"github.com/bitxx/bitesla/common/iniconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

const (
	DevMode  = "dev"
	ProdMode = "prod"
	TestMode = "test"
)

// 最终从此处获取配置信息
var CurrentConfig *CommonConfig

// Config 基础通用配置
type Config struct {
	Base BaseConfig `ini:"base"`
}

// BaseConfig 涉及到开项目模式配置
type BaseConfig struct {
	Mode string `ini:"mode"`
}

// CommonConfig 项目配置
type CommonConfig struct {
	Mode       string
	ServerConf ServerConfig `ini:"server"`
	LoggerConf LoggerConfig `ini:"logger"`
	MySQL      MySQL        `ini:"mysql"`
	Redis      Redis        `ini:"redis"`
}

type Redis struct {
	Url         string `ini:"url"`
	Password    string `ini:"password"`
	MaxIdle     int    `ini:"max_idle"`
	MaxActive   string `ini:"max_active"`
	IdleTimeout int    `ini:"idle_timeout"`
	DbIndex     int    `ini:"db_index"`
	DefaultKey  string `ini:"bitesla_key"`
}

type MySQL struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Url      string `ini:"url"`
	DBName   string `ini:"db_name"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Ip          string `ini:"ip"`
	Port        string `ini:"port"`
	Proxy       string `ini:"proxy"`
	JwtSecret   string `ini:"jwt_secret"`
	JwtIssuer   string `ini:"jwt_issuer"`
	JwtDuration string `ini:"jwt_duration"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	LogLevel      logrus.Level  `ini:"log_level"`
	EnableDynamic bool          `ini:"enable_dynamic"`
	JSONFormat    bool          `ini:"json_format"`
	BaseFileName  string        `ini:"base_file_name"`
	MaxAgeDays    time.Duration `ini:"max_age_Days"`
}

func LoadConfig() {
	commonConfig, err := readConfig()
	if err != nil {
		panic(err)
	}
	CurrentConfig = commonConfig
}

func readConfig() (*CommonConfig, error) {
	data, err := ioutil.ReadFile("./conf/bitesla-config.ini")
	if err != nil {
		fmt.Println(err)
		return &CommonConfig{}, err
	}
	var conf Config
	err = iniconfig.UnMarshal(data, &conf)
	if err != nil {
		return &CommonConfig{}, err
	}
	switch conf.Base.Mode {
	case DevMode:
		return readConfigByMode("./conf/bitesla-env-dev.ini", DevMode)
	case TestMode:
		return readConfigByMode("./conf/bitesla-env-test.ini", TestMode)
	case ProdMode:
		return readConfigByMode("./conf/bitesla-env-prod.ini", ProdMode)
	}
	return &CommonConfig{}, errors.New("config is error")
}

// 当前环境所有的配置均集中在此处
func readConfigByMode(path string, mode string) (*CommonConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &CommonConfig{}, err
	}
	var conf CommonConfig
	err = iniconfig.UnMarshal(data, &conf)

	if err != nil {
		return &conf, err
	}

	conf.Mode = mode
	return &conf, nil
}
