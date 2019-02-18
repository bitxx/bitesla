package trader

import (
	"bytes"
	"encoding/json"
	"github.com/jason-wj/bitesla/common/errs"
	"github.com/jason-wj/bitesla/common/logger"
	"github.com/jason-wj/bitesla/common/util"
	"github.com/jason-wj/bitesla/service/service-strategy/client"
	"github.com/jason-wj/bitesla/service/service-trader/conf"
	"github.com/jason-wj/bitesla/service/service-trader/proto"
	"github.com/nsqio/go-nsq"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//此为nsq消息队列，专门接收要执行的策略

var (
	strategyClient = client.NewStrategyClient()
	//runPath        string
)

// 消费者
type CustomerTraderQueue struct {
}

//处理消息
func (c *CustomerTraderQueue) HandleMessage(msg *nsq.Message) error {
	msg.Finish()

	trader := &bitesla_srv_trader.TraderInfo{}
	err := json.Unmarshal(msg.Body, trader)
	if err != nil {
		logger.Error("最终错误原因：", err)
	}
	c.Run(trader)

	//若返回err!=nil，数据会被重复提交
	return nil
}

//运行策略
func (c *CustomerTraderQueue) Run(trader *bitesla_srv_trader.TraderInfo) {
	//目录:runPath/用户id/交易所id/策略id/策略运行id/
	rootPath := conf.CurrentConfig.ServerConf.RunPath + strconv.FormatInt(trader.CurrentLoginUserID, 10) + "/" + strconv.FormatInt(trader.ExchangeId, 10) + "/" + strconv.FormatInt(trader.StrategyId, 10) + "/" + strconv.FormatInt(trader.TraderId, 10) + "/"
	//defer os.RemoveAll(rootPath)
	//生成路径
	cmd := exec.Command("mkdir", "-p", rootPath)
	err := cmd.Run()
	if err != nil {
		logger.Error("运行路径生成失败，错误原因：", err)
		return
	}

	//获取到策略
	btTraderInfo, err := json.Marshal(trader)
	if err != nil {
		logger.Error("执行策略json转换失败，错误原因：", err)
		return
	}
	strategyDetail, code, err := strategyClient.GetStrategyDetail(btTraderInfo)
	if err != nil || code != errs.Success {
		if err != nil {
			logger.Error("策略详情获取失败，错误原因：", err)
		} else {
			logger.Error("策略详情获取失败，错误原因：", errs.GetMsg(code))
		}
		return
	}
	strategyInfo := &bitesla_srv_trader.StrategyInfo{}
	err = util.ToStruct(strategyDetail, strategyInfo)
	if err != nil {
		logger.Error("策略详情已获取，但转换为结构体时异常，错误原因：", err)
		return
	}

	//生成文件，并将策略代码写入到文件中
	var strategyPath string
	switch strategyInfo.Language {
	case int32(bitesla_srv_trader.Language_GOLANG):
		strategyPath = rootPath + conf.CurrentConfig.ServerConf.GolangDefualtFileName
		//将编写策略需要的main.go文件复制到当前执行目录下
		cmd = exec.Command("cp", "-p", conf.CurrentConfig.ServerConf.GolangMainFilepath+conf.CurrentConfig.ServerConf.GolangMainFileName, rootPath)
		err = cmd.Run()
		if err != nil {
			logger.Info("main.go 文件拷贝失败，错误原因:", err)
		}
	case int32(bitesla_srv_trader.Language_PYTHON):
		strategyPath = rootPath + conf.CurrentConfig.ServerConf.PythonDefualtFileName
	default:
		logger.Error("所选语言不支持，请重新选择")
		return
	}

	codefile, err := os.Create(strategyPath)
	if err != nil {
		logger.Error("创建代码文件错误，错误原因：", err)
		return
	}
	defer codefile.Close()

	_, err = codefile.WriteString(strategyInfo.Script)
	if err != nil {
		logger.Error("策略代码写入文件失败，错误原因：", err)
		return
	}

	//运行策略
	switch strategyInfo.Language {
	case int32(bitesla_srv_trader.Language_GOLANG):
		var out bytes.Buffer
		var outErr bytes.Buffer
		uid := strconv.FormatInt(trader.CurrentLoginUserID, 10)
		exchangeId := strconv.FormatInt(trader.ExchangeId, 10)
		apiKey := trader.ApiKey
		apiSecret := trader.ApiSecret
		token := trader.Token

		//主文件，依赖用的
		mainPath := rootPath + conf.CurrentConfig.ServerConf.GolangMainFileName

		cmd = exec.Command("go", "run", mainPath, strategyPath, uid, exchangeId, token, apiKey, apiSecret)
		cmd.Stdout = &out
		cmd.Stderr = &outErr
		err = cmd.Run()

		logger.Info("输出结果:", out.String())
		if err != nil {
			logger.Error("策略运行失败，错误原因：'", err, "，另一个原因：", outErr.String())
		}

	case int32(bitesla_srv_trader.Language_PYTHON):
	default:
		logger.Error("所选语言不支持，请重新选择")
		return
	}
}

func InitTraderQueue(topic string, channel string, address string) error {
	//runPath = os.Getenv(conf.CurrentConfig.ServerConf.RunPath)

	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 15 * time.Second     //设置服务发现的轮询时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		return err
	}

	//添加策略执行
	consumer := &CustomerTraderQueue{}
	c.AddConcurrentHandlers(consumer, 8) // 添加消费者接口

	//建立NSQLookupd连接
	if err := c.ConnectToNSQD(address); err != nil {
		return err
	}
	return nil
}
