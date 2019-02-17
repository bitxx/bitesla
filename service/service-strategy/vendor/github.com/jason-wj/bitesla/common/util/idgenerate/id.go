package idgenerate

import (
	"fmt"
	"github.com/sony/sonyflake"
)

//账号id规则：
//账号id不能重复
//账号id没有规则，即编码规则不能加入任何和公司运营相关的数据，外部人员无法通过账号id猜测到账号id相关信息。不能被遍历。
//账号id长度固定，且不能太长
//易读，易沟通，不要出现数字字母换乱现象
//生成耗时

var (
	sonyFlake *sonyflake.Sonyflake
)

func getMachineID() (uint16, error) {
	//TODO 后期若账号id量大，则使用zookeeper集群生成
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

func GetId() (id int64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not initd")
		return
	}
	tid, err := sonyFlake.NextID()
	id = int64(tid)
	return
}
