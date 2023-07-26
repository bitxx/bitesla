package handler

import (
	"errors"
	"github.com/bitxx/bitesla/common/net/http"
	"github.com/bitxx/bitesla/common/util/idgenerate"
	"github.com/bitxx/bitesla/service/service-trader/conf"
	"github.com/bitxx/bitesla/service/service-trader/db"
)

type traderRepository struct {
}

func (t *traderRepository) listTrader(reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfos *bitesla_srv_trader.TraderInfos) error {
	traders, err := db.GetTraderList(reqTraderInfo.Size, reqTraderInfo.Page, reqTraderInfo.CurrentLoginUserID)
	if err != nil {
		return err
	}
	for _, trader := range traders {
		tmp := &bitesla_srv_trader.TraderInfo{
			StrategyId:  trader.StrategyId,
			ExchangeId:  trader.ExchangeId,
			Description: trader.Description,
			Name:        trader.Name,
			Status:      int32(trader.Status),
			CreateTime:  trader.CreateTime.Unix(),
			UpdateTime:  trader.UpdateTime.Unix(),
		}
		respTraderInfos.TraderInfos = append(respTraderInfos.TraderInfos, tmp)
	}

	return nil
}

func (t *traderRepository) putTrader(reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	var err error
	traderID := reqTraderInfo.TraderId
	if traderID <= 0 {
		traderID, err = idgenerate.GetId()
		if err != nil {
			return errors.New("执行策略所需的id生成失败")
		}
	} else {
		exist := db.IsTraderExist(traderID, reqTraderInfo.CurrentLoginUserID)
		if !exist {
			return errors.New("traderId不存在")
		}
	}

	err = db.AddOrUpdateTrader(traderID, reqTraderInfo.CurrentLoginUserID, reqTraderInfo.StrategyId, reqTraderInfo.ExchangeId, reqTraderInfo.Name, reqTraderInfo.Description)
	return err
}

func (t *traderRepository) getTraderDetail(reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	trader, err := db.GetTraderDetail(reqTraderInfo.CurrentLoginUserID, reqTraderInfo.TraderId)
	respTraderInfo.StrategyId = trader.StrategyId
	respTraderInfo.Name = trader.Name
	respTraderInfo.Description = trader.Description
	respTraderInfo.ExchangeId = trader.ExchangeId
	respTraderInfo.StrategyId = trader.StrategyId
	respTraderInfo.Status = int32(trader.Status)
	respTraderInfo.CreateTime = trader.CreateTime.Unix()
	respTraderInfo.UpdateTime = trader.UpdateTime.Unix()
	return err
}

func (t *traderRepository) updateTraderStatus(reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	err := db.UpdateTraderStatus(reqTraderInfo.CurrentLoginUserID, reqTraderInfo.TraderId, int(reqTraderInfo.Status))
	return err
}

func (t *traderRepository) switchTrader(reqTraderInfo *bitesla_srv_trader.TraderInfo, respTraderInfo *bitesla_srv_trader.TraderInfo) error {
	trader, err := db.GetTraderDetail(reqTraderInfo.CurrentLoginUserID, reqTraderInfo.TraderId)
	if err != nil {
		return err
	}
	reqTraderInfo.Status = int32(trader.Status)
	reqTraderInfo.Description = trader.Description
	reqTraderInfo.Name = trader.Name
	reqTraderInfo.CurrentLoginUserID = trader.UserId
	reqTraderInfo.Status = int32(trader.Status)
	reqTraderInfo.TraderId = trader.Id
	reqTraderInfo.StrategyId = trader.StrategyId
	reqTraderInfo.ExchangeId = trader.ExchangeId

	switch trader.Status {
	case int(bitesla_srv_trader.Status_STOP):
		//推送到队列
		//TODO 此处http client工具需要完善，目前重复实例化，并且代理使用不方便，需要改造到和交易所使用的方式类似
		h := http.NewHttpSend(conf.CurrentConfig.Nsq.HttpUrl + "/pub?topic=" + conf.CurrentConfig.Nsq.TopicDefaultName)
		h.SendType = http.SendtypeJson
		h.SetBody(reqTraderInfo)
		_, err = h.Post(false)
		return err
	case int(bitesla_srv_trader.Status_START):
	}
	return nil

}

// TODO 暂时不考虑实现
func (t *traderRepository) deleteTrader(*bitesla_srv_trader.TraderInfo, *bitesla_srv_trader.TraderInfo) error {
	panic("implement me")
}

//初始化策略执行任务
/*func (t *traderRepository) initTaskTrader(reqTraderInfo *bitesla_srv_trader.TraderInfo) (*model.Global, error) {
	btTraderInfo, err := json.Marshal(reqTraderInfo)
	if err != nil {
		return nil, err
	}
	strategyDetail, code, err := strategyClient.GetStrategyDetail(btTraderInfo)
	if err != nil || code != errs.Success {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errs.GetMsg(code))
	}
	strategyInfo := &bitesla_srv_trader.StrategyInfo{}
	err = util.ToStruct(strategyDetail, strategyInfo)
	if err != nil {
		return nil, err
	}

	global := &model.Global{}
	global.StrategyInfo = strategyInfo
	global.Trader = reqTraderInfo
	global.Client = client.NewStrategyClient()
	global.Ctx.Interrupt = make(chan func(), 1)
	for _, c := range consts {
		err := global.Ctx.Set(c, c)
		if err != nil {
			return nil, err
		}
	}

	err = global.Ctx.Set("Global", &global)
	if err != nil {
		return nil, err
	}
	err = global.Ctx.Set("G", &global)
	if err != nil {
		return nil, err
	}
	err = global.Ctx.Set("Exchange", global.Client)
	if err != nil {
		return nil, err
	}
	err = global.Ctx.Set("E", global.Client)
	if err != nil {
		return nil, err
	}
	return global, nil
}

//停止正在运行的策略
func (t *traderRepository) stopTaskTrader(reqTraderInfo *bitesla_srv_trader.TraderInfo) error {
	panic("implement me")
}

//启动正在运行的策略
func (t *traderRepository) startTaskTrader(reqTraderInfo *bitesla_srv_trader.TraderInfo) error {
	global, err := t.initTaskTrader(reqTraderInfo)
	if err != nil {
		return err
	}
	go func() {
		defer func() {

		}()
	}()
	if _, err := global.Ctx.Run(global.StrategyInfo.Script); err != nil {
		//log
	}
	if main, err := global.Ctx.Get("main"); err != nil || !main.IsFunction() {
		//log
	} else {
		if _, err := main.Call(main); err != nil {
			//log
		}
	}
	return nil
}*/
