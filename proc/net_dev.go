package proc

import (
	"strconv"
	"strings"
)

// Default net/dev file path
const netDevPath = "/proc/net/dev"

// NetDev the net/dev file info
type NetDev struct {
	Iface        string `json:"iface"`
	RxBytes      uint64 `json:"rxbytes"`
	RxPackets    uint64 `json:"rxpackets"`
	RxErrs       uint64 `json:"rxerrs"`
	RxDrop       uint64 `json:"rxdrop"`
	RxFifo       uint64 `json:"rxfifo"`
	RxFrame      uint64 `json:"rxframe"`
	RxCompressed uint64 `json:"rxcompressed"`
	RxMulticast  uint64 `json:"rxmulticast"`
	TxBytes      uint64 `json:"txbytes"`
	TxPackets    uint64 `json:"txpackets"`
	TxErrs       uint64 `json:"txerrs"`
	TxDrop       uint64 `json:"txdrop"`
	TxFifo       uint64 `json:"txfifo"`
	TxColls      uint64 `json:"txcolls"`
	TxCarrier    uint64 `json:"txcarrier"`
	TxCompressed uint64 `json:"txcompressed"`
}

// GetNetDev read the net/dev folder
func GetNetDev(path string) ([]NetDev, error) {
	lines, err := readFile(path, netDevPath)
	if err != nil {
		return nil, err
	}

	netdev := make([]NetDev, len(lines[2:])-1)

	for i, line := range lines[2:] {
		colon := strings.Index(line, ":")

		if colon > 0 {
			fields := strings.Fields(line[colon+1:])

			netdev[i].Iface = strings.Replace(line[0:colon], " ", "", -1)
			netdev[i].RxBytes, _ = strconv.ParseUint(fields[0], 10, 64)
			netdev[i].RxPackets, _ = strconv.ParseUint(fields[1], 10, 64)
			netdev[i].RxErrs, _ = strconv.ParseUint(fields[2], 10, 64)
			netdev[i].RxDrop, _ = strconv.ParseUint(fields[3], 10, 64)
			netdev[i].RxFifo, _ = strconv.ParseUint(fields[4], 10, 64)
			netdev[i].RxFrame, _ = strconv.ParseUint(fields[5], 10, 64)
			netdev[i].RxCompressed, _ = strconv.ParseUint(fields[6], 10, 64)
			netdev[i].RxMulticast, _ = strconv.ParseUint(fields[7], 10, 64)
			netdev[i].TxBytes, _ = strconv.ParseUint(fields[8], 10, 64)
			netdev[i].TxPackets, _ = strconv.ParseUint(fields[9], 10, 64)
			netdev[i].TxErrs, _ = strconv.ParseUint(fields[10], 10, 64)
			netdev[i].TxDrop, _ = strconv.ParseUint(fields[11], 10, 64)
			netdev[i].TxFifo, _ = strconv.ParseUint(fields[12], 10, 64)
			netdev[i].TxColls, _ = strconv.ParseUint(fields[13], 10, 64)
			netdev[i].TxCarrier, _ = strconv.ParseUint(fields[14], 10, 64)
			netdev[i].TxCompressed, _ = strconv.ParseUint(fields[15], 10, 64)
		}
	}

	return netdev, nil
}
