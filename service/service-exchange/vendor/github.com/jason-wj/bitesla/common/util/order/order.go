package order

import (
	"fmt"
	"github.com/sony/sonyflake"
)

//订单规则：
//订单号不能重复
//订单号没有规则，即编码规则不能加入任何和公司运营相关的数据，外部人员无法通过订单ID猜测到订单量。不能被遍历。
//订单号长度固定，且不能太长
//易读，易沟通，不要出现数字字母换乱现象
//生成耗时

var (
	sonyFlake *sonyflake.Sonyflake
)

func getMachineID() (uint16, error) {
	//TODO 后期若订单号量大，则使用zookeeper集群生成
	return 0, nil
}

//Init isDefaultmachineId若为true，表示使用sonyflake自己生成的machineId(当前机器内网ip的低16位)
func Init(isDefaultmachineId bool) {
	settings := sonyflake.Settings{}
	if !isDefaultmachineId {
		settings.MachineID = getMachineID
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
}

func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not initd")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
