package iniconfig

import (
	"io/ioutil"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MySqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Password string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout""`
}

func TestIniConfig(t *testing.T) {
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file failed")
	}
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal error:%v", err)
	}
	t.Logf("unmarshal success,config:%#v", conf)
	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("marshal failed,errs:%v", err)
	}
	t.Logf("marshal success,config:%s", string(confData))

	err = MarshalFile(conf, "./test.ini")
	if err != nil {
		t.Errorf("marshal save file failed,errs:%v", err)
	}

}

func TestIniConfigFile(t *testing.T) {
	filename := "./config.ini"
	var conf Config
	conf.ServerConf.Ip = "localhost"
	conf.ServerConf.Port = 8080
	conf.MySqlConf.Username = "root"
	conf.MySqlConf.Port = 3306
	conf.MySqlConf.Timeout = 3.5
	conf.MySqlConf.Password = "pwd"
	conf.MySqlConf.Host = "192.168.3.2"
	conf.MySqlConf.Database = "test"
	err := MarshalFile(conf, filename)
	if err != nil {
		t.Errorf("marshal file error:%v", err)
	}

	var conf2 Config
	err = UnMarshalFile(filename, &conf2)
	if err != nil {
		t.Errorf("marshal save file failed,errs:%v", err)
	}

	t.Logf("marshal success,config:%#v", conf2)
}
