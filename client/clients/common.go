package clients

import "net"

type PartKey struct {
	Name        string `json:"name"`
	Broker_name string `json:"brokername"`
	Broker_H_P  string `json:"brokerhp"`
	Offset      int64  `json:"offset"`
	Err         string `json:"err"`
}

type Parts struct {
	PartKey []PartKey `json:"partkeys"`
}

type BrokerInfo struct {
	Name      string `json:"name"`
	Host_port string `json:"hostport"`
}

func GetIpport() string {
	interfaces, err := net.Interfaces()
	ipport := ""
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces { //遍历所有接口 interfaces
		mac := inter.HardwareAddr //获取本机MAC地址
		ipport += mac.String()
	}
	return ipport
}
