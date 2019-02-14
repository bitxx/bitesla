package db

import (
	"github.com/jason-wj/bitesla/service/service-strategy/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type ConnectPool struct {
}

var instance *ConnectPool
var once sync.Once

var db *gorm.DB
var errDB error

func GetInstance() *ConnectPool {
	once.Do(func() {
		instance = &ConnectPool{}
	})
	return instance
}

//InitPool 初始化数据库连接
// 通过配置文件初始化
func (m *ConnectPool) InitPool() (issucc bool, err error) {
	username := conf.CurrentConfig.MySQL.Username
	password := conf.CurrentConfig.MySQL.Password
	url := conf.CurrentConfig.MySQL.Url
	dbName := conf.CurrentConfig.MySQL.DBName
	db, err = gorm.Open("mysql", username+":"+password+
		"@tcp("+url+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return false, err
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true, nil
}

//InitPoolByParam 直接通过参数初始化
func (m *ConnectPool) InitPoolByParam(username, password, url, dbName string) (issucc bool, err error) {
	db, err = gorm.Open("mysql", username+":"+password+
		"@tcp("+url+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return false, err
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true, nil
}

/*
* @fuc  对外获取数据库连接对象db
 */
func (m *ConnectPool) GetMysqlDB() *gorm.DB {
	return db
}
