package proc

import (
	"strings"
)

// Default net/arp file path
const netArpPath = "/proc/net/arp"

// NetArp the net/arp file info
type NetArp struct {
	IPAddress string `json:"ip_address"`
	HWType    string `json:"hw_type"`
	Flags     string `json:"flags"`
	HWAddress string `json:"hw_address"`
	Mask      string `json:"mask"`
	Device    string `json:"device"`
}

// GetNetArp read the net/arp folder
func GetNetArp(path string) ([]NetArp, error) {
	lines, err := readFile(path, netArpPath)
	if err != nil {
		return nil, err
	}

	netarp := make([]NetArp, len(lines)-2)

	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		arp := NetArp{}
		for i, f := range fields {
			switch i {
			case 0:
				arp.IPAddress = f
			case 1:
				arp.HWType = f
			case 2:
				arp.Flags = f
			case 3:
				arp.HWAddress = f
			case 4:
				arp.Mask = f
			case 5:
				arp.Device = f
			}
		}
		netarp[i-1] = arp
	}

	return netarp, nil
}
