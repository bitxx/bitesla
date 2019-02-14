package conf

import (
	"fmt"
	"testing"
)

func TestIniConfig(t *testing.T) {
	conf, _ := readConfig()
	fmt.Println(conf.IpfsClusterServerConfig.Address)
}
