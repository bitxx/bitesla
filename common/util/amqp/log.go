package amqp

import (
	"database/sql"
	"zcm_tools/email"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysql_url string
)

func registerDataBase(mysqlurl string) {
	mysql_url = mysqlurl
}

func addErrorRecord(exchange, routingkey, content, errmsg, state string) {
	db, err := sql.Open("mysql", mysql_url)
	defer db.Close()
	if err != nil {
		go email.SendEmail("添加mq错误记录失败", "exchange:"+exchange+",routingkey:"+routingkey+",content:"+content+",errmsg:"+errmsg+",err:"+err.Error(), "qxw@zcmlc.com;lxy@zcmlc.com")
		return
	}
	sql := "INSERT INTO rbmq_fail_record (exchange_name,routingkey_name,publish_content,err_msg,create_time, state,modify_time) VALUES(?,?,?,?,now(),?,now()) "
	res, sqlerr := db.Exec(sql, exchange, routingkey, content, errmsg, state)
	if sqlerr != nil {
		go email.SendEmail("添加mq错误记录失败", "exchange:"+exchange+",routingkey:"+routingkey+",content:"+content+",errmsg:"+errmsg+",err:插入语句报错::"+sqlerr.Error(), "qxw@zcmlc.com;lxy@zcmlc.com;zhulj@zcmlc")
		return
	}
	num, _ := res.RowsAffected()
	if num <= 0 {
		go email.SendEmail("添加mq错误记录失败", "exchange:"+exchange+",routingkey:"+routingkey+",content:"+content+",errmsg:"+errmsg+",err:影响行数为0", "qxw@zcmlc.com;lxy@zcmlc.com")
	}
}
