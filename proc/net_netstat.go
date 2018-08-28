package proc

import (
	"strconv"
	"strings"
)

// Default net/netstat file path
const netStatPath = "/proc/net/netstat"

// GetNetStat read the net/netstat folder
func GetNetStat(path string) (map[string]map[string]uint64, error) {
	lines, err := readFile(path, netStatPath)
	if err != nil {
		return nil, err
	}

	netStatMap := make(map[string]map[string]uint64)

	for i := 1; i < len(lines); i += 2 {
		header := strings.SplitN(lines[i-1], ":", 2)
		if len(header) < 2 {
			continue
		}
		value := strings.SplitN(lines[i], ":", 2)
		if len(value) < 2 {
			continue
		}
		protocol := strings.TrimSpace(header[0])
		protocolMap := make(map[string]uint64)
		keys := strings.Fields(header[1])
		values := strings.Fields(value[1])
		for j := 0; j < len(keys) && j < len(values); j++ {
			protocolMap[keys[j]], _ = strconv.ParseUint(values[j], 10, 64)
		}
		netStatMap[protocol] = protocolMap
	}

	return netStatMap, nil
}
