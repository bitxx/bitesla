package amqp

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Callback func(data []byte)

var (
	conn          *amqp.Connection
	channel       *amqp.Channel
	mqurl         string //链接字符串
	disconnet     bool   //是否失去连接
	isclose       bool
	receives      map[string]chan struct{}
	heartexchange = "heartbeat"
	mu            *sync.Mutex
)

func Init(mq_url, log_db_url string) {
	isclose = false
	mqurl = mq_url
	err := connect()
	if err != nil {
		fmt.Println(mq_url, "can't connection to "+mq_url+" server.err:"+err.Error())
	} else {
		fmt.Println("连接成功")
	}
	receives = make(map[string]chan struct{})
	registerDataBase(log_db_url)
	mu = &sync.Mutex{}
	go heartbeat()
}

func Push(exchange, routingkey string, data []byte) error {
	if channel == nil {
		err := connect()
		if err != nil {
			addErrorRecord(exchange, routingkey, string(data), err.Error(), "FAIL")
			return err
		} else if channel == nil {
			addErrorRecord(exchange, routingkey, string(data), "channel/connection is not open", "FAIL")
			return errors.New("channel/connection is not open")
		}
	}
	err := channel.Publish(exchange, routingkey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         data,
		DeliveryMode: 2,
	})
	if err != nil {
		failOnErr(err, "错误")
		addErrorRecord(exchange, routingkey, string(data), err.Error(), "FAIL")
	} else {
		addErrorRecord(exchange, routingkey, string(data), "", "SUCCESS")
	}
	return err
}

func Receive(queueName string, callback Callback) {
	if channel == nil {
		connect()
	}
	if channel == nil {
		if receive1, ok := receives[queueName]; ok {
			<-receive1
			Receive(queueName, callback)
		}
		return
	}
	msgs, err := channel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		failOnErr(err, "receive init err")
		return
	}
	receive := make(chan struct{})
	setReceive(queueName, receive)
	go func() {
		for d := range msgs {
			callback(d.Body)
		}
		disconnet = true
		if !isclose {
			//等待重连
			<-receive
			Receive(queueName, callback)
		} else {
			fmt.Println(queueName + " receive exit.")
		}
	}()
}

func setReceive(queueName string, receive chan struct{}) {
	mu.Lock()
	defer mu.Unlock()
	receives[queueName] = receive
}

func connect() error {
	var err error
	conn, err = amqp.Dial(mqurl)
	if err != nil {
		failOnErr(err, "connecttion err")
		return err
	}
	channel, err = conn.Channel()
	if err != nil {
		failOnErr(err, "open channel err")
		return err
	}
	return nil
}

func Close() {
	isclose = true
	channel.Close()
	conn.Close()
}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println("mqurl", mqurl, msg, err)
	}
}

//心跳包 检测连接健康状况 自动重连
func heartbeat() {

	for {
		if channel == nil {
			connect()
		}
		if channel == nil {
			sleep()
			continue
		}
		err := channel.Publish(heartexchange, "heartbeat", false, false, amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte("heartbeat"),
			DeliveryMode: 2,
		})
		if err != nil {
			failOnErr(err, "send heartbeat err")
			// channel.Close()
			// conn.Close()
			connect()
			disconnet = true
		} else if disconnet {
			disconnet = false
			for _, r := range receives {
				r <- struct{}{}
			}
		}
		sleep()
	}
}

//心跳包休眠时间
func sleep() {
	time.Sleep(30 * time.Second)
}
