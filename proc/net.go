package proc

import (
)

// Net the net info
type Net struct {
	Arp      []NetArp                      `json:"arp"`
	Dev      []NetDev                      `json:"dev"`
	TCP      []NetTCP                      `json:"tcp"`
	UDP      []NetUDP                      `json:"udp"`
	Netstat  map[string]map[string]uint64  `json:"netstat"`
}

// GetNet read the net folder
func GetNet(path string) (*Net, error) {
	var net = Net{}
	var err error
	if net.Arp, err = GetNetArp(""); err != nil {
		return nil, err
	}

	if net.Dev, err = GetNetDev(""); err != nil {
		return nil, err
	}

	if net.TCP, err = GetNetTCP(""); err != nil {
		return nil, err
	}
	if net.UDP, err = GetNetUDP(""); err != nil {
		return nil, err
	}
	if net.Netstat, err = GetNetStat(""); err != nil {
		return nil, err
	}
	return &net, nil
}
